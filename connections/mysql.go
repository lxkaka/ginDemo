package connections

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func MysqlConn() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("cant not connect to mysql: %s", err.Error()))
	}
	return db
}
