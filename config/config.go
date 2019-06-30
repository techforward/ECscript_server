package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Environment struct {
	Development      Conf `yaml:"development"`
	LocalDevelopment Conf `yaml:"local_development"`
	Production       Conf `yaml:"production"`
	Test             Conf `yaml:"test"`
}

type Conf struct {
	Database Database `yaml:"db"`
}

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	IP       string `yaml:"ip"`
	Port     string `yaml:"port"`
}

func SetEnvironment(env string) Conf {

	var Config Conf

	buf, err := ioutil.ReadFile("../config/environment.yaml")
	// buf, err := ioutil.ReadFile("./config/environment.yaml")
	if err != nil {
		panic(err)
	}

	var environment Environment

	err = yaml.Unmarshal(buf, &environment)
	if err != nil {
		panic(err)
	}

	switch env {
	case "development":
		Config = environment.Development
	case "local_development":
		Config = environment.LocalDevelopment
	case "production":
		Config = environment.Production
	case "test":
		Config = environment.Test
	}
	return Config
}
