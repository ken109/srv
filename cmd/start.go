package cmd

import (
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start environment",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func start(framework string, project string) {
	emptyCheck()
	color.Green("Copying docker-compose.yml...")
	startCopyCompose(brewPrefix+"/etc/srv/"+framework+".yml", project)
	color.Green("Creating database...")
	createDB(project, config.Mysql.User, config.Mysql.Password)

	color.Green("Starting...")
	composeUp()
	color.Green("Completed.")
}

func emptyCheck() {
	var fileCount = 0
	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		fileCount++
		return nil
	})

	if fileCount > 1 {
		color.Red("Current directory is not empty.")
		os.Exit(1)
	}
}

func startCopyCompose(path string, project string) {
	ft, err := os.Open(path)
	if err != nil {
		color.Red("Could not read " + path)
		os.Exit(1)
	}
	defer ft.Close()
	b, _ := ioutil.ReadAll(ft)
	var compose = string(b)
	compose = strings.ReplaceAll(compose, "APP_NAME", project)
	compose = strings.ReplaceAll(compose, "HOST_NAME", project)
	fc, err := os.Create("./docker-compose.yml")
	if err != nil {
		color.Red("Could not read template")
		os.Exit(1)
	}
	defer fc.Close()
	fc.WriteString(compose)
}
