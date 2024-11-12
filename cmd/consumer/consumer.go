/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package route

import (
	"fmt"
	"kongcli/cmd"
	"kongcli/internal/config"
	"kongcli/internal/consumers"

	"log"

	"github.com/kong/go-kong/kong"
	"github.com/spf13/cobra"
)

// kong Client
var client *kong.Client

// Common arg vars
var consumer string
var tags []string
var customid string

// consumerCmd represents the consumer command
var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Manages consumers",
	Long:  `Interacts with kong consumers`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("consumer called")
	// },
}

// Add subcommand for consumer
var addConsumer = &cobra.Command{
	Use:       "add",
	Short:     "Adds a consumer",
	ValidArgs: []string{"user", "customid", "tags"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating customer...")
		pointerTags := make([]*string, len(tags))
		for i, s := range tags {
			pointerTags[i] = &s
		}
		if consumer == "" {
			log.Fatal("No consumer name given")
		}
		cons, err := consumers.AddConsumer(cmd.Context(), client, consumer, customid, pointerTags)
		if err != nil {
			log.Fatal("ERROR: Cannot create consumer", err.Error())
		}
		fmt.Println("Lets go!", cons.CreatedAt)
	},
}

func init() {
	var err error
	client, err = config.CreateClient("https://PEPEGRILLO.com")
	if err != nil {
		log.Fatal("Unable to initialize kong client", err.Error())
	}
	cmd.RootCmd.AddCommand(consumerCmd)
	consumerCmd.AddCommand(addConsumer)
	consumerCmd.PersistentFlags().StringSliceVarP(&tags, "tags", "t", nil, "Tags to use with the command")
	consumerCmd.PersistentFlags().StringVar(&consumer, "consumer-id", "", "Consumer ID to use with the command")
	consumerCmd.PersistentFlags().StringVar(&customid, "custom-id", "", "Custom ID to use with the command")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
