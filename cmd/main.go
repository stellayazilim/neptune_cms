package main

import (
	"log"
	"os"

	"github.com/stellayazilim/neptune_cms/internal/interfaces/rest"
	"github.com/stellayazilim/neptune_cms/pkg/utils"
)

func main() {
	// init .env files
	utils.InjectEnv(utils.GetRootDir() + "/env/.env")

	// init rest interface
	if err := rest.NewRestWithHandlers().Run(os.Getenv("NEPTUNE_REST_ADDR")); err != nil {
		log.Fatal(err)
	}

}
