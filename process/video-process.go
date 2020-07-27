package process

import (
	"ginDemo/connections"
	"ginDemo/models"
)

//type VideoProcess interface {
//	Save(video models.Video)
//	Update(video models.Video)
//	Delete(video models.Video)
//	Find(username string) []models.Video
//}

func Save(video models.Video) {
	conn := connections.MysqlConn()
	defer conn.Close()
	conn.Create(&video)
}

func Update(video models.Video) {
	conn := connections.MysqlConn()
	defer conn.Close()
	conn.Model(&video).Update(&video)
}

func Delete(video models.Video) {
	conn := connections.MysqlConn()
	defer conn.Close()
	conn.Delete(&video)
}

func Find(username string) []models.Video {
	conn := connections.MysqlConn()
	defer conn.Close()
	conn.AutoMigrate(&models.Video{}, &models.Author{})
	var videos []models.Video
	conn.Set("gorm:auto_preload", true).Where("username=?", username).Find(&videos)
	return videos
}

func FindAuthor(username string) (models.Author, error) {
	conn := connections.MysqlConn()
	defer conn.Close()
	var author models.Author
	err := conn.Where("username=?", username).First(&author).Error
	if err != nil {
		return models.Author{}, err
	}
	return author, nil
}
