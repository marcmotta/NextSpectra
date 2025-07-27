// cmd/nextspectra/main.go
package main

import (
	"flag"
	"log"
	"os"

	"nextspectra/internal/nextspectra"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := nextspectra.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
