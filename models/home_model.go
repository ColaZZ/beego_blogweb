package models

import (
	"beego_blogweb/utils"
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
	"strings"
)

type HomeBlockParam struct {
	Id         int
	Title      string
	Tags       []TagLink
	Short      string
	Content    string
	Author     string
	CreateTime string

	Link string

	UpdateLink string
	DeleteLink string

	IsLogin bool
}

type TagLink struct {
	TagName string
	TagUrl  string
}

type HomeFooterPageCode struct {
	HasPre   bool
	HasNext  bool
	ShowPage string
	PreLink  string
	NextLink string
}

func MakeHomeBlocks(artList []Article, isLogin bool) template.HTML {
	htmlHome := ""
	for _, art := range artList {
		homeParam := HomeBlockParam{}
		homeParam.Id = art.Id
		homeParam.Title = art.Title
		homeParam.Tags = createTagsLinks(art.Tags)
		homeParam.Short = art.Short
		homeParam.Content = art.Content
		homeParam.Author = art.Author
		homeParam.CreateTime = utils.SwitchTimeStampToData(art.CreateTime)
		homeParam.Link = "/article/" + strconv.Itoa(art.Id)
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		homeParam.IsLogin = isLogin

		t, _ := template.ParseFiles("views/block/home_block.html")
		buffer := bytes.Buffer{}
		_ = t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	fmt.Println("htmlHome-->", htmlHome)
	return template.HTML(htmlHome)
}

func createTagsLinks(tags string) []TagLink {
	var tagLink []TagLink
	tagsParam := strings.Split(tags, "&")
	for _, tag := range tagsParam {
		tagLink = append(tagLink, TagLink{tag, "/?tag=" + tags})
	}
	return tagLink
}

func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}
	//查询出总的条数
	num := GetArticleRowsNum()
	pageRow, _ := beego.AppConfig.Int("articleListPageNum")
	//计算总页数
	allPageNum := (num-1) / pageRow + 1

	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

	//当前页数小于等于1，name上一页的按钮不能点击
	if page < 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}

	//当前页数大于等于总页数，name下一页不能点击
	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}
	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.PreLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode
}
