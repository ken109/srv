package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var djStartCmd = &cobra.Command{
	Use:     "dj [project name]",
	Aliases: []string{"django"},
	Short:   "django: start",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a project name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		start("django", args[0])
	},
}

func init() {
	startCmd.AddCommand(djStartCmd)
}
