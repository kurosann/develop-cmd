package push

import (
	"context"
	"develop-cmd/pkg/env"
	"develop-cmd/pkg/git"
	"develop-cmd/pkg/logger"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

func Push() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "push",
		Short: "推送所有仓库的当前分支到远程仓库",
		RunE: func(cmd *cobra.Command, args []string) error {
			os.Setenv("GIT_PAGER", "cat")
			repos := env.GetRepoName()
			wg := sync.WaitGroup{}
			for _, repo := range repos {
				wg.Add(1)
				logger.Info("正在推送仓库 %s 的更改...", repo)
				ctx := context.WithValue(cmd.Context(), "dir", repo)
				go func(repo string) {
					defer wg.Done()
					if err := git.Push(ctx, ""); err != nil {
						logger.Info("推送仓库 %s 失败: %v", repo, err)
						return
					}
					logger.Info("仓库 %s 推送成功", repo)
				}(repo)
			}
			wg.Wait()
			return nil
		},
	}
	return cmd
}
