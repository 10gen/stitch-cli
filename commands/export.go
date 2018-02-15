package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"strings"

	"github.com/10gen/stitch-cli/api"
	u "github.com/10gen/stitch-cli/user"
	"github.com/10gen/stitch-cli/utils"

	"github.com/mitchellh/cli"
)

var (
	errExportMissingFilename = errors.New("the app export response did not specify a filename")
)

// NewExportCommandFactory returns a new cli.CommandFactory given a cli.Ui
func NewExportCommandFactory(ui cli.Ui) cli.CommandFactory {
	return func() (cli.Command, error) {
		return &ExportCommand{
			exportToDirectory: utils.WriteZipToDir,
			BaseCommand: &BaseCommand{
				Name: "export",
				UI:   ui,
			},
		}, nil
	}
}

// ExportCommand is used to export a Stitch App
type ExportCommand struct {
	*BaseCommand

	exportToDirectory func(dest string, zipData io.Reader) error

	flagAppID   string
	flagGroupID string
	flagOutput  string
}

// Help returns long-form help information for this command
func (ec *ExportCommand) Help() string {
	return `Export a stitch application to a local directory.

REQUIRED:
  --app-id [string]

  --group-id [string]

OPTIONS:
  -o, --output [string]
	Directory to write the exported configuration. Defaults to "<app_name>_<timestamp>"` +
		ec.BaseCommand.Help()
}

// Synopsis returns a one-liner description for this command
func (ec *ExportCommand) Synopsis() string {
	return `Export a stitch application to a local directory.`
}

// Run executes the command
func (ec *ExportCommand) Run(args []string) int {
	set := ec.NewFlagSet()

	set.StringVar(&ec.flagAppID, flagAppIDName, "", "")
	set.StringVar(&ec.flagGroupID, flagGroupIDName, "", "")
	set.StringVarP(&ec.flagOutput, "output", "o", "", "")

	if err := ec.BaseCommand.run(args); err != nil {
		ec.UI.Error(err.Error())
		return 1
	}

	if err := ec.run(); err != nil {
		ec.UI.Error(err.Error())
		return 1
	}

	return 0
}

func (ec *ExportCommand) run() error {
	if ec.flagAppID == "" {
		return errAppIDRequired
	}

	if ec.flagGroupID == "" {
		return errGroupIDRequired
	}

	user, err := ec.User()
	if err != nil {
		return err
	}

	if !user.LoggedIn() {
		return u.ErrNotLoggedIn
	}

	authClient, err := ec.AuthClient()
	if err != nil {
		return err
	}

	filename, body, err := fetchApp(api.NewStitchClient(ec.flagBaseURL, authClient), ec.flagGroupID, ec.flagAppID)
	if err != nil {
		return err
	}

	defer body.Close()

	if ec.flagOutput != "" {
		filename = ec.flagOutput
	}

	return ec.exportToDirectory(strings.Replace(filename, ".zip", "", 1), body)
}

func fetchApp(client api.StitchClient, groupID, appID string) (string, io.ReadCloser, error) {
	res, err := client.Export(groupID, appID)
	if err != nil {
		return "", nil, err
	}

	if res.StatusCode != http.StatusOK {
		var stitchResponse api.StitchResponse
		defer res.Body.Close()

		if err := json.NewDecoder(res.Body).Decode(&stitchResponse); err != nil {
			return "", nil, err
		}

		return "", nil, fmt.Errorf("error: %s", stitchResponse.Error)
	}

	_, params, err := mime.ParseMediaType(res.Header.Get("Content-Disposition"))
	if err != nil {
		res.Body.Close()
		return "", nil, err
	}

	filename := params["filename"]
	if len(filename) == 0 {
		res.Body.Close()
		return "", nil, errExportMissingFilename
	}

	return filename, res.Body, nil
}
