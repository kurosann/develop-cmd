package checkout

import (
	"context"
	"fmt"

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
			for _, repo := range repos {
				ctx := context.WithValue(c.Context(), "dir", repo)
				if err := git.HandleBranch(ctx, branch); err != nil {
					return err
				}
				fmt.Println("repo", repo, "checkout", branch, "success")
			}
			return nil
		},
	}
	checkoutCmd.Flags().StringVarP(&branch, "branch", "b", "", "分支名称 (required)")
	checkoutCmd.MarkFlagRequired("branch")
	checkoutCmd.AddCommand(feature.Feature())
	return checkoutCmd
}
