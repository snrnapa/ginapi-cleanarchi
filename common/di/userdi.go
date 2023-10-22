package di

import (
	"ginapi-cleanarchi/handler"
	"ginapi-cleanarchi/interface/database"
	"ginapi-cleanarchi/usecase"

	"github.com/go-xorm/xorm"
)

func InitUser(db *xorm.Engine) handler.IUserHandler {
	r := database.NewUserRepository(db)
	s := usecase.NewUserCreate(r)
	return handler.NewUserHandler(s)

}
