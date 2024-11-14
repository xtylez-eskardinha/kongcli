/*
Copyright © 2024 Miguel Angel Sanchez <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"kongcli/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var Ctx = context.Background()

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "kongcli",
	Short: "CLI application to manage Kong Admin API",
	Long:  `Kong Admin API manager from CLI with simple commands that allows interacting with services, routes, and more easily.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.ExecuteContext(Ctx)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kongcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cobra.OnInitialize(config.InitConfig)
}
