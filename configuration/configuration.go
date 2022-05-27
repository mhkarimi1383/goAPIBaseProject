// configuration implimentation is here
package configuration

import (
	"github.com/mhkarimi1383/goAPIBaseProject/types"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	cfg types.Configuration
	err error
)

func init() {
	err = cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		err = cleanenv.ReadEnv(&cfg)
		if err != nil {
			cfg = types.Configuration{}
		}
	}
}

func GetConfig() (types.Configuration, error) {
	emptyCfg := types.Configuration{}
	if cfg != emptyCfg {
		return cfg, err
	} else {
		return types.Configuration{}, err
	}
}
