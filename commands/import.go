package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/10gen/stitch-cli/api"
	"github.com/10gen/stitch-cli/models"
	u "github.com/10gen/stitch-cli/user"
	"github.com/10gen/stitch-cli/utils"

	"github.com/mitchellh/cli"
	"github.com/mitchellh/go-homedir"
)

// NewImportCommandFactory returns a new cli.CommandFactory given a cli.Ui
func NewImportCommandFactory(ui cli.Ui) cli.CommandFactory {
	return func() (cli.Command, error) {
		workingDirectory, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		return &ImportCommand{
			BaseCommand: &BaseCommand{
				Name: "import",
				UI:   ui,
			},
			workingDirectory: workingDirectory,
			writeToDirectory: utils.WriteZipToDir,
			fetchApp:         fetchApp,
		}, nil
	}
}

// ImportCommand is used to import a Stitch App
type ImportCommand struct {
	*BaseCommand

	writeToDirectory func(dest string, zipData io.Reader) error
	fetchApp         func(stitchClient api.StitchClient, groupID, appID string) (string, io.ReadCloser, error)
	workingDirectory string

	flagAppID   string
	flagGroupID string
	flagAppPath string
}

// Help returns long-form help information for this command
func (ic *ImportCommand) Help() string {
	return `Import and deploy a stitch application from a local directory.

REQUIRED:
  --app-id [string]

  --group-id [string]

OPTIONS:
  -o, --app-path [string]` +
		ic.BaseCommand.Help()
}

// Synopsis returns a one-liner description for this command
func (ic *ImportCommand) Synopsis() string {
	return `Import and deploy a stitch application from a local directory.`
}

// Run executes the command
func (ic *ImportCommand) Run(args []string) int {
	set := ic.NewFlagSet()

	set.StringVar(&ic.flagAppID, flagAppIDName, "", "")
	set.StringVar(&ic.flagGroupID, flagGroupIDName, "", "")
	set.StringVar(&ic.flagAppPath, "app-path", "", "")

	if err := ic.BaseCommand.run(args); err != nil {
		ic.UI.Error(err.Error())
		return 1
	}

	if err := ic.importApp(); err != nil {
		ic.UI.Error(err.Error())
		return 1
	}

	return 0
}

func (ic *ImportCommand) importApp() error {
	user, err := ic.User()
	if err != nil {
		return err
	}

	if !user.LoggedIn() {
		return u.ErrNotLoggedIn
	}

	authClient, err := ic.AuthClient()
	if err != nil {
		return err
	}

	appPath, err := ic.resolveAppDirectory()
	if err != nil {
		return err
	}

	appInstanceData, err := ic.resolveAppInstanceData(appPath)
	if err != nil {
		return err
	}

	app, err := utils.UnmarshalFromDir(appPath)
	if err != nil {
		return err
	}

	appData, err := json.Marshal(app)
	if err != nil {
		return err
	}

	stitchClient := api.NewStitchClient(ic.flagBaseURL, authClient)

	res, err := stitchClient.Import(appInstanceData.GroupID, appInstanceData.AppID, appData)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		var stitchResponse api.StitchResponse

		if err := json.NewDecoder(res.Body).Decode(&stitchResponse); err != nil {
			return err
		}

		return fmt.Errorf("error: %s", stitchResponse.Error)
	}

	// re-fetch imported app to pull down IDs
	_, body, err := ic.fetchApp(stitchClient, appInstanceData.GroupID, appInstanceData.AppID)
	if err != nil {
		return fmt.Errorf("failed to sync app with local directory after import: %s", err)
	}

	defer body.Close()

	if err := ic.writeToDirectory(appPath, body); err != nil {
		return fmt.Errorf("failed to sync app with local directory after import: %s", err)
	}

	return nil
}

func (ic *ImportCommand) resolveAppDirectory() (string, error) {
	if ic.flagAppPath != "" {
		path, err := homedir.Expand(ic.flagAppPath)
		if err != nil {
			return "", err
		}

		if _, err := os.Stat(path); err != nil {
			return "", errors.New("directory does not exist")
		}
		return path, nil
	}

	return utils.GetRootAppDirectory(ic.workingDirectory)
}

// resolveAppInstanceData loads data for an app from a .stitch file located in the provided directory path
func (ic *ImportCommand) resolveAppInstanceData(path string) (*models.AppInstanceData, error) {
	appInstanceData := &models.AppInstanceData{
		AppID:   ic.flagAppID,
		GroupID: ic.flagGroupID,
	}

	if appInstanceData.AppID == "" || appInstanceData.GroupID == "" {
		if err := mergeAppInstanceDataFromPath(appInstanceData, path); err != nil {
			return nil, err
		}
	}

	if appInstanceData.AppID == "" {
		return nil, errAppIDRequired
	}

	if appInstanceData.GroupID == "" {
		return nil, errGroupIDRequired
	}

	return appInstanceData, nil
}

func mergeAppInstanceDataFromPath(appInstanceData *models.AppInstanceData, path string) error {
	var appInstanceDataFromDotfile models.AppInstanceData
	err := appInstanceDataFromDotfile.UnmarshalFile(path)

	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		return err
	}

	if appInstanceData.GroupID == "" {
		appInstanceData.GroupID = appInstanceDataFromDotfile.GroupID
	}

	if appInstanceData.AppID == "" {
		appInstanceData.AppID = appInstanceDataFromDotfile.AppID
	}

	return nil
}
