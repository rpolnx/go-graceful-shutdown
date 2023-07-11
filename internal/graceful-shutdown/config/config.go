package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	App App `yaml:"app"`
	Api Api `yaml:"api"`
}

type App struct {
	Name    string `yaml:"name"`
	Profile string `yaml:"profile"`
}

type Api struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	ApiTimeout      int    `yaml:"apiTimeout"`
	GracefulTimeout int    `yaml:"gracefulTimeout"`
}

func LoadConfig() (config *Configuration, e error) {
	initial := time.Now()

	logrus.Infof("[Env Config] Initializing env variable configurations")

	viper.SetConfigName("configs/application")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Load Config from File
	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			logrus.Warnf("No Config file found, loaded config from Environment - Default path ./conf")
		default:
			return nil, errors.Wrap(err, "config.LoadConfig")
		}
	}

	err := viper.Unmarshal(&config)

	if err != nil {
		logrus.Error("[Env Config] Error serializing config", err)
		return nil, errors.Wrap(err, "[Env Config] Error serializing config")
	}

	delta := time.Since(initial).Milliseconds()
	logrus.Infof(fmt.Sprintf("[Env Config] Finalized env variable configurations in %dus", delta))

	return config, nil
}
