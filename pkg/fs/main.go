package fs

import (
	"github.com/rskvp/necode/services/chore/config"
)

type (
	fsConfigProvider struct {
		configDir string
		env       string
	}
)

var _ config.ConfigProvider = (*fsConfigProvider)(nil)

// NewFSConfigProvider creates a default file system based Config Provider
func NewFSConfigProvider(configDir string, env string) config.ConfigProvider {
	return &fsConfigProvider{
		configDir,
		env,
	}
}

func (a *fsConfigProvider) GetConfig() (*config.Config, error) {
	cfg, err := LoadConfig(a.configDir, a.env)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
