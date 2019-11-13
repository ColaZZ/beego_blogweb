package models

import (
	"beego_blogweb/utils"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"strconv"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	CreateTime int64
}

func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	SetArticleRowsNum()
	return i, err
}

//插入一篇文章
func insertArticle(article Article) (int64, error) {
	return utils.ModifyDB("insert into article(title,tags,short,content,author,createtime) values(?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.CreateTime)
}

//根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	return QueryArticleWithPage(page, num)
}

func QueryArticleWithPage(page, num int) ([]Article, error) {
	sqlStr := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryArticlesWithCon(sqlStr)
}

func QueryArticlesWithCon(sqlStr string) ([]Article, error) {
	rows, err := utils.QueryDB(sqlStr)
	if err != nil {
		return nil, err
	}
	var articleList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		_ = rows.Scan(&id, &title, &tags, &short, &content, &author, createtime)
		art := Article{
			Id:         id,
			Title:      title,
			Tags:       tags,
			Short:      short,
			Content:    content,
			Author:     author,
			CreateTime: createtime,
		}
		articleList = append(articleList, art)
	}
	return articleList, nil
}

//存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var artcileRowsNum = 0

//只有首次获取行数的时候采取统计表里的行数
func GetArticleRowsNum() int {
	if artcileRowsNum == 0 {
		artcileRowsNum = QueryArticleRowNum()
	}
	return artcileRowsNum
}

//查询文章的总条数
func QueryArticleRowNum() int {
	row := utils.QueryRowDB("select count(id) from article")
	num := 0
	_ = row.Scan(&num)
	return num
}

// 设置页数
func SetArticleRowsNum() {
	artcileRowsNum = QueryArticleRowNum()
}

//获取id所对应的文章信息
func QueryArticleWithId(id int) Article {
	row := utils.QueryRowDB("select id,title,tags,short,content,author,creatime from article where id=" + strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	_ = row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{
		Id:         id,
		Title:      title,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Author:     author,
		CreateTime: createtime,
	}
	return art
}

//修改文章数据
func UpdateArticle(article Article) (int64, error) {
	return utils.ModifyDB("update article set title=?,tags=?,short=?,content=? where id=?",
		article.Title, article.Tags, article.Short, article.Content, article.Id)
}

//删除文章
func DeleteArticle(artID int) (int64, error) {
	i, err := deleteArticleWithArtId(artID)
	SetArticleRowsNum()
	return i, err
}

//按articleID删除文章
func deleteArticleWithArtId(artID int) (int64, error) {
	return utils.ModifyDB("delete from article where id=?", artID)
}

//查询所有的(参数)
func QueryArticleWithParam(param string) []string {
	rows, err := utils.QueryDB(fmt.Sprintf("select %s from article", param))
	if err != nil {
		log.Println(err)
	}

	var paramList []string
	for rows.Next() {
		tag := ""
		rows.Scan(&tag)
		paramList = append(paramList, tag)
	}
	return paramList
}
