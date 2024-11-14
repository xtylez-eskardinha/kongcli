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
		fmt.Printf("Consumer created: %s\n", cons.ID)
	},
}

var deleteConsumer = &cobra.Command{
	Use:   "rm",
	Short: "Deletes a consumer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting consumer...")
		err := consumers.DeleteConsumer(cmd.Context(), client, consumer)
		if err != nil {
			log.Fatal("ERROR: Cannot delete consumer", err.Error())
		}
		fmt.Println("Deleted consumer")
	},
}

var listConsumers = &cobra.Command{
	Use:   "ls",
	Short: "Lists consumers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing consumers...")
		var consumerList []*kong.Consumer
		var err error
		if len(tags) > 0 {
			pointerTags := make([]*string, len(tags))
			for i, s := range tags {
				pointerTags[i] = &s
			}
			consumerList, err = consumers.ListConsumersFiltered(cmd.Context(), client, pointerTags)
		} else {
			consumerList, err = consumers.ListConsumers(cmd.Context(), client)
		}
		if err != nil {
			log.Fatal("ERROR: Cannot list consumers", err.Error())
		}
		for _, c := range consumerList {
			fmt.Printf("Consumer: %s\n", c.ID)
		}
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
	consumerCmd.AddCommand(deleteConsumer)
	consumerCmd.AddCommand(listConsumers)
	consumerCmd.PersistentFlags().StringSliceVarP(&tags, "tags", "t", nil, "Tags to use with the command")
	consumerCmd.PersistentFlags().StringVar(&consumer, "consumer-id", "", "Consumer ID to use with the command")
	consumerCmd.PersistentFlags().StringVar(&customid, "custom-id", "", "Custom ID to use with the command")
}
