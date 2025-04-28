package gitlab

import (
	"develop-cmd/cmd/gitlab/doctor"
	"develop-cmd/cmd/gitlab/mr"

	"github.com/spf13/cobra"
)

func Glab() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "glab",
		Short: "gitlab 命令",
		Run: func(c *cobra.Command, args []string) {
			c.Help()
		},
	}

	cmd.AddCommand(doctor.Doctor())
	cmd.AddCommand(mr.Mr())
	return cmd
}
