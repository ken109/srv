package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var djStgCmd = &cobra.Command{
	Use:     "dj [project name]",
	Aliases: []string{"django"},
	Short:   "django: staging",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a project name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		staging("django", args[0])
	},
}

func init() {
	stgCmd.AddCommand(djStgCmd)
}
