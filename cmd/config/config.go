/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package config_cmd

import (
	"fmt"
	"kongcli/cmd"
	"kongcli/internal/config"

	"github.com/spf13/cobra"
)

var name string
var user string
var password string
var url string

// routeCmd represents the route command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manages configs",
	Long:  `Interacts with kong configs`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("route called")
	// },
}

var addServer = &cobra.Command{
	Use:       "add",
	Short:     "Adds a server to config",
	ValidArgs: []string{"user", "pass", "url", "name"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding to config")
		if url == "" {
			panic("No URL given, exiting...")
		}
		if name == "" {
			panic("No name given, exiting...")
		}
		config.AddToConfig(name, url, user, password)
	},
}

func init() {
	cmd.RootCmd.AddCommand(configCmd)
	addServer.PersistentFlags().StringVarP(&name, "name", "n", "", "Name for the server")
	addServer.PersistentFlags().StringVarP(&user, "user", "u", "", "Username for the auth")
	addServer.PersistentFlags().StringVarP(&password, "pass", "p", "", "Password for the user auth")
	addServer.PersistentFlags().StringVarP(&url, "url", "U", "", "URL for the kong server")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
