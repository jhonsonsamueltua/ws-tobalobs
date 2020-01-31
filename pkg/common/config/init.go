package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/ws-tobalobs/pkg/models"
)

func InitConfig() *models.Config {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("./config")
	v.SetConfigType("toml")
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}
	conf := &models.Config{}
	if err := v.Unmarshal(&conf); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}

	return conf
}
