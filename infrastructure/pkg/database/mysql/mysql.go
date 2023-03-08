package mysql

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/viafcccy/garden-be/global"
)

func NewDB() *gorm.DB {
	return initDB()
}

func initDB() *gorm.DB {
	c := global.Gconfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Username, c.Password, c.DbHost, c.DbPort, c.DbName)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true) //如果使用 gorm 来帮忙创建表时，这里填写 false 的话 gorm 会给表添加 s 后缀，填写 true 则不会
	db.LogMode(true)       //打印 sql 语句

	db.SetLogger(&MyLogger{})

	db.DB().SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有 sql 任务需要执行的连接数大于 20，超过的连接会被连接池关闭。
	db.DB().SetMaxOpenConns(100) //设置数据库连接池最大连接数
	db.DB().SetConnMaxLifetime(30 * time.Second)

	return db
}

type MyLogger struct {
}

func (logger *MyLogger) Print(values ...interface{}) {
	var (
		level  interface{}
		source string
		doTime interface{}
		sql    string
	)

	for i, val := range values {
		if i == 0 {
			level = val
		}
		if i == 1 {
			source = val.(string)
		}
		if i == 2 {
			doTime = val
		}
		if i == 3 {
			sql = val.(string)
		}
	}

	if level == "sql" {
		logStr := fmt.Sprintf("%s", doTime) + " " + source + " " + sql
		log.Println(logStr)
	}
}
