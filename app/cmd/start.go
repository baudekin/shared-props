/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	shared_props "github.com/baudekin/shared-props"
	"github.com/spf13/cobra"
	"sync"
)

var cancel context.CancelFunc

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		var ctx context.Context
		// Note cancel can only be accessed from cmd package and
		// should only be called by stop
		ctx, cancel = context.WithCancel(context.Background())
		var wg sync.WaitGroup
		wg.Add(2)
		var propchan = make(chan map[string]string)
		// Start Property Stream
		go func() {
			defer wg.Done()
			defer cancel()
			shared_props.StartUpdating(ctx, propchan)
		}()
		// Respond to Property Stream Changes
		go func() {
			defer wg.Done()
			defer cancel()
			shared_props.StartMonitoring(ctx, propchan)
		}()
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
