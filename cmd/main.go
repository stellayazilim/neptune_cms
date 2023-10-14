package main

import (
	"fmt"
	"os"

	"github.com/stellayazilim/neptune_cms/internal/rest"
	"github.com/stellayazilim/neptune_cms/pkg/neptune/account"
	"github.com/stellayazilim/neptune_cms/pkg/neptune/auth"
	"github.com/stellayazilim/neptune_cms/pkg/storage/postgres"
	"github.com/stellayazilim/neptune_cms/pkg/utils"
)

func main() {

	// load env
	utils.InjectEnv(utils.GetRootDir() + "/env/.env")

	// connect postgres
	pg := postgres.NewPostgres()

	//*** account module ****//
	//***				 ****//
	// account repository
	accountRepository := account.AccountRepository(pg)
	//***				 ***//

	//***  auth module ****//
	// auth service
	authService := auth.AuthService(accountRepository, auth.AuthHelper())
	//***  			   ***//

	// init rest app
	r := rest.Rest(authService)

	// start listen port
	if err := r.Run(os.Getenv("NEPTUNE_REST_ADDR")); err != nil {
		fmt.Println(err)
	}

}
