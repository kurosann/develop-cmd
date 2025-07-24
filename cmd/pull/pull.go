package pull

import (
	"context"
	"develop-cmd/pkg/config"
	"develop-cmd/pkg/git"
	"fmt"

	"github.com/spf13/cobra"
)

func Pull() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "从远程仓库拉取所有仓库的更新",
		RunE: func(cmd *cobra.Command, args []string) error {
			paths := config.ProjectPath()

			for _, path := range paths {
				fmt.Println("拉取项目: ", path)
				ctx := context.WithValue(cmd.Context(), "dir", path)
				if err := git.Pull(ctx); err != nil {
					return err
				}
			}

			fmt.Println("拉取完成")
			return nil
		},
	}
	return cmd
}
