package models

import (
	"path/filepath"

	"github.com/10gen/stitch-cli/utils"

	"gopkg.in/yaml.v2"
)

const appInstanceDataFileName string = ".stitch"

// AppInstanceData defines data pertaining to a specific deployment of a Stitch application
type AppInstanceData struct {
	AppID   string `yaml:"app_id"`
	GroupID string `yaml:"group_id"`
}

// UnmarshalFile unmarshals data from a local .stitch project file into an AppInstanceData
func (aic *AppInstanceData) UnmarshalFile(path string) error {
	return utils.ReadAndUnmarshalInto(yaml.Unmarshal, filepath.Join(path, appInstanceDataFileName), &aic)
}
