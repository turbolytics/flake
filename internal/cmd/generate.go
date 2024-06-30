package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/turbolytics/flake/pkg/flake"
)

func init() {
	var (
		regionID  uint16
		machineID uint16
		count     int
	)

	var generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate a new 128-bit Flake ID",
		Long:  `Generates one or more 128-bit Flake IDs using the specified region and machine IDs.`,
		Run: func(cmd *cobra.Command, args []string) {
			fg, err := flake.NewGenerator(
				flake.GeneratorWithRegionID(regionID),
				flake.GeneratorWithMachineID(machineID),
			)
			if err != nil {
				panic(err)
			}
			for i := 0; i < count; i++ {
				id, err := fg.GenerateFlakeID()
				if err != nil {
					panic(err)
				}
				fmt.Println(id.String())
			}
		},
	}

	// Add flags for the generate command
	generateCmd.Flags().Uint16VarP(&regionID, "region", "r", 1, "Region ID")
	generateCmd.Flags().Uint16VarP(&machineID, "machine", "m", 1, "Machine ID")
	generateCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of IDs to generate")

	// Add generate command to the root command
	rootCmd.AddCommand(generateCmd)
}
