package pull

import (
	"context"
	"develop-cmd/pkg/env"
	"develop-cmd/pkg/git"
	"develop-cmd/pkg/logger"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

func Pull() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "从远程仓库拉取所有仓库的更新",
		RunE: func(cmd *cobra.Command, args []string) error {
			os.Setenv("GIT_PAGER", "cat")
			repos := env.GetRepoName()
			wg := sync.WaitGroup{}
			for _, repo := range repos {
				wg.Add(1)
				logger.Info("正在拉取仓库 %s 的更新...", repo)
				ctx := context.WithValue(cmd.Context(), "dir", repo)
				go func(repo string) {
					defer wg.Done()
					if err := git.Pull(ctx); err != nil {
						logger.Info("拉取仓库 %s 失败: %v", repo, err)
						return
					}
					logger.Info("仓库 %s 拉取成功", repo)
				}(repo)
			}
			wg.Wait()
			return nil
		},
	}
	return cmd
}
