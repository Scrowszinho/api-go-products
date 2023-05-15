package database

// import (
// 	"testing"

// 	"github.com/Scrowszinho/api-go-products/configs"
// 	"github.com/Scrowszinho/api-go-products/internal/entity"
// 	"github.com/Scrowszinho/api-go-products/migrations"
// )

// func TestCreateBet(t *testing.T) {
// 	configs.ConnectGorm()
// 	db := configs.GetDB()
// 	migrations.MigrateTable()
// 	betsDB := NewBet(db)
// 	userDB := NewUser(db)
// 	user, err := userDB.FindByEmailOrNickname("gustavo@gmail.com")
// 	if err != nil {
// 		panic(err)
// 	}
// 	bets, err := entity.NewBet(user, )
// }
