package viper

import (
	"github.com/spf13/viper"
)

func Init() (*viper.Viper, error) {
	viper := viper.New()

	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return viper, nil
}
