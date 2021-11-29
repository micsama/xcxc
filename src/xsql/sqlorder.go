package xsql

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type order struct {
	Aoid      string `form:"aoid"`
	Order_id  string `form:"order_id"`
	Pay_price string `form:"pay_price"`
	Pay_time  string `form:"pay_time"`
	More      string `form:"more"`
	Detail    string `form:"detail"`
	Sign      string `form:"sign"`
}

const (
	Aid = "19169"
	// appsecret = "74b2501932504c9680970d1d8d4b5add"
	Appsecret = "dab20eb7f4c048ecb2a19ca315a1566e"
)

func Notify(c *gin.Context) {
	id := c.Query("s")
	fmt.Println("notify id: ", id)
	// o := order{c.PostForm("aoid")}
	var o order
	if err := c.ShouldBind(&o); err == nil {
		fmt.Println(o)
		S := o.Aoid + o.Order_id + o.Pay_price + o.Pay_time + Appsecret
		sign := fmt.Sprintf("%x", md5.Sum([]byte(S)))
		Ssign := o.Sign
		if sign == strings.ToLower(Ssign) {
			fmt.Println("sign ok")
			c.JSON(http.StatusOK, gin.H{})
			Userdb = Userdb.Table("orders")
			if o.Detail == "{}" {
				o.Detail = "{WechatPay}"
			}
			if err := Userdb.Create(o).Error; err != nil {
				fmt.Println("写入数据库失败", err)
			} else { //彻底成功
				Userdb = Userdb.Table("students")
				x := Userdb.Where("id = ?", o.Order_id)
				x.Update("ispayed", "已付")
				x.Update("order", o.Aoid)
			}
		} else {
			fmt.Println("sign error")
		}
	} else {
		fmt.Println("notify error")
	}
}

func CheckById(id string) bool {
	var s Student
	Userdb = Userdb.Table("students")
	Userdb.Where("id = ?", id).Take(&s)
	if s.Ispayed == "已付" {
		return true
	} else if s.Ispayed == "未付" {
		return false
	} else {
		fmt.Println("error")
		return false
	}
}

func UserUrl(id string, t string) (apiurl string, sign string) {
	var (
		name     string
		pay_type string
		price    string
		order_id string
		// order_uid  string
		notify_url string
		return_url string
		cancel_url string
		more       string
		// exprie     int
		c city
		s StudentS
	)
	fmt.Println(id)
	Userdb = Userdb.Table("students")
	Userdb.Where("id = ?", id).Take(&s)
	var tt string
	if t == "alipay" {
		tt = "pay"
		pay_type = "alipay"
		Userdb = Userdb.Table("citys")
		Userdb.Where("city = ?", s.City).Take(&c)
		price = c.Price
		order_id = s.Id
		more = order_id
		saw := "32h9fsdfh8#@$WEFsafajfweifhioag4352345hioapsdjf24"
		fmt.Println(sign + saw)
		ss := sha256.Sum256([]byte(sign + saw))
		sss := fmt.Sprintf("%x", ss)
		// TODO: 更换域名后记得更改  <02-08-21, MicSama> //
		notify_url = "http://121.5.31.22:13488/notify?c=" + sss
		name = price
		S := name + pay_type + price + order_id + notify_url + Appsecret
		fmt.Println(S)
		sign = fmt.Sprintf("%x", md5.Sum([]byte(S)))
		apiurl = fmt.Sprintf("https://xorpay.com/api/%s/%s?name=%s&pay_type=%s&price=%s&order_id=%s&notify_url=%s&return_url=%s&more=%s&sign=%s",
			tt, Aid, name, pay_type, price, order_id, notify_url, return_url, more, sign)
		return
	} else if t == "wechat" {
		tt = "cashier"
		pay_type = "jsapi"
		Userdb = Userdb.Table("citys")
		Userdb.Where("city = ?", s.City).Take(&c)
		price = c.Price
		order_id = s.Id
		more = order_id
		saw := "32h9fsdfh8#@$WEFsafajfweifhioag4352345hioapsdjf24"
		fmt.Println(sign + saw)
		ss := sha256.Sum256([]byte(sign + saw))
		sss := fmt.Sprintf("%x", ss)
		// TODO: 更换域名后记得更改  <02-08-21, MicSama> //
		notify_url = "http://121.5.31.22:13488/notify?c=" + sss
		return_url = "http://121.5.31.22:13488/wechatpay/ok"
		cancel_url = "http://121.5.31.22:13488/wechatpay/cancel"
		name = price
		S := name + pay_type + price + order_id + notify_url + Appsecret
		fmt.Println(S)
		sign = fmt.Sprintf("%x", md5.Sum([]byte(S)))
		apiurl = fmt.Sprintf("https://xorpay.com/api/%s/%s?name=%s&pay_type=%s&price=%s&order_id=%s&notify_url=%s&cancel_url=%s&return_url=%s&more=%s&sign=%s",
			tt, Aid, name, pay_type, price, order_id, notify_url, cancel_url, return_url, more, sign)
		return
	} else {
		apiurl = "/"
		return
	}
}
