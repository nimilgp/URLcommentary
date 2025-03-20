package main

import (
	"log"

	"github.com/nimilgp/URLcommentary/cmd/api"
)

func main() {
	// server := api.GetAPIServer(":3333")
	if err := api.GetAPIServer(":3333").Run(); err != nil {
		log.Fatal(err)
	}
}
