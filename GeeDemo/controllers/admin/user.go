package controllers

import (
	"gee"
	"net/http"
)

type UserControllers struct {
	AdminControllers
}

func (con UserControllers) Index(c *gee.Context) {
	c.String(200, "用户列表")
}

func (con UserControllers) Add(c *gee.Context) {
	c.HTML(http.StatusOK, "admin/useradd.html", gee.H{})
}

/*
func (con UserControllers) DoUpload(c *gee.Context) {
	username := c.PostForm("username")
	file, err := c.FormFile("face") //gee暂未实现，待补充（2022.3.1）

	if err == nil {
		//判断后缀
		extName := path.Ext(file.Filename)

		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}

		if _, ok := allowExtMap[extName]; !ok {
			c.String(200, "上传类型不合法")
			return
		}

		//创建文件保存目录

		//day := models.GetDay()
		//····
		//此部分与web关系不太大，暂且略过
		//接下来判断目录是否存在，不存在创建目录，
		//更改文件名，存入文件夹
		//按时间存放，好找，也不会顶替之前同名图片

		dst := path.Join("./static/upload", file.Filename)
		c.SaveUploadedFile(file, dst) //gee暂未实现，待补充（2022.3.1）
		c.JSON(http.StatusOK, gee.H{
			"success":  true,
			"username": username,
			"dst":      dst,
		})
	} else {
		c.String(200, "没文件")
	}
}
*/
