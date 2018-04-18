package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	cpi "github.com/yizenghui/fire"
	c "github.com/yizenghui/fire/controller"
)

//CheckSubcribeUpdate  每天处理订阅更新
func CheckSubcribeUpdate() {
	ticker := time.NewTicker(time.Hour * 6)
	for _ = range ticker.C {
		go cpi.RunSubcribePostUpdateCheck()
	}
}

func main() {
	// go CheckSubcribeUpdate()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome to fire minapp api, this build by yizenghui.com for go!")
	})
	// 获取openid
	e.GET("/getopenid", func(c echo.Context) error {
		code := c.QueryParam("code")
		ret, _ := cpi.GetOpenID(code)
		return c.JSON(http.StatusOK, ret)
	})

	// 用户签名
	e.GET("/sign", c.Sign)
	// 解密数据内容
	e.GET("/crypt", c.Crypt)

	// 记录分享 (我们现在通过分享次数进行排序)
	// e.GET("/push", func(c echo.Context) error {
	// 	openID := c.QueryParam("openid")
	// 	url := c.QueryParam("url")
	// 	cs := cpi.ShareLog(openID, url)
	// 	type Ret struct {
	// 		Status bool
	// 	}
	// 	return c.JSON(http.StatusOK, Ret{Status: cs})
	// })

	// 记录分享 (我们现在通过分享次数进行排序)
	e.POST("/push", c.NewPush)

	// 创建任务
	e.POST("/task", func(c echo.Context) error {
		openID := c.QueryParam("openid")
		url := c.QueryParam("url")
		cs := cpi.ShareLog(openID, url)
		type Ret struct {
			Status bool
		}
		return c.JSON(http.StatusOK, Ret{Status: cs})
	})

	// 获取任务列表
	e.GET("/task", func(c echo.Context) error {
		openID := c.QueryParam("openid")
		url := c.QueryParam("url")
		cs := cpi.ShareLog(openID, url)
		type Ret struct {
			Status bool
		}
		return c.JSON(http.StatusOK, Ret{Status: cs})
	})

	// 获取推荐码(图片资源)
	e.GET("/qrcode", func(c echo.Context) error {
		scene := c.QueryParam("scene")
		page := `pages/index/index`
		if scene == "" {
			return c.HTML(http.StatusOK, "")
		}
		fileName, err := cpi.GetwxCodeUnlimit(scene, page)
		if err == nil {
			http.ServeFile(c.Response().Writer, c.Request(), fileName)
		} else {
			http.ServeFile(c.Response().Writer, c.Request(), fileName)
		}
		var err2 error
		return err2
	})
	// 图标
	e.File("favicon.ico", "images/favicon.ico")
	e.Logger.Fatal(e.Start(":8009"))
	// e.Logger.Fatal(e.StartAutoTLS(":443"))

}