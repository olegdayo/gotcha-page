package sharedData

import "github.com/offluck/gotcha-page/backend/internal/config"

var (
	conf config.Config
)

func InitConfig(path string) error {
	var err error
	conf, err = config.Import(path)
	return err
}

func GetConfig() config.Config {
	return conf
}
