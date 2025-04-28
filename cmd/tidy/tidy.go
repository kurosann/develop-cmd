package tidy

import (
	"context"
	"fmt"

	"develop-cmd/pkg/C"
	"develop-cmd/pkg/env"

	"github.com/spf13/cobra"
)

func Tidy() *cobra.Command {
	var tidyCmd = &cobra.Command{
		Use:   "tidy",
		Short: "格式化代码",
		Run: func(c *cobra.Command, args []string) {
			// 循环遍历当前目录下的所有bg-开头的文件夹 进入目录执行make tidy
			repos := env.GetRepoName()
			for _, repo := range repos {
				ctx := context.WithValue(c.Context(), "dir", repo)
				C.CmdOutByte(ctx, "make", "tidy")
				fmt.Println("repo", repo, "tidy success")
			}
		},
	}
	return tidyCmd
}
