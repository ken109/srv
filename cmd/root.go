package cmd

import (
	"fmt"
	"github.com/ken109/srv/util"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string
var config util.Config
var home string
var brewPrefix = util.GetOutput("brew", "--prefix")

var rootCmd = &cobra.Command{
	Use:   "srv",
	Short: "Build staging environment",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	initConfig()

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		fmt.Sprintf("config file (default is %s)", brewPrefix+"/etc/srv/config.yaml"))

	home, _ = homedir.Dir()
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile(brewPrefix + "/etc/srv/config.yml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Config file Read error")
		fmt.Println(err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Config file Unmarshal error")
		fmt.Println(err)
		os.Exit(1)
	}
}
