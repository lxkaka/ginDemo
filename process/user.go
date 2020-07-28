package process

import (
	"ginDemo/connections"
	"ginDemo/models"
)

type UserProcess struct {
}

func (p *UserProcess) Find(username string, password string) (models.Author, error) {
	conn := connections.MysqlConn()
	defer conn.Close()
	var author models.Author
	err := conn.First(&author, "username=? and password=?", username, password).Error
	if err != nil {
		return models.Author{}, err
	}
	return author, nil
}

func (p *UserProcess) Create(user models.Author) {
	conn := connections.MysqlConn()
	defer conn.Close()
	conn.Create(&user)
}
