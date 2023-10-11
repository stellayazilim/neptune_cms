package main

import (
	"os"

	"github.com/stellayazilim/neptune_cms/internal/rest"
	"github.com/stellayazilim/neptune_cms/pkg/utils/config"
)

func main() {

	// load env
	config.InjectEnv()

	// init rest app
	r := rest.Rest()

	// start listen port
	r.Run(os.Getenv("NEPTUNE_REST_ADDR"))
}
