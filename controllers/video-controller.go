package controllers

import (
	"ginDemo/models"
	"ginDemo/process"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//type VideoController interface {
//	Find(u string) []models.Video
//	Save(c *gin.Context) error
//	Update(c *gin.Context) error
//	Delete(c *gin.Context) error
//}

func Find(c *gin.Context) {
	username := c.Query("username")
	c.JSON(200, process.Find(username))
}

func Save(c *gin.Context) {
	var video models.Video
	var author models.Author
	err := c.ShouldBindJSON(&video)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	author, err = process.FindAuthor(video.Author.Username)
	if err == nil {
		video.AuthorID = author.ID
		video.Author = models.Author{}
	}
	process.Save(video)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func Update(c *gin.Context) {
	var video models.Video
	err := c.ShouldBindJSON(&video)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	video.ID = id
	process.Update(video)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func Delete(c *gin.Context) {
	var video models.Video

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return

	}
	video.ID = id
	process.Delete(video)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
