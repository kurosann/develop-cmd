package initialize

import (
	"develop-cmd/pkg/config"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// Initialize init .devctl 文件
func Initialize() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "init 初始化配置，文件位置：~/.devctl/config",
		Run: func(c *cobra.Command, args []string) {
			dir, err := os.UserHomeDir()
			if err != nil {
				panic(err)
			}

			devctlDir := path.Join(dir, ".devctl")
			if _, err := os.Stat(devctlDir); os.IsNotExist(err) {
				if err := os.Mkdir(devctlDir, os.ModePerm); err != nil {
					panic(err)
				}
			}

			filePath := path.Join(devctlDir, "config")
			file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err != nil {
				panic(err)
			}

			// 判断文件内容是否为空
			stat, err := file.Stat()
			if err != nil {
				panic(err)
			}
			if stat.Size() != 0 {
				fmt.Println("配置文件已存在且非空，无需初始化。")
				if err = file.Close(); err != nil {
					panic(err)
				}
				return
			}

			if _, err = file.Write(config.DefaultConfig()); err != nil {
				panic(err)
			}
			if err = file.Sync(); err != nil {
				panic(err)
			}
			if err = file.Close(); err != nil {
				panic(err)
			}

			tip := `
配置文件初始化 %s 完毕。
修改配置文件中，完成配置
Repo      - 项目 git 地址
Project   - 项目名称
Workspace - 项目工作目录
`
			fmt.Printf(tip, filePath)
		},
	}
}
