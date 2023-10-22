package router

import (
	"ginapi-cleanarchi/common/di"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

func SetRoutes(engine *gin.Engine, db *xorm.Engine) {

	user := di.InitUser(db)
	engine.GET("/user", user.GetUser)

}
