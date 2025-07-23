package main

import (
	"develop-cmd/cmd/add"
	"develop-cmd/cmd/branch"
	"develop-cmd/cmd/build"
	"develop-cmd/cmd/checkout"
	"develop-cmd/cmd/commit"
	"develop-cmd/cmd/fetch"
	"develop-cmd/cmd/gitlab"
	"develop-cmd/cmd/initialize"
	"develop-cmd/cmd/merge"
	"develop-cmd/cmd/pull"
	"develop-cmd/cmd/push"
	"develop-cmd/cmd/remove"
	"develop-cmd/cmd/status"
	"develop-cmd/cmd/tidy"
	"develop-cmd/pkg/env"
	"develop-cmd/pkg/logger"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var verbose bool
	var debug bool
	var rootCmd = &cobra.Command{
		Use: "devctl",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// 查看当前目录
			dir, _ := os.Getwd()
			repoNames := env.GetRepoName()

			for _, repoName := range repoNames {
				// 路径内是否包含 repoName 切换到根目录
				if strings.Contains(dir, repoName) {
					workspace := strings.SplitN(dir, repoName, 2)[0]
					os.Chdir(workspace)
					break
				}
			}

			logger.Init(verbose, debug)
		},
	}
	rootCmd.PersistentFlags().BoolVarP(&verbose, "v", "", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&debug, "vv", "", false, "debug output")
	rootCmd.AddCommand(initialize.Initialize())
	rootCmd.AddCommand(branch.Branch())
	rootCmd.AddCommand(checkout.Checkout())
	rootCmd.AddCommand(tidy.Tidy())
	rootCmd.AddCommand(push.Push())
	rootCmd.AddCommand(pull.Pull())
	rootCmd.AddCommand(build.Build())
	rootCmd.AddCommand(status.Status())
	rootCmd.AddCommand(remove.Remove())
	rootCmd.AddCommand(fetch.Fetch())
	rootCmd.AddCommand(add.Add())
	rootCmd.AddCommand(commit.Commit())
	rootCmd.AddCommand(gitlab.Glab())
	rootCmd.AddCommand(merge.Merge())
	rootCmd.Execute()
}
