package cmd

import (
	"github.com/fatih/color"
	"github.com/ken109/srv/util"
	"github.com/spf13/cobra"
	"os/exec"
)

var baseStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "base: stop",
	Run: func(cmd *cobra.Command, args []string) {
		var dir = util.Pwd()
		util.Cd(brewPrefix + "/etc/srv/base")
		color.Green("Stopping...")
		if err := exec.Command("docker-compose", "down").Run(); err != nil {
			color.Red("Failed to stop.")
		} else {
			color.Green("Successfully stopped.")
		}
		util.Cd(dir)
	},
}

func init() {
	baseCmd.AddCommand(baseStopCmd)
}
