package commands

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/10gen/stitch-cli/user"
	u "github.com/10gen/stitch-cli/utils/test"
	gc "github.com/smartystreets/goconvey/convey"

	"github.com/mitchellh/cli"
)

func TestExportCommand(t *testing.T) {
	setup := func() (*ExportCommand, *cli.MockUi) {
		mockUI := cli.NewMockUi()
		cmd, err := NewExportCommandFactory(mockUI)()
		if err != nil {
			panic(err)
		}

		exportCommand := cmd.(*ExportCommand)
		exportCommand.storage = u.NewEmptyStorage()
		return exportCommand, mockUI
	}

	t.Run("should require an app-id", func(t *testing.T) {
		exportCommand, mockUI := setup()
		exitCode := exportCommand.Run([]string{})
		u.So(t, exitCode, gc.ShouldEqual, 1)

		u.So(t, mockUI.ErrorWriter.String(), gc.ShouldContainSubstring, errAppIDRequired.Error())
	})

	t.Run("should require the user to be logged in", func(t *testing.T) {
		exportCommand, mockUI := setup()
		exitCode := exportCommand.Run([]string{`--app-id=my-cool-app`, `--group-id=group-a-doop`})
		u.So(t, exitCode, gc.ShouldEqual, 1)

		u.So(t, mockUI.ErrorWriter.String(), gc.ShouldContainSubstring, user.ErrNotLoggedIn.Error())
	})

	t.Run("when the user is logged in", func(t *testing.T) {
		setup := func(response *http.Response) (*ExportCommand, *cli.MockUi) {
			mockUI := cli.NewMockUi()
			cmd, err := NewExportCommandFactory(mockUI)()
			if err != nil {
				panic(err)
			}

			mockClient := u.NewMockClient([]*http.Response{response})

			exportCommand := cmd.(*ExportCommand)
			exportCommand.client = mockClient
			exportCommand.user = &user.User{
				APIKey:      "my-api-key",
				AccessToken: u.GenerateValidAccessToken(),
			}
			exportCommand.exportToDirectory = func(dest string, r io.Reader) error {
				return nil
			}
			exportCommand.storage = u.NewEmptyStorage()

			return exportCommand, mockUI
		}

		t.Run("it writes response data to the default directory on success", func(t *testing.T) {
			dest, data, r := buildValidExportResponse()
			exportCommand, mockUI := setup(r)

			destination := ""
			var zipData string

			exportCommand.exportToDirectory = func(dest string, r io.Reader) error {
				destination = dest
				b, err := ioutil.ReadAll(r)
				if err != nil {
					panic(err)
				}
				zipData = string(b)
				return nil
			}

			exitCode := exportCommand.Run([]string{`--app-id=my-cool-app`, `--group-id=group-a-doop`})
			u.So(t, exitCode, gc.ShouldEqual, 0)
			u.So(t, mockUI.ErrorWriter.String(), gc.ShouldBeEmpty)

			u.So(t, destination, gc.ShouldEqual, dest)
			u.So(t, zipData, gc.ShouldEqual, data)
		})

		t.Run("it writes response data to the provided destination directory on success", func(t *testing.T) {
			_, data, r := buildValidExportResponse()
			exportCommand, mockUI := setup(r)

			destination := ""
			var zipData string

			exportCommand.exportToDirectory = func(dest string, r io.Reader) error {
				destination = dest
				b, err := ioutil.ReadAll(r)
				if err != nil {
					panic(err)
				}
				zipData = string(b)
				return nil
			}

			outputDir := "some/other/directory/my_app"
			exitCode := exportCommand.Run([]string{`--app-id=my-cool-app`, `--group-id=group-a-doop`, `--output=` + outputDir})
			u.So(t, exitCode, gc.ShouldEqual, 0)
			u.So(t, mockUI.ErrorWriter.String(), gc.ShouldBeEmpty)

			u.So(t, destination, gc.ShouldEqual, outputDir)
			u.So(t, zipData, gc.ShouldEqual, data)
		})

		t.Run("returns an error when the response from the API is unexpected", func(t *testing.T) {
			exportCommand, mockUI := setup(&http.Response{
				StatusCode: http.StatusTeapot,
				Body:       u.NewResponseBody(strings.NewReader(`{ "error": "oh noes" }`)),
			})

			exitCode := exportCommand.Run([]string{`--app-id=my-cool-app`, `--group-id=group-a-doop`})
			u.So(t, exitCode, gc.ShouldEqual, 1)

			u.So(t, mockUI.ErrorWriter.String(), gc.ShouldContainSubstring, "oh noes")
		})
	})
}

func buildValidExportResponse() (string, string, *http.Response) {
	dest := "my_app_123456"
	data := "myZipData"

	r := &http.Response{
		StatusCode: http.StatusOK,
		Body:       u.NewResponseBody(strings.NewReader(data)),
		Header: http.Header{
			"Content-Disposition": []string{fmt.Sprintf(`attachment; filename="%s"`, dest+".zip")},
		},
	}

	return dest, data, r
}
