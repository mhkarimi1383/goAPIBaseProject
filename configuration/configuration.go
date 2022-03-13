// configuration implimentation is here
package configuration

import (
	"github.com/mhkarimi1383/goAPIBaseProject/structures"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	cfg *structures.Configuration
	err error
)

func init() {
	err = cleanenv.ReadConfig("config.yml", cfg)
	if err != nil {
		err = cleanenv.ReadEnv(cfg)
		if err != nil {
			cfg = nil
		}
	}
}

func GetConfig() (structures.Configuration, error) {
	if cfg != nil {
		return *cfg, err
	} else {
		return structures.Configuration{}, err
	}
}
