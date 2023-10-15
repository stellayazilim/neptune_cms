package main

import (
	"context"
	"fmt"

	token_entity "github.com/stellayazilim/neptune_cms/pkg/entities/token.entity"
	"github.com/stellayazilim/neptune_cms/pkg/repositories"
	"github.com/stellayazilim/neptune_cms/pkg/storage/postgres"
	"github.com/stellayazilim/neptune_cms/pkg/utils"
)

func main() {

	utils.InjectEnv(utils.GetRootDir() + "/env/.env")
	pg := postgres.NewPostgres()

	//ar := repositories.AccountRepository(pg)

	tr := repositories.TokenRepository(pg)

	// as := auth.New(ar, tr)

	// sdt := auth.SignInDto{
	// 	Email:    "jhon@doe.com",
	// 	Password: "1234",
	// }
	//as.SignIn(context.TODO(), &sdt)

	//tokens, _ := tr.Find(context.TODO())

	tid := token_entity.ID(uint64(23))
	tbid, _ := tr.FindByTokenId(context.TODO(), &tid)

	fmt.Println(tbid.Value, tbid.TokenStatus)
	tbid.TokenStatus = tbid.TokenStatus.INVALID()

	tr.Update(context.TODO(), tbid)
	fmt.Println(tbid.Value, tbid.TokenStatus)
}
