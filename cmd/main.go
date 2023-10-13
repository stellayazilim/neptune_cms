package main

import (
	"fmt"
	"log"

	"github.com/stellayazilim/neptune_cms/pkg/models"
	"github.com/stellayazilim/neptune_cms/pkg/neptune/account"
	"github.com/stellayazilim/neptune_cms/pkg/storage/postgres"
	"github.com/stellayazilim/neptune_cms/pkg/utils"
)

func main() {

	// load env
	utils.InjectEnv()

	// init rest app
	// r := rest.Rest()

	// // start listen port
	// r.Run(os.Getenv("NEPTUNE_REST_ADDR"))

	ac := &models.Account{
		Email: "jhon@doe.com",
	}
	pg := postgres.NewPostgres()

	ar := account.AccountRepository(pg)

	err := ar.GetAccountByEmail(ac)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ac)

}
