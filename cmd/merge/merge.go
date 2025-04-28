package merge

import (
	"context"
	"develop-cmd/pkg/env"
	"develop-cmd/pkg/git"

	"github.com/spf13/cobra"
)

func Merge() *cobra.Command {
	var mergeCmd = &cobra.Command{
		Use:   "merge",
		Short: "合并分支",
		Args:  cobra.ExactArgs(1),
		RunE: func(c *cobra.Command, args []string) error {
			repos := env.GetRepoName()
			for _, repo := range repos {
				ctx := context.WithValue(c.Context(), "dir", repo)
				if err := git.Merge(ctx, args[0]); err != nil {
					return err
				}
			}
			return nil
		},
	}
	return mergeCmd
}
