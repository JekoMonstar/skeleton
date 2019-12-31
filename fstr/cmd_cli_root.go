package fstr

var Cmd_Cli_Root_Str = `package cli

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile = "conf/dev.toml"
	Cmd2    = ""

	RootCmd = &cobra.Command{
		Use:   "%s",
		Short: "",
		Long:  "",
	}
)

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println("root err:", err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", cfgFile, "config_path")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	// 配置自动更新
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config is change :%%s \n", e.String())
	})
	viper.WatchConfig()
}
`
