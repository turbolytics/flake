package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flake",
	Short: "Flake generates 128-bit GUIDs based on Twitter Snowflake IDs",
	Long:  `Flake is a CLI tool and library for generating 128-bit globally unique identifiers (GUIDs) using a system inspired by Twitter's Snowflake IDs.`,
	Run: func(cmd *cobra.Command, args []string) {
		// This function is called when no subcommands are provided.
		// You can show the help message by default.
		fmt.Println("Flake CLI - use 'flake generate' to create a Flake ID")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
