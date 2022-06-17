// configuration implimentation is here
package configuration

import (
	"github.com/mhkarimi1383/goAPIBaseProject/types"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	// cfg is a global variable that holds the configuration
	cfg types.Configuration

	// err is a global variable that holds error if any
	err error
)

// initialize the configuration after first import and store it in cfg global variable
func init() {
	err = cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		err = cleanenv.ReadEnv(&cfg)
		if err != nil {
			cfg = types.Configuration{}
		}
	}
}

// others will use this function to get the configuration
func GetConfig() (types.Configuration, error) {
	emptyCfg := types.Configuration{}
	if cfg != emptyCfg {
		return cfg, err
	} else {
		return types.Configuration{}, err
	}
}
