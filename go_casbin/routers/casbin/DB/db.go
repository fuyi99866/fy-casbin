package DB

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_casbin/pkg/logger"
)

//连接数据库
func MysqlTool() *gorm.DB {
	mysql, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/casbin_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logger.Error("connect to DB error")
		panic(err)
	}
	return mysql
}
