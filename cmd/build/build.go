package build

import (
	"context"
	"develop-cmd/pkg/C"
	"develop-cmd/pkg/env"
	"fmt"

	"github.com/spf13/cobra"
)

func Build() *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "构建项目",
		RunE: func(cmd *cobra.Command, args []string) error {
			repos := env.GetRepoName()
			for _, repo := range repos {
				ctx := context.WithValue(cmd.Context(), "dir", repo)
				C.CmdOutByte(ctx, "make", "build")
				fmt.Println("repo", repo, "build success")
			}
			return nil
		},
	}
}
