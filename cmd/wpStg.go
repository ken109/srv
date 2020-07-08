package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var wpStgCmd = &cobra.Command{
	Use:     "wp [project name]",
	Aliases: []string{"wordpress"},
	Short:   "wordpress: staging",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a project name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		staging("wordpress", args[0])
	},
}

func init() {
	stgCmd.AddCommand(wpStgCmd)
}
