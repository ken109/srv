package cmd

import (
	"github.com/fatih/color"
	"github.com/ken109/stg/util"
	"github.com/spf13/cobra"
	"os/exec"
)

var baseStartCmd = &cobra.Command{
	Use:   "start",
	Short: "base: start",
	Run: func(cmd *cobra.Command, args []string) {
		var dir = util.Pwd()
		util.Cd(home + "/..stg/base")
		color.Green("Starting...")
		if err := exec.Command("docker-compose", "up", "-d", "--remove-orphans").Run(); err != nil {
			color.Red("Failed to start.")
		} else {
			color.Green("Successfully started.")
		}
		util.Cd(dir)
	},
}

func init() {
	baseCmd.AddCommand(baseStartCmd)
}
