package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port string
	Host string
}

type DbConfig struct {
	Host string
}

type Configuration struct {
	Server ServerConfig
	DB     DbConfig
}

func LoadEnvConfiguration(configPath, configName, configType string) (*Configuration, error) {
	var config *Configuration

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

	return config, nil
}
