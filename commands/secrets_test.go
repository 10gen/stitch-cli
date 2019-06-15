package commands

import (
	"testing"
	"errors"

	"github.com/10gen/stitch-cli/user"
	u "github.com/10gen/stitch-cli/utils/test"
	"github.com/10gen/stitch-cli/secrets"
	"github.com/10gen/stitch-cli/models"

	"github.com/mitchellh/cli"
	gc "github.com/smartystreets/goconvey/convey"
)

func setUpBasicSecretsCommand(
	addSecretFn func(groupID, appID string, secret secrets.Secret) error,
	removeSecretFn func(groupID, appID, secretID string) error,
) (*SecretsCommand, *cli.MockUi) {
	mockUI := cli.NewMockUi()
	cmd, err := NewSecretsCommandFactory(mockUI)()
	if err != nil {
		panic(err)
	}

	secretsCommand := cmd.(*SecretsCommand)
	secretsCommand.storage = u.NewEmptyStorage()

	mockStitchClient := &u.MockStitchClient{
		AddSecretFn:             addSecretFn,
		RemoveSecretFn:          removeSecretFn,
		FetchAppByClientAppIDFn: func(clientAppID string) (*models.App, error) {
			return &models.App{
				GroupID: "group-id",
				ID:      "app-id",
			}, nil
		},
	}
	secretsCommand.stitchClient = mockStitchClient
	return secretsCommand, mockUI
}

func TestSecretsCommand(t *testing.T) {
	validAddArgs := []string{"--app-id=my-app-abcdef", "--name=foo", "--value=bar"}
	validRemoveArgs := []string{"--app-id=my-app-abcdef", "--secret-id=asdf"}

	t.Run("adding a secret should require the user to be logged in", func(t *testing.T) {
		secretsCommand, mockUI := setUpBasicSecretsCommand(nil, nil)
		exitCode := secretsCommand.Run(append([]string{"add"}, validAddArgs...))
		u.So(t, exitCode, gc.ShouldEqual, 1)

		u.So(t, mockUI.ErrorWriter.String(), gc.ShouldContainSubstring, user.ErrNotLoggedIn.Error())
	})

	t.Run("removing a secret should require the user to be logged in", func(t *testing.T) {
		secretsCommand, mockUI := setUpBasicSecretsCommand(nil, nil)
		exitCode := secretsCommand.Run(append([]string{"remove"}, validRemoveArgs...))
		u.So(t, exitCode, gc.ShouldEqual, 1)

		u.So(t, mockUI.ErrorWriter.String(), gc.ShouldContainSubstring, user.ErrNotLoggedIn.Error())
	})

	t.Run("when the user is logged in", func(t *testing.T) {
		setup := func(
			addSecretsFn func(appID, groupID string, secret secrets.Secret) error,
			removeSecretsFn func(appID, groupID, secretID string) error,
	) (*SecretsCommand, *cli.MockUi) {
			secretsCommand, mockUI := setUpBasicSecretsCommand(addSecretsFn, removeSecretsFn)

			secretsCommand.user = &user.User{
				APIKey:      "my-api-key",
				AccessToken: u.GenerateValidAccessToken(),
			}

			return secretsCommand, mockUI
		}

		t.Run("it fails if there is no sub command", func(t *testing.T) {
				secretsCommand, _ := setup(nil, nil)
				exitCode := secretsCommand.Run(append([]string{}, validAddArgs...))
				u.So(t, exitCode, gc.ShouldEqual, 127)
		})

		t.Run("it fails if there is an invalid sub command", func(t *testing.T) {
			secretsCommand, _ := setup(nil, nil)
			exitCode := secretsCommand.Run(append([]string{"invalid"}, validAddArgs...))
			u.So(t, exitCode, gc.ShouldEqual, 127)
		})

		t.Run("adding a secret fails if the secret name is missing", func(t *testing.T) {
			secretsCommand, mockUI := setup(nil, nil)
			exitCode := secretsCommand.Run(append([]string{"add", "--app-id=my-app-abcdef", "--value=bar"}))
			u.So(t, exitCode, gc.ShouldEqual, 1)
			u.So(t, mockUI.ErrorWriter.String(), gc.ShouldContainSubstring, "must be supplied to create a Secret")
		})

		t.Run("adding a secret it fails if the secret value is missing", func(t *testing.T) {
			secretsCommand, mockUI := setup(nil, nil)
			exitCode := secretsCommand.Run(append([]string{"add", "--app-id=my-app-abcdef", "--name=foo"}))
			u.So(t, exitCode, gc.ShouldEqual, 1)
			u.So(t, mockUI.ErrorWriter.String(), gc.ShouldContainSubstring, "must be supplied to create a Secret")
		})

		t.Run("removing a secret it fails if the secret value is missing", func(t *testing.T) {
			secretsCommand, mockUI := setup(nil, nil)
			exitCode := secretsCommand.Run(append([]string{"remove", "--app-id=my-app-abcdef"}))
			u.So(t, exitCode, gc.ShouldEqual, 1)
			u.So(t, mockUI.ErrorWriter.String(), gc.ShouldContainSubstring, "must be supplied to remove a Secret")
		})

		t.Run("adding a secret fails if adding the secret fails", func(t *testing.T) {
			secretsCommand, mockUI := setup(func(appID, groupID string, secret secrets.Secret) error {
				return errors.New("oopsies")
			}, nil)
			exitCode := secretsCommand.Run(append([]string{"add"}, validAddArgs...))
			u.So(t, exitCode, gc.ShouldEqual, 1)
			u.So(t, mockUI.ErrorWriter.String(), gc.ShouldContainSubstring, "oopsies")
		})

		t.Run("removing a secret fails if removing the secret fails", func(t *testing.T) {
			secretsCommand, mockUI := setup(nil, func(appID, groupID, secretID string) error {
				return errors.New("oopsies")
			})
			exitCode := secretsCommand.Run(append([]string{"remove"}, validRemoveArgs...))
			u.So(t, exitCode, gc.ShouldEqual, 1)
			u.So(t, mockUI.ErrorWriter.String(), gc.ShouldContainSubstring, "oopsies")
		})

		t.Run("it passes the correct flags to AddSecret", func(t *testing.T) {
			var secretName string
			var secretValue string
			secretsCommand, _ := setup(func(appID, groupID string, secret secrets.Secret) error {
				secretName = secret.Name
				secretValue = secret.Value
				return nil
			}, nil)
			exitCode := secretsCommand.Run(append([]string{"add" }, validAddArgs...))
			u.So(t, exitCode, gc.ShouldEqual, 0)
			u.So(t, secretName, gc.ShouldEqual, "foo")
			u.So(t, secretValue, gc.ShouldEqual, "bar")
		})

		t.Run("it passes the correct flags to RemoveSecret", func(t *testing.T) {
			var secretID string
			secretsCommand, _ := setup(nil, func(appID, groupID, id string) error {
				secretID = id
				return nil
			})
			exitCode := secretsCommand.Run(append([]string{"remove"}, validRemoveArgs...))
			u.So(t, exitCode, gc.ShouldEqual, 0)
			u.So(t, secretID, gc.ShouldEqual, "asdf")
		})

		t.Run("adding a secret works", func(t *testing.T) {
			secretsCommand, mockUI := setup(nil, nil)
			exitCode := secretsCommand.Run(append([]string{"add"}, validAddArgs...))
			u.So(t, exitCode, gc.ShouldEqual, 0)
			u.So(t, mockUI.OutputWriter.String(), gc.ShouldContainSubstring, "New secret created")
		})

		t.Run("removing a secret works", func(t *testing.T) {
			secretsCommand, mockUI := setup(nil, nil)
			exitCode := secretsCommand.Run(append([]string{"remove"}, validRemoveArgs...))
			u.So(t, exitCode, gc.ShouldEqual, 0)
			u.So(t, mockUI.OutputWriter.String(), gc.ShouldContainSubstring, "Secret removed")
		})
	})
}
