package config

import (
	"errors"
	"os"
	"path"
	"testing"
)

// 测试 configDir 和 configFile 函数
func TestConfigDirAndFile(t *testing.T) {
	dir := configDir()
	if dir == "" {
		t.Error("configDir 返回空字符串")
	}
	file := configFile()
	if file == "" {
		t.Error("configFile 返回空字符串")
	}
	if path.Dir(file) != dir {
		t.Errorf("configFile 路径不在 configDir 下: %s, %s", file, dir)
	}
}

// 测试 InitConfig 创建配置文件
func TestInitConfig(t *testing.T) {
	// 清理环境
	dir := configDir()
	file := configFile()
	_ = os.Remove(file)
	_ = os.RemoveAll(dir)

	err := InitConfig()
	if err != nil {
		t.Fatalf("InitConfig 执行失败: %v", err)
	}

	// 检查文件是否存在
	if _, err := os.Stat(file); err != nil {
		t.Errorf("配置文件未创建: %v", err)
	}
}

// 测试 LoadConfig 加载配置
func TestLoadConfig(t *testing.T) {
	// 先初始化
	err := InitConfig()
	if err != nil {
		t.Fatalf("InitConfig 执行失败: %v", err)
	}

	err = LoadConfig()
	if err != nil {
		t.Errorf("LoadConfig 执行失败: %v", err)
	}

	// 检查 GlobalConfig 字段
	if GlobalConfig == nil {
		t.Error("GlobalConfig 为空")
	}
	if !GlobalConfig.Init {
		t.Error("GlobalConfig.Init 应为 true")
	}
	if len(GlobalConfig.Repo) != 0 {
		t.Error("GlobalConfig.Repo 应为空")
	}
	if len(GlobalConfig.Project) != 0 {
		t.Error("GlobalConfig.Project 应为空")
	}
	if GlobalConfig.Workspace != "" {
		t.Error("GlobalConfig.Workspace 应为空字符串")
	}
}

// 测试 LoadConfig 处理空文件
func TestLoadConfig_EmptyFile(t *testing.T) {
	// 创建空文件
	file := configFile()
	dir := configDir()
	_ = os.MkdirAll(dir, 0755)
	f, err := os.Create(file)
	if err != nil {
		t.Fatalf("创建空配置文件失败: %v", err)
	}
	f.Close()

	err = LoadConfig()
	if err == nil {
		t.Error("空配置文件应返回错误")
	}
	if !errors.Is(err, errors.New("config file is empty")) && err.Error() != "config file is empty" {
		t.Errorf("期望错误 'config file is empty'，实际: %v", err)
	}
}
