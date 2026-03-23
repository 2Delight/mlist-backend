package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const dbConnectionStringTemplate = "%s://%s:%s@%s:%d/%s?sslmode=%s"

type DB struct {
	Schema         string `yaml:"schema"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	IP             string `yaml:"ip"`
	Port           uint16 `yaml:"port"`
	DataBase       string `yaml:"database"`
	SSL            string `yaml:"ssl"`
	MigrationsPath string `yaml:"migrations_path"`
}

func (db DB) GetDBURL() string {
	return fmt.Sprintf(
		dbConnectionStringTemplate,
		db.Schema,
		db.User,
		db.Password,
		db.IP,
		db.Port,
		db.DataBase,
		db.SSL,
	)
}

type Config struct {
	DB     `yaml:"db"`
	Server struct {
		Port uint16 `yaml:"port"`
	} `yaml:"server"`
	Metrics struct {
		Port uint16 `yaml:"port"`
	} `yaml:"metrics"`
}

func Parse(path string) (Config, error) {
	bytes, err := os.ReadFile(path) // nolint:gosec
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
