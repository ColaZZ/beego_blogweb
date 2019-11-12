package utils

import (
	_ "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sqlx.DB

func InitMysql() {
	fmt.Println("InitMysql...")
	dsn := "root:0312@tcp(127.0.0.1:3306)/myblogweb"
	db, _ = sqlx.Connect("mysql", dsn)

	str := MD5("11")
	fmt.Println(str)
	//db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)

	CreateTableWithUser()
}

func CreateTableWithUser() {
	sqlStr := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	_, _ = ModifyDB(sqlStr)
}

// 操作数据库（返回row affected)
// 执行sql的exec语句
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
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


