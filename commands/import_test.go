package commands

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/10gen/stitch-cli/api"
	"github.com/10gen/stitch-cli/user"
	u "github.com/10gen/stitch-cli/utils/test"
	gc "github.com/smartystreets/goconvey/convey"

	"github.com/mitchellh/cli"
)

func TestImportCommand(t *testing.T) {
	setUpBasicCommand := func() (*ImportCommand, *cli.MockUi) {
		mockUI := cli.NewMockUi()
		cmd, err := NewImportCommandFactory(mockUI)()
		if err != nil {
			panic(err)
		}

		importCommand := cmd.(*ImportCommand)
		importCommand.storage = u.NewEmptyStorage()
		importCommand.writeToDirectory = func(dest string, r io.Reader) error {
			return nil
		}

		importCommand.fetchApp = func(stitchClient api.StitchClient, groupID, appID string) (string, io.ReadCloser, error) {
			return "", u.NewResponseBody(bytes.NewReader([]byte{})), nil
		}
		return importCommand, mockUI
	}

	validArgs := []string{"--app-id=my-cool-app", "--group-id=group-a-doop"}

	t.Run("should require the user to be logged in", func(t *testing.T) {
		importCommand, mockUI := setUpBasicCommand()
		exitCode := importCommand.Run(validArgs)
		u.So(t, exitCode, gc.ShouldEqual, 1)

		u.So(t, mockUI.ErrorWriter.String(), gc.ShouldContainSubstring, user.ErrNotLoggedIn.Error())
	})

	t.Run("when the user is logged in", func(t *testing.T) {
		setup := func(responses []*http.Response) (*ImportCommand, *cli.MockUi) {
			importCommand, mockUI := setUpBasicCommand()

			importCommand.client = u.NewMockClient(responses)
			importCommand.user = &user.User{
				APIKey:      "my-api-key",
				AccessToken: u.GenerateValidAccessToken(),
			}

			return importCommand, mockUI
		}

		t.Run("it fails if given an invalid flagAppPath", func(t *testing.T) {
			importCommand, mockUI := setup([]*http.Response{
				{
					Body: u.NewResponseBody(strings.NewReader(`{ "error": "ruh roh" }`)),
				},
			})
			exitCode := importCommand.Run(append([]string{"--app-path=/somewhere/bogus"}, validArgs...))
			u.So(t, exitCode, gc.ShouldEqual, 1)

			u.So(t, mockUI.ErrorWriter.String(), gc.ShouldContainSubstring, "directory does not exist")
		})

		t.Run("it succeeds if given a valid flagAppPath", func(t *testing.T) {
			responses := []*http.Response{
				{
					StatusCode: http.StatusNoContent,
					Body:       u.NewResponseBody(bytes.NewReader([]byte{})),
				},
			}
			importCommand, mockUI := setup(responses)
			exitCode := importCommand.Run(append([]string{"--app-path=../testdata/full_app"}, validArgs...))
			u.So(t, exitCode, gc.ShouldEqual, 0)
			u.So(t, mockUI.ErrorWriter.String(), gc.ShouldBeEmpty)
		})

		t.Run("syncing data after a successful import", func(t *testing.T) {
			t.Run("on failure to fetch the app reports an error", func(t *testing.T) {
				responses := []*http.Response{
					{
						StatusCode: http.StatusNoContent,
						Body:       u.NewResponseBody(bytes.NewReader([]byte{})),
					},
				}
				importCommand, mockUI := setup(responses)
				importCommand.fetchApp = func(stitchClient api.StitchClient, groupID, appID string) (string, io.ReadCloser, error) {
					return "", nil, fmt.Errorf("oh no")
				}

				exitCode := importCommand.Run(append([]string{"--app-path=../testdata/full_app"}, validArgs...))
				u.So(t, exitCode, gc.ShouldEqual, 1)
				u.So(t, mockUI.ErrorWriter.String(), gc.ShouldContainSubstring, "failed to sync app")
			})

			t.Run("on success", func(t *testing.T) {
				type testCase struct {
					Description       string
					Args              []string
					WorkingDirectory  string
					ExpectedDirectory string
				}

				for _, tc := range []testCase{
					{
						Description:       "it writes data to the provided directory",
						Args:              append([]string{"--app-path=../testdata/simple_app"}, validArgs...),
						WorkingDirectory:  "",
						ExpectedDirectory: "../testdata/simple_app",
					},
					{
						Description:       "it writes data to the working directory when using a .stitch.file",
						Args:              []string{},
						WorkingDirectory:  "../testdata/simple_app_with_instance_data",
						ExpectedDirectory: "../testdata/simple_app_with_instance_data",
					},
				} {
					t.Run(tc.Description, func(t *testing.T) {
						responses := []*http.Response{
							{
								StatusCode: http.StatusNoContent,
								Body:       u.NewResponseBody(bytes.NewReader([]byte{})),
							},
						}
						importCommand, mockUI := setup(responses)
						importCommand.workingDirectory = tc.WorkingDirectory
						importCommand.fetchApp = func(stitchClient api.StitchClient, groupID, appID string) (string, io.ReadCloser, error) {
							return "", u.NewResponseBody(strings.NewReader("export response")), nil
						}

						destinationDirectory := ""
						writeContent := ""

						importCommand.writeToDirectory = func(dest string, zipData io.Reader) error {
							b, err := ioutil.ReadAll(zipData)
							if err != nil {
								return err
							}
							destinationDirectory = dest
							writeContent = string(b)
							return nil
						}

						exitCode := importCommand.Run(tc.Args)
						u.So(t, exitCode, gc.ShouldEqual, 0)
						u.So(t, mockUI.ErrorWriter.String(), gc.ShouldBeEmpty)
						u.So(t, destinationDirectory, gc.ShouldEqual, tc.ExpectedDirectory)
						u.So(t, writeContent, gc.ShouldEqual, "export response")
					})
				}
			})
		})

		t.Run("it fails with an error if the response to the import request is invalid", func(t *testing.T) {
			responses := []*http.Response{
				{
					StatusCode: http.StatusBadRequest,
					Body:       u.NewResponseBody(strings.NewReader(`{ "error": "oh noes" }`)),
				},
			}
			importCommand, mockUI := setup(responses)
			exitCode := importCommand.Run(append([]string{"--app-path=../testdata/simple_app"}, validArgs...))
			u.So(t, exitCode, gc.ShouldEqual, 1)
			u.So(t, mockUI.ErrorWriter.String(), gc.ShouldContainSubstring, "oh noes")
		})

		t.Run("it fails if an app-id is not provided", func(t *testing.T) {
			importCommand, mockUI := setup([]*http.Response{})
			exitCode := importCommand.Run([]string{"--app-path=../testdata/simple_app"})
			u.So(t, exitCode, gc.ShouldEqual, 1)

			u.So(t, mockUI.ErrorWriter.String(), gc.ShouldContainSubstring, errAppIDRequired.Error())
		})

		t.Run("it succeeds if it can grab instance data from the .stitch file at the provided path", func(t *testing.T) {
			responses := []*http.Response{
				{
					StatusCode: http.StatusNoContent,
					Body:       u.NewResponseBody(bytes.NewReader([]byte{})),
				},
			}
			importCommand, _ := setup(responses)
			exitCode := importCommand.Run([]string{"--app-path=../testdata/simple_app_with_instance_data"})
			u.So(t, exitCode, gc.ShouldEqual, 0)
		})

		t.Run("it succeeds if it can grab instance data from the .stitch file in the current directory", func(t *testing.T) {
			responses := []*http.Response{
				{
					StatusCode: http.StatusNoContent,
					Body:       u.NewResponseBody(bytes.NewReader([]byte{})),
				},
			}
			importCommand, _ := setup(responses)
			importCommand.workingDirectory = "../testdata/simple_app_with_instance_data"
			exitCode := importCommand.Run([]string{})
			u.So(t, exitCode, gc.ShouldEqual, 0)
		})
	})
}
