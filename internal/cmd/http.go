package cmd

import (
	"github.com/spf13/cobra"
	"github.com/turbolytics/flake/internal/flakehttp"
	"github.com/turbolytics/flake/pkg/flake"
	"log"
	"net/http"
	"strconv"
	"time"
)

func init() {
	var (
		workerID uint64
		port     int
	)

	var httpCmd = &cobra.Command{
		Use:   "http",
		Short: "Starts a Flake Generator HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			fg, err := flake.NewGenerator(
				flake.GeneratorWithWorkerID(workerID),
			)

			if err != nil {
				panic(err)
			}

			handlers := flakehttp.Handlers{
				FlakeGen: fg,
			}

			mux := http.NewServeMux()
			mux.HandleFunc("/generate", handlers.GenerateFlakeIDHandler)

			server := &http.Server{
				Addr:         ":" + strconv.Itoa(port),
				Handler:      mux,
				ReadTimeout:  1 * time.Second,
				WriteTimeout: 10 * time.Second,
			}

			log.Printf("Handlers listening on port %d...", port)
			log.Fatal(server.ListenAndServe())
		},
	}

	httpCmd.Flags().Uint64VarP(&workerID, "worker", "r", 1, "Worker ID")
	httpCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port for HTTP server (default is 8080)")

	rootCmd.AddCommand(httpCmd)
}
