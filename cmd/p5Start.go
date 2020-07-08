package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var p5StartCmd = &cobra.Command{
	Use:     "p5 [project name]",
	Aliases: []string{"php5"},
	Short:   "php5: start",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a project name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		start("php5", args[0])
	},
}

func init() {
	startCmd.AddCommand(p5StartCmd)
}
