package process

import (
	"fmt"
	"ginDemo/connections"
	"ginDemo/models"
)

type VideoProcess struct {
}

func (p *VideoProcess) Save(video models.Video) {
	conn := connections.MysqlConn()
	defer conn.Close()
	conn.Create(&video)
}

func (p *VideoProcess) Update(video models.Video) {
	conn := connections.MysqlConn()
	defer conn.Close()
	conn.Model(&video).Update(&video)
}

func (p *VideoProcess) Delete(video models.Video) {
	conn := connections.MysqlConn()
	defer conn.Close()
	conn.Delete(&video)
}

func (p *VideoProcess) Find(title string) []models.Video {
	conn := connections.MysqlConn()
	defer conn.Close()
	conn.AutoMigrate(&models.Video{}, &models.Author{})
	var videos []models.Video
	conn.Set("gorm:auto_preload", true).Where("title like ?", fmt.Sprintf("%s%%", title)).Find(&videos)
	return videos
}

func (p *VideoProcess) FindAuthor(username string, password string) (models.Author, error) {
	conn := connections.MysqlConn()
	defer conn.Close()
	var author models.Author
	err := conn.First(&author, "username=? and password=?", username, password).Error
	if err != nil {
		return models.Author{}, err
	}
	return author, nil
}
