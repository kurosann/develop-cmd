package fetch

import (
	"context"
	"develop-cmd/pkg/C"
	"develop-cmd/pkg/env"
	"sync"

	"github.com/spf13/cobra"
)

func Fetch() *cobra.Command {
	return &cobra.Command{
		Use:   "fetch",
		Short: "git fetch",
		RunE: func(c *cobra.Command, args []string) error {
			repos := env.GetRepoName()
			wg := sync.WaitGroup{}
			for _, repo := range repos {
				wg.Add(1)
				ctx := context.WithValue(c.Context(), "dir", repo)
				go func(repo string) {
					defer wg.Done()
					C.CmdOutByte(ctx, "git", "fetch")
				}(repo)
			}
			wg.Wait()
			return nil
		},
	}
}
