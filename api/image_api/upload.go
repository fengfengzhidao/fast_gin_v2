package image_api

import (
	"fast_gin/global"
	"fast_gin/utils/find"
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
	"path"
	"path/filepath"
	"strings"
)

var whiteList = []string{
	".jpg",
	".jpeg",
	".png",
	".webp",
}

func (ImageApi) UploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		res.FailWithMsg("请选择文件", c)
		return
	}

	// 大小限制
	if fileHeader.Size > global.Config.Upload.Size*1024*1024 {
		res.FailWithMsg("上传文件过大", c)
		return
	}
	// 后缀判断
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))

	if !find.InList(whiteList, ext) {
		res.FailWithMsg("上传文件后缀非法", c)
		return
	}

	// 处理文件名重复

	fp := path.Join("uploads", global.Config.Upload.Dir, fileHeader.Filename)

	c.SaveUploadedFile(fileHeader, fp)

	res.Ok("/"+fp, "上传成功", c)
}
