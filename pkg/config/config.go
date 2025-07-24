package config

import (
	"encoding/json"
)

var GlobalConfig *Config = &Config{
	Init:      false,
	Repo:      []string{},
	Project:   []string{},
	Workspace: "",
}

type Config struct {
	Init      bool     `json:"init"`
	Repo      []string `json:"repo"`
	Project   []string `json:"project"`
	Workspace string   `json:"workspace"`
}

func (c *Config) JSON() []byte {
	bytes, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return bytes
}

func (c *Config) Load(bytes []byte) error {
	if err := json.Unmarshal(bytes, c); err != nil {
		return err
	}
	return nil
}
