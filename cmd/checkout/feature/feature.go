package feature

import (
	"context"
	"fmt"
	"time"

	"develop-cmd/pkg/env"
	"develop-cmd/pkg/git"

	"github.com/spf13/cobra"
)

func Feature() *cobra.Command {
	featureCmd := &cobra.Command{
		Use:   "feature",
		Short: "创建 feature 分支",
		RunE: func(c *cobra.Command, args []string) error {
			repos := env.GetRepoName()

			// 获取 git 用户名
			username, err := git.GetUserName(c.Context())
			if err != nil {
				return fmt.Errorf("get git username failed: %v", err)
			}
			for _, repo := range repos {
				ctx := context.WithValue(c.Context(), "dir", repo)

				// 获取当前日期
				date := time.Now().Format("20060102")

				// 构建分支名
				branch := fmt.Sprintf("feature/%s/%s_%s", username, date, args[0])

				// 处理分支
				if err := git.HandleBranch(ctx, branch); err != nil {
					return err
				}
				fmt.Println("repo", repo, "checkout", branch, "success")
			}
			return nil
		},
	}
	return featureCmd
}
