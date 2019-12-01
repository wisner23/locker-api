package handlers

import (
	"github.com/wisner23/locker-api/src/models"
	"net/http"
	
	"github.com/gin-gonic/gin"
	_"github.com/astaxie/beego/orm"
)

func CreateUser(c *gin.Context) {

	ORM := models.GetOrmObject()
	var newUser models.Users
	c.BindJSON(&newUser)

	_, err := ORM.Insert(&newUser)

	if err == nil {
		c.JSON(
			http.StatusCreated,
			gin.H{ "email": newUser.Email},
		)
	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
	 }
}

func GetUsers(c *gin.Context) {
	var users []models.Users
	ORM := models.GetOrmObject()
	_, err := ORM.QueryTable("users").All(&users)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"users" : &users,
		},
	)
}
