package main

import (
	"os"

	"github.com/stellayazilim/neptune_cms/internal/rest"
	"github.com/stellayazilim/neptune_cms/pkg/utils"
)

func main() {

	// load env
	utils.InjectEnv()

	// init rest app
	r := rest.Rest()

	// start listen port
	r.Run(os.Getenv("NEPTUNE_REST_ADDR"))

}
