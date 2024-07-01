package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/turbolytics/flake/pkg/flake"
	"time"
)

func init() {
	var (
		workerID uint64
		count    int
	)

	var generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate a new 128-bit Flake ID",
		Long:  `Generates one or more 128-bit Flake IDs using the specified region and machine IDs.`,
		Run: func(cmd *cobra.Command, args []string) {
			fg, err := flake.NewGenerator(
				flake.GeneratorWithWorkerID(workerID),
			)
			if err != nil {
				panic(err)
			}
			for i := 0; i < count; i++ {
				id, err := fg.GenerateFlakeID()
				if err != nil {
					panic(err)
				}
				fmt.Printf(
					"id=%q timestamp=%q\n",
					id.String(),
					time.UnixMilli(int64(id.Timestamp)),
				)
			}
		},
	}

	// Add flags for the generate command
	generateCmd.Flags().Uint64VarP(&workerID, "worker", "r", 1, "Worker ID")
	generateCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of IDs to generate")

	// Add generate command to the root command
	rootCmd.AddCommand(generateCmd)
}
