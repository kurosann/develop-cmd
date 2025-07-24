package initialize

import (
	"develop-cmd/pkg/config"
	"fmt"

	"github.com/spf13/cobra"
)

// Initialize init .devctl 文件
func Initialize() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "init 初始化配置，文件位置：" + config.ConfigFile(),
		Run: func(c *cobra.Command, args []string) {
			if config.GlobalConfig.Init {
				fmt.Println("配置文件已初始化")
				return
			}

			if err := config.InitConfig(); err != nil {
				fmt.Println("配置文件初始化失败", err)
				return
			}

			tip := `
配置文件初始化 %s 完毕。
修改配置文件中，完成配置
Repo      - 项目 git 地址
Project   - 项目名称
Workspace - 项目工作目录
`
			fmt.Printf(tip, config.ConfigFile())
		},
	}
}
