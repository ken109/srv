package cmd

import (
	"github.com/fatih/color"
	"github.com/ken109/srv/util"
	"github.com/spf13/cobra"
	"os/exec"
)

var baseStartCmd = &cobra.Command{
	Use:   "staging",
	Short: "base: staging",
	Run: func(cmd *cobra.Command, args []string) {
		var dir = util.Pwd()
		util.Cd(brewPrefix + "/etc/srv/base")
		color.Green("Starting...")
		if err := exec.Command("docker-compose", "up", "-d", "--remove-orphans").Run(); err != nil {
			color.Red("Failed to staging.")
		} else {
			color.Green("Successfully started.")
		}
		util.Cd(dir)
	},
}

func init() {
	baseCmd.AddCommand(baseStartCmd)
}
