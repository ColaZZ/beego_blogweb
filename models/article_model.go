package models

import (
	"beego_blogweb/utils"
	"fmt"
	"github.com/astaxie/beego"
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
		_ := rows.Scan(&id, &title, &tags, &short, &content, &author, createtime)
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
