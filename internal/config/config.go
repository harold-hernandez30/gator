package config

import (
	"encoding/json"
	"log"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl  string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (config *Config) SetUser (username string) {

	config.CurrentUserName = username
	err := writeConfig(*config)

	if err != nil {
		log.Fatal("Could not write to config")
	}
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return homeDir + string(os.PathSeparator) + configFileName, nil
}

func writeConfig(config Config) error {
	configFilePath, err := getConfigFilePath()

	if err != nil {
		return err
	}

	res, err := json.Marshal(config)

	if err != nil {
		return err
	}

	writeErr := os.WriteFile(configFilePath, res, 0666)
	if writeErr != nil {
		return writeErr
	}

	return nil
}

func Read() Config {

	config, err := getConfigFilePath()

	if err != nil {
		log.Fatal("could not find config")
		return Config{}
	}

	file, err := os.ReadFile(config)

	if err != nil {
		log.Fatal("error reading file config")
		return Config{}
	}

	dbConfig := Config{}
	json.Unmarshal(file, &dbConfig)

	return dbConfig
}
