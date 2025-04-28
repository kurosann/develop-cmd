package branch

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"develop-cmd/pkg/C"
	"develop-cmd/pkg/env"

	"github.com/spf13/cobra"
)

// 查看所有分支
func Branch() *cobra.Command {
	var current bool
	var keyword string
	var branchCmd = &cobra.Command{
		Use:   "branch",
		Short: "查看所有分支",
		RunE: func(c *cobra.Command, args []string) error {
			os.Setenv("GIT_PAGER", "cat")
			repos := env.GetRepoName()

			// 用于存储所有分支
			allBranches := make(map[string]bool)

			for _, repo := range repos {
				// 先进行 fetch
				ctx := context.WithValue(c.Context(), "dir", repo)
				var output []byte
				var err error
				if current {
					output, err = C.CmdOutByte(ctx, "git", "branch", "--show-current")
				} else {
					output, err = C.CmdOutByte(ctx, "git", "branch", "-a")
				}
				if err != nil {
					return err
				}

				// 处理每个分支
				branches := strings.Split(string(output), "\n")
				for _, b := range branches {
					b = strings.TrimSpace(b)
					if b != "" {
						allBranches[b] = true
					}
				}
			}
			sortls := make([]string, 0, len(allBranches))
			for b := range allBranches {
				sortls = append(sortls, b)
			}
			sort.Strings(sortls)

			// 打印所有不重复的分支，并根据关键字过滤
			for _, b := range sortls {
				if keyword == "" || strings.Contains(b, keyword) {
					fmt.Println(b)
				}
			}

			return nil
		},
	}
	branchCmd.Flags().StringVarP(&keyword, "keyword", "k", "", "过滤分支")
	branchCmd.Flags().BoolVarP(&current, "current", "c", false, "显示当前分支")
	return branchCmd
}
