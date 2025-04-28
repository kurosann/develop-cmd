package commit

import (
	"context"
	"develop-cmd/pkg/C"
	"develop-cmd/pkg/env"

	"github.com/spf13/cobra"
)

func Commit() *cobra.Command {
	return &cobra.Command{
		Use:   "commit",
		Short: "git commit 提交所有文件到版本控制中",
		Run: func(c *cobra.Command, args []string) {
			repos := env.GetRepoName()
			for _, repo := range repos {
				ctx := context.WithValue(c.Context(), "dir", repo)
				C.CmdOutByte(ctx, "git", "commit", "-m", args[0], "-a")
			}
		},
	}
}
