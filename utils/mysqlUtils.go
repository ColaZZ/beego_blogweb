package utils

import (
	_ "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

func InitMysql() {
	fmt.Println("init mysql...")
	if db == nil {
		db, _ := sqlx.Open("mysql", "root:0312@tcp(127.0.0.1:3306)/myblogweb")
		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(16)
		CreateTableWithUser()
	}
}

// 操作数据库（返回row affected)
func ModifyDB(sqlStr string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sqlStr)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func QueryRowDB(sqlStr string) *sqlx.Row {
	return db.QueryRowx(sqlStr)
}

// 创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	_, _ = ModifyDB(sql)
}

