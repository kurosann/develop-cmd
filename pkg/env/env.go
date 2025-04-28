package env

import (
	"os"
	"strings"
)

var (
	RepoAddress = []string{
		"ssh://git@imageharbor.xyz:39423/bitguys/bg-agent.git",
		"ssh://git@imageharbor.xyz:39423/bitguys/bg-component.git",
		"ssh://git@imageharbor.xyz:39423/bitguys/bg-controller.git",
		"ssh://git@imageharbor.xyz:39423/bitguys/bg-framework.git",
		"ssh://git@imageharbor.xyz:39423/bitguys/bg-internal-sdk.git",
		"ssh://git@imageharbor.xyz:39423/bitguys/bitguys-proxy.git",
		"ssh://git@imageharbor.xyz:39423/bitguys/bitguys-mdw-sdk.git",
	}

	// 环境变量名称
	EnvRepoAddress    = "REPO_ADDRESS"
	EnvAddRepoAddress = "ADD_REPO_ADDRESS"
)

// getRepoAddressFromEnv 从环境变量获取仓库地址
func getRepoAddressFromEnv(envVar string, defaultValue []string) []string {
	if value := os.Getenv(envVar); value != "" {
		return strings.Split(value, ",")
	}
	if value := os.Getenv(EnvAddRepoAddress); value != "" {
		return append(defaultValue, strings.Split(value, ",")...)
	}
	return defaultValue
}

func GetRepoAddress() []string {
	return getRepoAddressFromEnv(EnvRepoAddress, RepoAddress)
}

func GetRepoName() []string {
	repoAddress := GetRepoAddress()
	repoName := []string{}
	for _, address := range repoAddress {
		// 获取 repo 名称 .git前 /后面的内容
		parts := strings.Split(address, "/")
		name := strings.Split(parts[len(parts)-1], ".git")[0]
		repoName = append(repoName, name)
	}
	return repoName
}
