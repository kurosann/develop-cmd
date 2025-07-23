package config

import (
	"encoding/json"
	"io"
	"os"
	"path"
)

type Config struct {
	Repo      []string
	Project   []string
	Workspace string
}

func ConfigPath() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return path.Join(dir, ".devctl", "config")
}

func DefaultConfig() []byte {
	config := Config{}
	bytes, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}
	return bytes
}

func LoadConfig() (*Config, error) {
	file, err := os.Open(ConfigPath())
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(ConfigPath())
			if err != nil {
				return nil, err
			}
			file.Write(DefaultConfig())
			file.Close()
			return LoadConfig()
		}
		return nil, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	if err = json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
