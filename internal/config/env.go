package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Environment struct {
	ServerHost string `mapstructure:"SERVER_HOST"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPwd      string `mapstructure:"DB_Password"`
	DBName     string `mapstructure:"DB_NAME"`
}

var Env *Environment

func LoadEnvConfiguration(configPath, configName, configType string) error {
	var config *Environment

	// config 파일 경로 설정
	viper.AddConfigPath(configPath)
	// config 파일 이름 설정
	viper.SetConfigName(configName)
	// config 파일 타입 설정 ( yaml / toml / env )
	viper.SetConfigType(configType)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("could not read the config file: %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Errorf("could not unmarshal: %v", err)
	}

	Env = config
	return nil
}
