package main

import (
	"github.com/wisner23/locker-api/src/handlers"
	"github.com/wisner23/locker-api/src/models"

	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
)

var ORM orm.Ormer

func init() {
	models.ConnectToDb()
	ORM = models.GetOrmObject()
}

func main() {
	router := gin.Default()
	router.POST("/users", handlers.CreateUser)
	router.GET("/users", handlers.GetUsers)

	router.Run(":3000")
}
