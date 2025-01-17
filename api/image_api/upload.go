package image_api

import (
	"fast_gin/global"
	"fast_gin/utils/find"
	"fast_gin/utils/md5"
	"fast_gin/utils/random"
	"fast_gin/utils/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
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
	// uploads/images/xx.jpg
	// uploads/images/xx_7hf.jpg
	fp := path.Join("uploads", global.Config.Upload.Dir, fileHeader.Filename)
	for {
		_, err1 := os.Stat(fp)
		if os.IsNotExist(err1) {
			break
		}
		// 文件存在
		// 算上传的图片和本身的图片是不是一样的，如果是一样的，那就直接返回之前的地址
		uploadFile, _ := fileHeader.Open()
		oldFile, _ := os.Open(fp)

		uploadFileHash := md5.MD5WithFile(uploadFile)
		oldFileHash := md5.MD5WithFile(oldFile)
		if uploadFileHash == oldFileHash {
			// 上传的图片，名称和内容都是一样的
			res.Ok("/"+fp, "上传成功", c)
			return
		}
		// 上传的图片，名称是一样的，但是内容不一样
		fileNameNotExt := strings.TrimSuffix(fileHeader.Filename, ext)
		newFileName := fmt.Sprintf("%s_%s%s", fileNameNotExt, random.RandStr(3), ext)
		fp = path.Join("uploads", global.Config.Upload.Dir, newFileName)
	}
	c.SaveUploadedFile(fileHeader, fp)

	res.Ok("/"+fp, "上传成功", c)
}
