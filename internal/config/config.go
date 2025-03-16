package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	var result Config
	file, err := os.Open(home + "/.gatorconfig.json")
	if err != nil {
		return Config{}, err
	}
	// dont forget close
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return Config{}, err
	}
	return result, nil
}

func (c *Config) SetUser(username string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	c.CurrentUserName = username

	filePath := home + "/.gatorconfig.json"

	//NOTE: convert struct to json
	jsonData, err := json.Marshal(&c)
	if err != nil {
		return err
	}

	if err = os.WriteFile(filePath, jsonData, 0644); err != nil {
		return err
	}

	return nil
}
