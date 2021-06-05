package config

import (
	"github.com/spf13/viper"
)

func Load(dest interface{}, configPath string) error {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(&dest)
}
