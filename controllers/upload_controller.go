package controllers

import (
	"fmt"
	"gin_blogweb/models"
	"io"
	"os"
	"path/filepath"
	"time"
)

type UploadController struct {
	BaseController
}

func (c *UploadController) Post() {
	fileData, fileHeader, err := c.GetFile("upload")
	if err != nil {
		c.responseErr(err)
		return
	}

	now := time.Now()
	fileType := "other"
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == "jepg" {
		fileType = "img"
	}
	//文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		c.responseErr(err)
		return
	}

	//文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathStr := filepath.Join(fileDir, fileName)
	desFile, err := os.Create(filePathStr)
	if err != nil {
		c.responseErr(err)
		return
	}

	_, err = io.Copy(desFile, fileData)
	if err != nil {
		c.responseErr(err)
		return
	}

	if fileType == "img" {
		album := models.Album{
			Id:         0,
			FilePath:   filePathStr,
			Filename:   fileName,
			Status:     0,
			CreateTime: timeStamp,
		}
		_, _ = models.InsertAlbum(album)
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message":"上传成功"}
	c.ServeJSON()
}

func (c *UploadController) responseErr(err error) {
	c.Data["json"] = map[string]interface{}{"code": 0, "message": err}
	c.ServeJSON()
}
