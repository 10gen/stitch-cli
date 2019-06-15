package commands

import (
	"os"
	"fmt"
	"path/filepath"

	"github.com/mitchellh/cli"
	"github.com/10gen/stitch-cli/utils"
	u "github.com/10gen/stitch-cli/user"
	"github.com/10gen/stitch-cli/secrets"
)

// NewDiffCommandFactory returns a new cli.CommandFactory given a cli.Ui
func NewSecretsCommandFactory(ui cli.Ui) cli.CommandFactory {
	return func() (cli.Command, error) {
		c := cli.NewCLI(filepath.Base(os.Args[0]), utils.CLIVersion)
		c.Args = os.Args[1:]

		sc := &SecretsCommand{
			BasePath: filepath.Base(os.Args[0]),
			BaseCommand: &BaseCommand{
				Name: "secrets",
				CLI:  c,
				UI:   ui,
			},
		}

		c.Commands = map[string]cli.CommandFactory{
			"add": NewSecretsAddCommandFactory(sc),
			"remove": NewSecretsRemoveCommandFactory(sc),
		}

		return sc, nil
	}
}

// SecretsCommand is used to run CRUD operations on a Stitch App's secrets
type SecretsCommand struct {
	*BaseCommand

	Name     string
	BasePath string
}

// Synopsis returns a one-liner description for this command
func (sc *SecretsCommand) Synopsis() string {
	return "Add or remove secrets for your Stitch App."
}

// Help returns long-form help information for this command
func (sc *SecretsCommand) Help() string {
	return sc.BaseCommand.CLI.HelpFunc(sc.CLI.Commands)
}

// Run executes the command
func (sc *SecretsCommand) Run(args []string) int {
	sc.BaseCommand.CLI.Args = args

	exitStatus, err := sc.BaseCommand.CLI.Run()
	if err != nil {
		sc.BaseCommand.UI.Error(err.Error())
	}

	return exitStatus
}

const (
	flagSecretName  = "name"
	flagSecretValue = "value"
	flagSecretID    = "secret-id"
)

var (
	errSecretNameRequired = fmt.Errorf("a name (--%s=[string]) must be supplied to create a Secret", flagSecretName)
	errSecretValueRequired = fmt.Errorf("a value (--%s=[string]) must be supplied to create a Secret", flagSecretValue)
	errSecretIDRequired = fmt.Errorf("an ID (--%s=[string]) must be supplied to remove a Secret", flagSecretID)
)

// NewSecretsAddCommandFactory returns a new cli.CommandFactory given a cli.Ui
func NewSecretsAddCommandFactory(sc *SecretsCommand) cli.CommandFactory {
	return func() (cli.Command, error) {
		sc.Name = "add"

		return &SecretsAddCommand{
			SecretsCommand: sc,
		}, nil
	}
}

// SecretsCommand is used to run CRUD operations on a Stitch App's secrets
type SecretsAddCommand struct {
	*SecretsCommand

	flagAppID       string
	flagSecretName  string
	flagSecretValue string
}

// Synopsis returns a one-liner description for this command
func (sac *SecretsAddCommand) Synopsis() string {
	return "Add a secret to your Stitch App."
}

// Help returns long-form help information for this command
func (sac *SecretsAddCommand) Help() string {
	return `Add a secret to your Stitch Application.

REQUIRED:
  --app-id [string]
	The App ID for your app (i.e. the name of your app followed by a unique suffix, like "my-app-nysja").

  --name [string]
	The name of your secret.

  --value [string]
	The value of your secret.
	` +
		sac.BaseCommand.Help()
}

// Run executes the command
func (sac *SecretsAddCommand) Run(args []string) int {
	flags := sac.NewFlagSet()

	flags.StringVar(&sac.flagAppID, flagAppIDName, "", "")
	flags.StringVar(&sac.flagSecretName, flagSecretName, "", "")
	flags.StringVar(&sac.flagSecretValue, flagSecretValue, "", "")

	if err := sac.BaseCommand.run(args); err != nil {
		sac.UI.Error(err.Error())
		return 1
	}

	if err := sac.addSecret(); err != nil {
		sac.UI.Error(err.Error())
		return 1
	}

	return 0
}

func (sac *SecretsAddCommand) addSecret() error {
	if sac.flagAppID == "" {
		return errAppIDRequired
	}

	if sac.flagSecretName == "" {
		return errSecretNameRequired
	}

	if sac.flagSecretValue == "" {
		return errSecretValueRequired
	}

	user, err := sac.User()
	if err != nil {
		return err
	}

	if !user.LoggedIn() {
		return u.ErrNotLoggedIn
	}

	stitchClient, err := sac.StitchClient()
	if err != nil {
		return err
	}

	app, err := stitchClient.FetchAppByClientAppID(sac.flagAppID)
	if err != nil {
		return err
	}

	if addErr := stitchClient.AddSecret(app.GroupID, app.ID, secrets.Secret{
		Name:  sac.flagSecretName,
		Value: sac.flagSecretValue,
	}); addErr != nil {
		return addErr
	}

	sac.UI.Info(fmt.Sprintf("New secret created: %s", sac.flagSecretName))
	return nil
}

// NewSecretsRemoveCommandFactory returns a new cli.CommandFactory given a cli.Ui
func NewSecretsRemoveCommandFactory(sc *SecretsCommand) cli.CommandFactory {
	return func() (cli.Command, error) {
		sc.Name = "remove"

		return &SecretsRemoveCommand{
			SecretsCommand: sc,
		}, nil
	}
}

// SecretsCommand is used to run CRUD operations on a Stitch App's secrets
type SecretsRemoveCommand struct {
	*SecretsCommand

	flagAppID       string
	flagSecretID    string
}

// Synopsis returns a one-liner description for this command
func (src *SecretsRemoveCommand) Synopsis() string {
	return "Remove a secret from your Stitch App."
}

// Help returns long-form help information for this command
func (src *SecretsRemoveCommand) Help() string {
	return `Remove a secret from your Stitch Application.

REQUIRED:
  --app-id [string]
	The App ID for your app (i.e. the name of your app followed by a unique suffix, like "my-app-nysja").

  --secret-id [string]
	The ID of your secret.
	` +
		src.BaseCommand.Help()
}

// Run executes the command
func (src *SecretsRemoveCommand) Run(args []string) int {
	flags := src.NewFlagSet()

	flags.StringVar(&src.flagAppID, flagAppIDName, "", "")
	flags.StringVar(&src.flagSecretID, flagSecretID, "", "")

	if err := src.BaseCommand.run(args); err != nil {
		src.UI.Error(err.Error())
		return 1
	}

	if err := src.removeSecret(); err != nil {
		src.UI.Error(err.Error())
		return 1
	}

	return 0
}

func (src *SecretsRemoveCommand) removeSecret() error {
	if src.flagAppID == "" {
		return errAppIDRequired
	}

	if src.flagSecretID == "" {
		return errSecretIDRequired
	}

	user, err := src.User()
	if err != nil {
		return err
	}

	if !user.LoggedIn() {
		return u.ErrNotLoggedIn
	}

	stitchClient, err := src.StitchClient()
	if err != nil {
		return err
	}

	app, err := stitchClient.FetchAppByClientAppID(src.flagAppID)
	if err != nil {
		return err
	}

	if removeErr := stitchClient.RemoveSecret(app.GroupID, app.ID, src.flagSecretID); removeErr != nil {
		return removeErr
	}

	src.UI.Info(fmt.Sprintf("Secret removed: %s", src.flagSecretID))
	return nil
}
