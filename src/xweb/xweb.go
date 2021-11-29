package xweb

import (
	"crypto/sha256"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"xcxc/xsql"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

var message string
var messagetype int
var r *gin.Engine

func InitWeb() (err error) {
	r = gin.Default()
	r.Static("/static", "../static") //css+picture
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
	r.SetFuncMap(template.FuncMap{ //html  |safe
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//	r.LoadHTMLGlob("./templates/**/*") //load templates
	r.HTMLRender = loadTemplates("../templates") //加载并准备渲染模板
	r.GET("/", home)
	r.GET("/index", home)
	r.GET("/student", student)
	r.GET("/login", login)

	r.GET("/wechatpay/:id", checkwechat)

	r.GET("/geturl/:id", GetChargeUrl)
	r.POST("/notify/", xsql.Notify)
	r.GET("/returnurl/:id", returnurl)
	r.GET("/checkpay/:id", checkpay)

	r.GET("/admin/:id", getStudent)
	r.DELETE("/admin/:id", Dele)
	r.POST("/admin", admin)

	// r.GET("/GetCity", getCity)

	r.GET("/edit/:id", startEdit)
	r.POST("/edit", submitEdit)

	r.GET("/dataSxy15530856229", data)

	r.POST("/submit", submit)
	r.GET("/submit", checkwechat)

	r.Run(":13488")
	return
}
func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "hi",
	})

}

func student(c *gin.Context) {
	c.HTML(http.StatusOK, "student.html", gin.H{
		// "citys": xsql.GetCityS(),
		"title": "学生登记",
	})

}
func login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})

}
func admin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "15530856229" && password == "Sxy15530856229" {
		t := strconv.Itoa(int(time.Now().Unix()))
		check := t + "tLRZBQk*b@OqubT^hNMo"
		s := sha256.Sum256([]byte(check))
		s2 := string(s[:])
		c.SetCookie("time", t, 1800, "/", "", false, true)
		c.SetCookie("acc", s2, 1800, "/", "", false, true)
		// citys := xsql.GetCityS()
		// fmt.Println("citys:==  ", citys)
		c.HTML(http.StatusOK, "table.html", gin.H{
			// "citys": xsql.GetCityS(),
		})
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"message":     "错误",
			"messagetype": 1,
		})
	}

}
func edit(c *gin.Context) {
	c.HTML(http.StatusOK, "edit.html", gin.H{
		"title": "编辑",
	})

}
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}
	// fmt.Println(includes)
	// 为layouts/和includes/目录生成 templates map
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
