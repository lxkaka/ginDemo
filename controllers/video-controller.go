package controllers

import (
	"ginDemo/models"
	"ginDemo/process"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VideoController struct {
}

var vp = process.VideoProcess{}

func (ctl *VideoController) Find(c *gin.Context) {
	title := c.Query("title")
	c.JSON(200, vp.Find(title))
}

func (ctl *VideoController) Save(c *gin.Context) {
	var video models.Video
	err := c.ShouldBindJSON(&video)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	userid, exists := c.Get("userid")
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	video.AuthorID, _ = userid.(uint64)
	video.Author = models.Author{}
	vp.Save(video)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (ctl *VideoController) Update(c *gin.Context) {
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
	vp.Update(video)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (ctl *VideoController) Delete(c *gin.Context) {
	var video models.Video

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return

	}
	video.ID = id
	vp.Delete(video)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
