package merge

import (
	"context"
	"develop-cmd/pkg/C"
	"develop-cmd/pkg/env"
	"develop-cmd/pkg/git"
	"fmt"

	"github.com/spf13/cobra"
)

func Merge() *cobra.Command {
	var mergeCmd = &cobra.Command{
		Use:   "merge",
		Short: "合并分支",
		Args:  cobra.ExactArgs(1),
		RunE: func(c *cobra.Command, args []string) error {
			repos := env.GetRepoName()
			// 检查所有分支是否存在
			for _, repo := range repos {
				ctx := context.WithValue(c.Context(), "dir", repo)
				_, err := C.CmdOutByte(ctx, "git", "branch", "--list", args[0])
				if err != nil {
					return fmt.Errorf("repo %s branch %s not found", repo, args[0])
				}
			}
			for _, repo := range repos {
				ctx := context.WithValue(c.Context(), "dir", repo)
				if err := git.Merge(ctx, args[0]); err != nil {
					return err
				}
				fmt.Println("repo", repo, "merge", args[0], "success")
			}
			return nil
		},
	}
	return mergeCmd
}
