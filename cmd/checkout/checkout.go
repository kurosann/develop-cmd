package checkout

import (
	"context"
	"fmt"
	"sync"

	"develop-cmd/cmd/checkout/bugfix"
	"develop-cmd/cmd/checkout/feature"
	"develop-cmd/pkg/env"
	"develop-cmd/pkg/git"

	"github.com/spf13/cobra"
)

func Checkout() *cobra.Command {
	var branch string
	var checkoutCmd = &cobra.Command{
		Use:   "checkout",
		Short: "切换分支",
		RunE: func(c *cobra.Command, args []string) error {
			repos := env.GetRepoName()
			wg := sync.WaitGroup{}
			for _, repo := range repos {
				wg.Add(1)
				ctx := context.WithValue(c.Context(), "dir", repo)
				go func(ctx context.Context) {
					defer wg.Done()
					if err := git.HandleBranch(ctx, branch); err != nil {
						fmt.Println("repo", repo, "checkout", branch, "failed", err)
					} else {
						fmt.Println("repo", repo, "checkout", branch, "success")
					}
				}(ctx)
			}
			wg.Wait()
			return nil
		},
	}
	checkoutCmd.Flags().StringVarP(&branch, "branch", "b", "", "分支名称 (required)")
	checkoutCmd.MarkFlagRequired("branch")
	checkoutCmd.AddCommand(feature.Feature())
	checkoutCmd.AddCommand(bugfix.Bugfix())
	return checkoutCmd
}
