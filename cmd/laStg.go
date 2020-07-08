package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var laStgCmd = &cobra.Command{
	Use:     "la [project name]",
	Aliases: []string{"laravel"},
	Short:   "laravel: staging",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a project name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		staging("laravel", args[0])
	},
}

func init() {
	stgCmd.AddCommand(laStgCmd)
}
