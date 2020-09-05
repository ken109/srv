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
var brewPrefix = util.GetOutput("brew", "--prefix")
var home string

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
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "",
		"config file (default is "+brewPrefix+"/etc/srv/config.yml)")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, _ = homedir.Dir()

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
