package cmd

import (
	"database/sql"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop environment",
	Run: func(cmd *cobra.Command, args []string) {
		project := GetProjectName()
		color.Green(" Stopping " + project + "...")
		ComposeDown()

		if dropDb, err := cmd.PersistentFlags().GetBool("drop-db"); err == nil {
			if dropDb {
				color.Green("Dropping database...")
				DropDB(project, config.Mysql.User, config.Mysql.Password)
			}
		}

		color.Green("Completed.")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	stopCmd.PersistentFlags().Bool("drop-db", false, "Share in the office")
}

func GetProjectName() string {
	ft, err := os.Open("./docker-compose.yml")
	if err != nil {
		color.Red("Could not read docker-compose.yml")
		os.Exit(1)
	}
	defer ft.Close()
	b, _ := ioutil.ReadAll(ft)

	pt := regexp.MustCompile("services:\n (.*):")
	group := pt.FindSubmatch(b)
	return strings.TrimSpace(string(group[1]))
}

func ComposeDown() {
	err := exec.Command("docker-compose", "down").Run()
	if err != nil {
		color.Red("Could not stop docker-compose")
		os.Exit(1)
	}
}

func DropDB(project string, user string, password string) {
	db, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	_, err = db.Exec("DROP DATABASE IF EXISTS `" + project + "`")
	if err != nil {
		color.Red("Could not drop database")
		os.Exit(1)
	}
}
