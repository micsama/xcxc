package xweb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"xcxc/xsql"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func GetChargeUrl(c *gin.Context) {
	id := c.Param("id")
	t := c.Query("t")
	if t == "alipay" {
		url, sign := xsql.UserUrl(id, t)
		fmt.Println("\n-----------------\nurl:", url)
		c.SetCookie("sign", sign, 1200, "/", "", false, true)
		contentType := "application/x-www-form-urlencoded"
		resp, err := http.Post(url, contentType, nil)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("get resp!!----error ---!!!", err)
		}
		fmt.Println(string(b))
		any := jsoniter.Get(b, "info", "qr")
		c.JSON(http.StatusOK, gin.H{
			"url": any.ToString(),
		})
	} else if t == "wechat" {
		url, _ := xsql.UserUrl(id, t)
		c.JSON(http.StatusOK, gin.H{
			"url": url,
		})

	}
}

func returnurl(c *gin.Context) {
	// check_ := c.Param("id")
	// fmt.Println(check_)
	// sign, err := c.Cookie("sign")
	// if err != nil {
	// 	c.HTML(http.StatusOK, "returnurl.html", gin.H{
	// 		"message":     "错误！",
	// 		"messagetype": 2,
	// 	})
	// }
	// saw := "32h9fsdfh8#@$WEFsafajfweifhioag4352345hioapsdjf24"
	// sss := string(sha256.New().Sum([]byte(sign + saw)))
	// if sss == check_ {
	// 	message = "成功！"
	// 	c.HTML(http.StatusOK, "returnurl.html", gin.H{
	// 		"message":     message,
	// 		"messagetype": 1,
	// 	})
	// } else {
	// 	c.HTML(http.StatusOK, "returnurl.html", gin.H{
	// 		"message":     "错误！",
	// 		"messagetype": 2,
	// 	})
	// }

}
func checkpay(c *gin.Context) {
	id := c.Param("id")
	if xsql.CheckById(id) {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "wait",
		})
	}
}
func checkwechat(c *gin.Context) {
	status := c.Param("id")
	if status == "ok" {
		c.HTML(http.StatusOK, "checkwechat.html", gin.H{
			"message":     "付款成功！",
			"messagetype": 1,
		})
	} else {
		c.HTML(http.StatusOK, "checkwechat.html", gin.H{
			"message":     "付款取消！请重新填写并付款",
			"messagetype": 2,
		})
	}
}
