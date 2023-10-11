package main

import (
	"log"
	"os"

	"github.com/stellayazilim/neptune_cms/config"
	"github.com/stellayazilim/neptune_cms/internal/rest"
)

func main() {

	// load env
	if err := config.InjectEnv(); err != nil {
		log.Fatal(err)
	}

	// init rest app
	r := rest.Rest()

	// start listen port
	r.Run(os.Getenv("NEPTUNE_REST_ADDR"))
}
