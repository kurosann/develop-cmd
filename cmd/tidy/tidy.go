package tidy

import (
	"context"
	"fmt"
	"sync"

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
			wg := sync.WaitGroup{}
			for _, repo := range repos {
				wg.Add(1)
				ctx := context.WithValue(c.Context(), "dir", repo)
				go func(ctx context.Context) {
					defer wg.Done()
					if _, err := C.CmdOutByte(ctx, "make", "tidy"); err != nil {
						fmt.Println("repo", repo, "tidy failed", err)
					} else {
						fmt.Println("repo", repo, "tidy success")
					}
				}(ctx)
			}
			wg.Wait()
		},
	}
	return tidyCmd
}
