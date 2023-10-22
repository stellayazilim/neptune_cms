package main

import (
	"log"
	"os"

	"github.com/stellayazilim/neptune_cms/internal/interfaces/rest"
	"github.com/stellayazilim/neptune_cms/pkg/common/utils"
	"github.com/stellayazilim/neptune_cms/pkg/storage/memory"
)

func main() {
	// init .env files
	utils.InjectEnv(utils.GetRootDir() + "/env/.env")

	// init memory storages
	memory.InitMemoryUser()

	// init rest interface
	if err := rest.NewRestWithHandlers().Run(os.Getenv("NEPTUNE_REST_ADDR")); err != nil {
		log.Fatal(err)
	}

}
