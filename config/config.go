package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/go-playground/validator/v10"
)

const (
	configPath = "config"
	configFileName = "config"
	configExtension = "json"
)

type Config struct {
	Server             Server
	Postgres           Postgres
}

type Server struct {
	Host string `validate:"required"`
	Port string `validate:"required"`
}

type Postgres struct {
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	DBName   string `validate:"required"`
	SSLMode  string `validate:"required"`
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath(fmt.Sprintf("./%s", configPath))
	v.SetConfigName(configFileName)
	v.SetConfigType(configExtension)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error){
	var c Config
	
	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}
	err = validator.New().Struct(c)
	if err != nil {
		return nil, err
	}
	return &c, err
}