package main

import (
	"log"
	"os"

	"github.com/k2t0f12d/prometheus-pingdom-exporter/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
