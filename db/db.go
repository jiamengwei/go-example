package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var pool *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:root@/go_blog?parseTime=true")
	if err != nil {
		log.Panic("数据库连接初始化失败_", err)
	}
	pool = db
}

func Conn() *sql.DB {
	return pool
}
