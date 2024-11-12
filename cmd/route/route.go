/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package route

import (
	"kongcli/cmd"

	"github.com/spf13/cobra"
)

// routeCmd represents the route command
var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "Manages routes",
	Long:  `Interacts with kong routes`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("route called")
	// },
}

func init() {
	cmd.RootCmd.AddCommand(routeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
