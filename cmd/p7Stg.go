package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var p7StgCmd = &cobra.Command{
	Use:     "p7 [project name]",
	Aliases: []string{"php7"},
	Short:   "php7: staging",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a project name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		staging("php7", args[0])
	},
}

func init() {
	stgCmd.AddCommand(p7StgCmd)
}
