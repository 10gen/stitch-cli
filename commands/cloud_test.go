package commands

import (
	"os/exec"
	"strings"
	"testing"

	u "github.com/10gen/stitch-cli/utils/test"
	"github.com/10gen/stitch-cli/utils/test/mdbcloud"
	gc "github.com/smartystreets/goconvey/convey"
)

func TestCloudCommands(t *testing.T) {
	u.SkipUnlessMongoDBCloudRunning(t)
	serverBaseURL := u.StitchServerBaseURL()

	// test login
	loginArgs := []string{
		"run",
		"../main.go",
		"login",
		"--config-path",
		"../cli_conf",
		"--base-url",
		serverBaseURL,
		"--username",
		u.MongoDBCloudUsername(),
		"--api-key",
		u.MongoDBCloudAPIKey(),
	}

	err := exec.Command("go", loginArgs...).Run()
	u.So(t, err, gc.ShouldBeNil)
	err = exec.Command("ls", "../cli_conf").Run()
	u.So(t, err, gc.ShouldBeNil)

	// test import
	importArgs := []string{
		"run",
		"../main.go",
		"import",
		"--config-path",
		"../cli_conf",
		"--base-url",
		serverBaseURL,
		"--path",
		"../testdata/simple_app_with_cluster",
		"--project-id",
		u.MongoDBCloudGroupID(),
		"--yes",
	}
	out, err := exec.Command("go", importArgs...).Output()
	u.So(t, err, gc.ShouldBeNil)

	// test export
	importOut := string(out)
	appID := importOut[strings.Index(importOut, "'simple-app-")+1 : len(importOut)-2]
	exportArgs := []string{
		"run",
		"../main.go",
		"export",
		"--config-path",
		"../cli_conf",
		"--base-url",
		serverBaseURL,
		"--app-id",
		appID,
		"-o",
		"../exported_app",
		"--yes",
	}
	err = exec.Command("go", exportArgs...).Run()
	u.So(t, err, gc.ShouldBeNil)

	atlasClient := mdbcloud.NewClient(u.MongoDBCloudPublicAPIBaseURL(), u.MongoDBCloudAtlasAPIBaseURL()).
		WithAuth(u.MongoDBCloudUsername(), u.MongoDBCloudAPIKey())

	defer atlasClient.DeleteDatabaseUser(u.MongoDBCloudGroupID(), "mongodb-stitch-"+appID)

	out, _ = exec.Command("cat", "../exported_app/stitch.json").Output()
	u.So(t, string(out), gc.ShouldContainSubstring, "\"app_id\":")
	out, _ = exec.Command(
		"diff",
		"../testdata/simple_app_with_cluster/stitch.json",
		"../exported_app/stitch.json",
	).Output()
	u.So(t, out, gc.ShouldHaveLength, 0)

	out, _ = exec.Command("cat", "../exported_app/services/mongodb-atlas/config.json").Output()
	u.So(t, string(out), gc.ShouldContainSubstring, "\"id\":")
	out, _ = exec.Command(
		"diff",
		"../testdata/simple_app_with_cluster/services/mongodb-atlas/config.json",
		"../exported_app/services/mongodb-atlas/config.json",
	).Output()
	u.So(t, out, gc.ShouldHaveLength, 0)
}
