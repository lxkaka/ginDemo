package controllers

import (
	"ginDemo/models"
	"ginDemo/process"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Credential struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var up = process.UserProcess{}

func Auth(c *gin.Context) {
	var cre Credential
	err := c.BindJSON(&cre)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	author, err := up.Find(cre.Username, cre.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	token, err := utils.GenerateToken(author)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func SignUp(c *gin.Context) {
	var user models.Author
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	up.Create(user)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
