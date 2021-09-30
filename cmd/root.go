package cmd

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte("index page"))
			if err != nil {
				log.Panicf("error while writing response: %v", err)
			}
		})

		log.Fatalln(http.ListenAndServe(":1234", nil))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
