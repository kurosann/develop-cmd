package status

import (
	"context"
	"develop-cmd/pkg/C"
	"develop-cmd/pkg/env"
	"fmt"

	"github.com/spf13/cobra"
)

func Status() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "查看当前分支状态",
		RunE: func(cmd *cobra.Command, args []string) error {
			repos := env.GetRepoName()
			for _, repo := range repos {
				ctx := context.WithValue(cmd.Context(), "dir", repo)
				status, err := C.CmdOutByte(ctx, "git", "status")
				if err != nil {
					return err
				}
				fmt.Println(repo, string(status))
			}
			return nil
		},
	}
}
