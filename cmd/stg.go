package cmd

import (
	"database/sql"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ken109/srv/util"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var stgCmd = &cobra.Command{
	Use:   "stg",
	Short: "Staging environment",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(stgCmd)
}

func staging(framework string, project string) {
	checkDir(project)
	deploy(project)
	color.Green("Copying docker-compose.yml...")
	util.Cd(home + "/staging/" + project)
	copyCompose(home+"/.srv/"+framework+".yml", project)

	color.Green("Creating database...")
	createDB(project, config.Mysql.User, config.Mysql.Password)

	color.Green("Starting...")
	composeUp()

	color.Green("Completed.")
}

func checkDir(project string) {
	if util.Exists(home + "/staging/" + project) {
		color.Green("Deleting old project...")
		util.Remove(home + "/staging/" + project)
	}
}

func deploy(project string) {
	color.Green("Deploying...")
	util.Mkdir(home + "/staging/" + project)
	if err := util.TryCommand("tar", "xf", "/tmp/srv-"+project+".tar.gz", "-C", home+"/staging/"+project); err != nil {
		color.Red("Could not deploy.")
		os.Exit(1)
	}
	util.Remove("/tmp/srv-" + project + ".tar.gz")
}

func copyCompose(path string, project string) {
	ft, err := os.Open(path)
	if err != nil {
		color.Red("Could not read " + path)
		os.Exit(1)
	}
	defer ft.Close()
	b, _ := ioutil.ReadAll(ft)
	var compose = string(b)
	compose = strings.ReplaceAll(compose, "APP_NAME", project)
	compose = strings.ReplaceAll(compose, "DOMAIN", config.Option.Domain)
	compose = strings.ReplaceAll(compose, "MAIL", config.Option.Mail)
	fc, err := os.Create("./docker-compose.yml")
	if err != nil {
		color.Red("Could not read template")
		os.Exit(1)
	}
	defer fc.Close()
	fc.WriteString(compose)
}

func createDB(project string, user string, password string) {
	db, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/")
	if err != nil {
		color.Red("データベースに接続できませんでした")
		os.Exit(1)
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS `" + project + "`")
	if err != nil {
		color.Red("データベースを作成できませんでした。")
		os.Exit(1)
	}
}

func composeUp() {
	err := exec.Command("docker-compose", "up", "-d", "--remove-orphans").Run()
	if err != nil {
		color.Red("Could not staging docker-compose")
		os.Exit(1)
	}
}
