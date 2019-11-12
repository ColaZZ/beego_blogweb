package models

import (
	"beego_blogweb/utils"
	"fmt"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int
	Createtime int64
}

// 插入用户
func InertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username,pasword,status,creattime values(?,?,?,?)",
		user.Username, user.Password, user.Status, user.Createtime)
}

// 按条件查询,返回id
func QueryUserWightCon(con string) int {
	sqlStr := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sqlStr)
	row := utils.QueryRowDB(sqlStr)
	id := 0
	_ = row.Scan(&id)
	return id
}

// 根据用户名查询ID
func QueryUserWithUsername(username string) int {
	sqlStr := fmt.Sprintf("where username = %s", username)
	return QueryUserWightCon(sqlStr)
}
