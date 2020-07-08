package cmd

import (
	"database/sql"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
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
	var host string
	color.Green("Copying docker-compose.yml...")
	host = copyCompose(home+"/..stg/"+framework+".yml", project)
	color.Green("Creating database...")
	createDB(project, config.Mysql.User, config.Mysql.Password)

	color.Green("Starting...")
	composeUp()
	color.Blue("URL: http://" + host)
	color.Green("Completed.")
}

func copyCompose(path string, project string) string {
	ft, err := os.Open(path)
	if err != nil {
		color.Red("Could not read " + path)
		os.Exit(1)
	}
	defer ft.Close()
	b, _ := ioutil.ReadAll(ft)
	var compose = string(b)
	var host string
	compose = strings.ReplaceAll(compose, "APP_NAME", project)
	compose = strings.ReplaceAll(compose, "HOST_NAME", project)
	host = project
	fc, err := os.Create("./docker-compose.yml")
	if err != nil {
		color.Red("Could not read template")
		os.Exit(1)
	}
	defer fc.Close()
	fc.WriteString(compose)

	return host
}

func createDB(project string, user string, password string) {
	db, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS `" + project + "`")
	if err != nil {
		color.Red("Could not create database")
		os.Exit(1)
	}
}

func composeUp() {
	err := exec.Command("docker-compose", "up", "-d", "--remove-orphans").Run()
	if err != nil {
		color.Red("Could not start docker-compose")
		os.Exit(1)
	}
}
