package config

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
)

func init() {
	if err := LoadConfig(); err != nil {
		fmt.Println("未初始化配置，请先执行 init 命令")
	}
}

// ================================  ================================

func ConfigDir() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return path.Join(dir, ".devctl")
}

func ConfigFile() string {
	return path.Join(ConfigDir(), "config")
}

// ================================  ================================

func InitConfig() error {
	dir := ConfigDir()
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	configPath := ConfigFile()
	// 只需确保文件存在即可，无需分别处理创建和打开
	file, err := os.OpenFile(configPath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	GlobalConfig = &Config{
		Init:      true,
		Repo:      []string{},
		Project:   []string{},
		Workspace: "",
	}

	if _, err = file.Write(GlobalConfig.JSON()); err != nil {
		return err
	}
	if err = file.Sync(); err != nil {
		return err
	}
	if err = file.Close(); err != nil {
		return err
	}

	return nil
}

func LoadConfig() error {
	file, err := os.Open(ConfigFile())
	if err != nil {
		return err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(bytes) == 0 {
		return errors.New("config file is empty")
	}

	if err = GlobalConfig.Load(bytes); err != nil {
		return err
	}

	return nil
}
