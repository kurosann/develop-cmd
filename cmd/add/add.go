package add

import (
	"context"
	"develop-cmd/pkg/C"
	"develop-cmd/pkg/env"

	"github.com/spf13/cobra"
)

// git add
func Add() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "git add 添加所有文件到版本控制中",
		Run: func(c *cobra.Command, args []string) {
			repos := env.GetRepoName()
			for _, repo := range repos {
				ctx := context.WithValue(c.Context(), "dir", repo)
				C.CmdOutByte(ctx, "git", "add", ".")
			}
		},
	}
}
