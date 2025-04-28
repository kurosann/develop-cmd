package remove

import (
	"context"
	"develop-cmd/pkg/C"
	"develop-cmd/pkg/env"
	"develop-cmd/pkg/logger"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// 删除分支
func Remove() *cobra.Command {
	return &cobra.Command{
		Use:   "remove",
		Short: "删除分支",
		RunE: func(c *cobra.Command, args []string) error {
			branch := args[0]
			if branch == "" {
				return fmt.Errorf("请输入分支名")
			}
			filterRepos := make([]string, 0)
			repos := env.GetRepoName()
			for _, repo := range repos {
				// 检查是否在当前分支
				ctx := context.WithValue(c.Context(), "dir", repo)
				output, err := C.CmdOutByte(ctx, "git", "branch", "--show-current")
				if err != nil {
					return err
				}
				ctx = context.WithValue(c.Context(), "dir", repo)
				output, err = C.CmdOutByte(ctx, "git", "branch", "-a")
				if err != nil {
					return err
				}
				filterBranches := make([]string, 0)
				branches := strings.Split(string(output), "\n")
				for _, b := range branches {
					if b != "" {
						if strings.Contains(string(b), branch) {
							if strings.Contains(string(b), "*") {
								fmt.Printf("当前分支为 %s，请先切换到其他分支\n", branch)
								os.Exit(1)
							}
							if len(filterBranches) == 0 {
								filterBranches = append(filterBranches, repo)
								filterRepos = append(filterRepos, repo)
							}
							filterBranches = append(filterBranches, b)
						}
					}
				}
				if len(filterBranches) != 0 {
					fmt.Println(strings.Join(filterBranches, " "))
				}
			}
			if len(filterRepos) == 0 {
				fmt.Printf("没有目标分支 %s\n", branch)
				return nil
			}
			logger.Info("目标仓库: %s", strings.Join(filterRepos, " "))
			// 提示用户是否删除
			fmt.Printf("是否删除以上仓库分支 %s (y/n):", branch)
			var confirm string
			fmt.Scanln(&confirm)
			if strings.ToLower(confirm) != "y" {
				return nil
			}
			for _, repo := range filterRepos {
				logger.Info("删除 %s 本地分支 %s", repo, branch)
				ctx := context.WithValue(c.Context(), "dir", repo)
				C.CmdOutByte(ctx, "git", "branch", "-d", branch)
				logger.Info("删除 %s 本地分支 %s 成功", repo, branch)
			}
			fmt.Println("删除成功")

			fmt.Print("是否推送(y/n):")
			fmt.Scanln(&confirm)
			if strings.ToLower(confirm) != "y" {
				return nil
			}
			for _, repo := range filterRepos {
				logger.Info("删除 %s 远程分支 %s", repo, branch)
				ctx := context.WithValue(c.Context(), "dir", repo)
				C.CmdOutByte(ctx, "git", "push", "origin", "--delete", branch)
				logger.Info("删除 %s 远程分支 %s 成功", repo, branch)
			}
			return nil
		},
	}
}
