package utils

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

func InitMysql() {
	fmt.Println("InitMysql...")
	dsn := "root:0312@tcp(127.0.0.1:3306)/myblogweb"
	db, _ = sqlx.Connect("mysql", dsn)

	//db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)

	CreateTableWithUser()
	CreateTableWithArticle()
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

func QueryDB(sqlStr string) (*sqlx.Rows, error) {
	return db.Queryx(sqlStr)
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

//创建文章表
func CreateTableWithArticle() {
	sql := `create table if not exists article(
        id int(4) primary key auto_increment not null,
        title varchar(30),
        author varchar(20),
        tags varchar(30),
        short varchar(255),
        content longtext,
        createtime int(10)
        );`
	_, _ = ModifyDB(sql)
}
