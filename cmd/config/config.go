/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package config_cmd

import (
	"fmt"
	"kongcli/cmd"
	"kongcli/internal/config"
	"log"

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
		if url == "" {
			log.Fatal("No URL given, exiting...")
		}
		if name == "" {
			log.Fatal("No name given, exiting...")
		}
		err := config.AddToConfig(name, url, user, password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Added %s to config\n", name)
	},
}

var deleteServer = &cobra.Command{
	Use:       "rm",
	Short:     "Deletes a server from config",
	ValidArgs: []string{"name"},
	Run: func(cmd *cobra.Command, args []string) {

		for _, arg := range args {
			if arg == "" {
				continue
			}
			err := config.DeleteFromConfig(arg)
			if err != nil {
				log.Printf("Error deleting %s: %v", arg, err)
				continue
			}
			fmt.Printf("Deleted %s from config\n", arg)
		}

		if len(args) == 0 && name == "" {
			log.Fatal("No names given to delete, exiting...")
		}
	},
}

var listServers = &cobra.Command{
	Use:   "ls",
	Short: "Lists all servers in config",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.ListServers()
		if err != nil {
			log.Fatal(err)
		}
	},
}

var setContext = &cobra.Command{
	Use:       "set",
	Short:     "Sets the current context",
	ValidArgs: []string{"name"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("No context name provided as argument, exiting...")
		}
		err := config.SetContext(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Set context to %s\n", args[0])
	},
}

func init() {
	cmd.RootCmd.AddCommand(configCmd)
	configCmd.AddCommand(addServer)
	configCmd.AddCommand(deleteServer)
	configCmd.AddCommand(listServers)
	configCmd.AddCommand(setContext)
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
