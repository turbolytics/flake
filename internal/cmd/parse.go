package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/turbolytics/flake/pkg/flake"
	"log"
	"time"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse a Flake ID and display its components",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		idStr := args[0]
		id, err := flake.NewIDFromStr(idStr)
		if err != nil {
			log.Fatalf("Error parsing Flake ID: %v", err)
		}
		fmt.Printf(
			"Parsed ID:\nTimestamp: %s\nWorkerID: %d\nSequence: %d\n",
			time.UnixMilli(int64(id.Timestamp)), id.WorkerID, id.Sequence)
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
