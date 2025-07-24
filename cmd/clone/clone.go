package clone

import (
	"context"
	"develop-cmd/pkg/config"
	"develop-cmd/pkg/git"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func Clone() *cobra.Command {
	return &cobra.Command{
		Use:   "clone",
		Short: "clone 克隆项目",
		Run: func(cmd *cobra.Command, args []string) {
			repos := config.RepoPath()
			workspace := config.Workspace()
			fmt.Println("工作目录: ", workspace)

			ctx := context.WithValue(cmd.Context(), "dir", workspace)
			for _, repo := range repos {
				fmt.Println("克隆项目: ", repo)

				ctx_, cancel := context.WithTimeout(ctx, 3*time.Second)
				defer cancel()
				if err := git.Clone(ctx_, repo); err != nil {
					fmt.Printf("克隆项目 %s 失败: %v\n", repo, err)
					continue
				}
			}

			fmt.Println("克隆项目完成")
		},
	}
}
