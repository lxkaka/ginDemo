package controllers

import (
	"ginDemo/process"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Credential struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Auth(c *gin.Context) {
	var cre Credential
	err := c.BindJSON(&cre)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var vp = process.VideoProcess{}
	author, err := vp.FindAuthor(cre.Username, cre.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	token, err := utils.GenerateToken(author)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
