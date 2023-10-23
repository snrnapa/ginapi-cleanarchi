package main

import (
	"fmt"
	"ginapi-cleanarchi/router"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func main() {
	// ginのエンジンインスタンスを生成
	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// env.EnvLoad()
	// driverName := "mysql"
	// DsName := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
	// err := errors.New("")
	DsName := "root:root@tcp([127.0.0.1]:3306)/user_manage?charset=utf8mb4&parseTime=true"

	dbEngine, err := xorm.NewEngine("mysql", DsName)
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
