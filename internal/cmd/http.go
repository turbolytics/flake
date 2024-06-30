package cmd

import (
	"github.com/spf13/cobra"
	"github.com/turbolytics/flake/internal/flakehttp"
	"github.com/turbolytics/flake/pkg/flake"
	"log"
	"net/http"
	"strconv"
)

func init() {
	var (
		regionID  uint16
		machineID uint16
		port      int
	)

	var httpCmd = &cobra.Command{
		Use:   "http",
		Short: "Starts a Flake Generator HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			fg, err := flake.NewGenerator(
				flake.GeneratorWithRegionID(regionID),
				flake.GeneratorWithMachineID(machineID),
			)

			if err != nil {
				panic(err)
			}

			handlers := flakehttp.Handlers{
				FlakeGen: fg,
			}

			http.HandleFunc("/generate", handlers.GenerateFlakeIDHandler)

			log.Printf("Handlers listening on port %d...", port)
			log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
		},
	}

	httpCmd.Flags().Uint16VarP(&regionID, "region", "r", 1, "Region ID")
	httpCmd.Flags().Uint16VarP(&machineID, "machine", "m", 1, "Machine ID")
	httpCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port for HTTP server (default is 8080)")

	rootCmd.AddCommand(httpCmd)
}
