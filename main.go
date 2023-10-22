package main

import (
	"errors"
	"fmt"
	"ginapi-cleanarchi/common/env"
	"ginapi-cleanarchi/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func main() {
	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	env.EnvLoad()

	driverName := "mysql"
	DsName := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
	err := errors.New("")
	dbEngine, err := xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	dbEngine.ShowSQL(true)
	dbEngine.SetMaxOpenConns(2)
	dbEngine.SetMapper(core.GonicMapper{})
	fmt.Println("init data base ok")

	router.SetRoutes(engine, dbEngine)

	engine.Run(":8080")
}
