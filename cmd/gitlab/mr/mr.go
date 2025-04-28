package mr

import (
	"context"
	"develop-cmd/pkg/C"
	"develop-cmd/pkg/env"
	"fmt"
	"strings"

	"github.com/avast/retry-go"
	"github.com/spf13/cobra"
)

func Mr() *cobra.Command {
	var description string
	var assignee string
	var removeSource bool
	cmd := &cobra.Command{
		Use:   "mr",
		Short: "gitlab mr 命令 args: target-branch title",
		Example: `
  devctl glab mr develop "mr title" // 创建mr
  devctl glab mr develop "mr title" -d "mr description" // 创建mr并指定描述
  devctl glab mr develop "mr title" -a "mr assignee" // 创建mr并指定合并人
  devctl glab mr develop "mr title" -d "mr description" -a "mr assignee" // 创建mr并指定描述和合并人
		`,
		Args: cobra.MinimumNArgs(2),
		Run: func(c *cobra.Command, args []string) {
			repos := env.GetRepoName()
			for _, repo := range repos {
				// 检查是否有提交
				ctx := context.WithValue(context.Background(), "dir", repo)
				bs, err := C.CmdOutByte(ctx, "git", "status")
				if err != nil {
					fmt.Println(err)
					return
				}
				// 若没有提交代码跳过mr
				// 有提交代码则创建mr
				if strings.Contains(string(bs), "nothing to commit") {
					// 检查当前分支和目标分支之间的差异
					diffCtx := context.WithValue(ctx, "dir", repo)
					_, err := C.CmdOutByte(diffCtx, "git", "diff", "--quiet", args[0]+"..HEAD")
					if err != nil {
						if err := retry.Do(func() error {
							args := []string{"mr", "create", "-y", "--target-branch", args[0], "--title", args[1]}
							if removeSource {
								args = append(args, "--remove-source-branch")
							}
							args = append(args, "-d", description)
							if assignee != "" {
								args = append(args, "-a", assignee)
							}
							// 如果有差异，则创建MR
							bs, err := C.CmdOutByte(ctx, "glab", args...)
							if err != nil {
								fmt.Println(err)
								if strings.Contains(err.Error(), "409") {
									fmt.Printf("%s 仓库MR已存在，跳过创建\n", repo)
									return nil
								}
								return err
							}
							fmt.Println(string(bs))
							return nil
						}, retry.OnRetry(func(n uint, err error) {
							fmt.Printf(" %s 仓库创建MR失败，重试第%d次\n", repo, n+1)
						}), retry.Attempts(3)); err != nil {
							fmt.Println(err)
						}
					} else {
						fmt.Printf(" %s 仓库当前分支与目标分支没有差异，跳过MR\n", repo)
					}
				} else {
					fmt.Printf(" %s 仓库有未提交的文件，请先提交\n", repo)
				}
			}
		},
	}

	cmd.Flags().StringVarP(&description, "description", "d", "", "描述")
	cmd.Flags().StringVarP(&assignee, "assignee", "a", "", "指定合并人")
	cmd.Flags().BoolVarP(&removeSource, "remove-source", "r", false, "删除源分支")
	return cmd
}
