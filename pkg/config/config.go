package config

import (
	"encoding/json"
	"fmt"
	"os"
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

func RepoPath() []string {
	if !GlobalConfig.Init {
		fmt.Println("未初始化配置，请先执行 init 命令")
		os.Exit(1)
	}

	return GlobalConfig.Repo
}

func Workspace() string {
	if !GlobalConfig.Init {
		fmt.Println("未初始化配置，请先执行 init 命令")
		os.Exit(1)
	}

	workspace := GlobalConfig.Workspace
	if workspace == "" {
		fmt.Println("未设置 workspace，请先配置 workspace")
		os.Exit(1)
	}
	if err := os.MkdirAll(workspace, 0755); err != nil && !os.IsExist(err) {
		fmt.Printf("创建 workspace 目录失败: %v\n", err)
		os.Exit(1)
	}

	return workspace
}
