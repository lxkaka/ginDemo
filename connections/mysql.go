package connections

import (
	"fmt"
	"ginDemo/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func MysqlConn() *gorm.DB {
	dbInfo := fmt.Sprintf("%s:%s@(%s:%d)/test?charset=utf8&parseTime=True&loc=Local",
		utils.Settings.Mysql.User,
		utils.Settings.Mysql.Password,
		utils.Settings.Mysql.Host,
		utils.Settings.Mysql.Port)
	db, err := gorm.Open("mysql", dbInfo)
	if err != nil {
		panic(fmt.Sprintf("cant not connect to mysql: %s", err.Error()))
	}
	return db
}
