package xweb

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"xcxc/email"
	"xcxc/xsql"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func submit(c *gin.Context) {
	var st, u xsql.Student
	myDate := c.PostForm("date")
	st.Year, _ = strconv.Atoi(myDate[0:4])
	st.Month, _ = strconv.Atoi(myDate[5:7])
	st.Day, _ = strconv.Atoi(myDate[8:10])
	st.Name = c.PostForm("name")
	st.Email = c.PostForm("email")
	st.Phone = c.PostForm("phone")
	st.Applydate = time.Now().Local().Format("2006-01-02-15:04:05")
	st.Id = st.Phone + strconv.Itoa(st.Month*100+st.Day) + st.Applydate
	fmt.Println(st.Id)
	st.School = c.PostForm("school")
	st.Isback = c.PostForm("isback")
	st.City = c.PostForm("city")
	st.Ispayed = "未付"

	xsql.Userdb = xsql.Userdb.Table("students")
	xsql.Userdb.Table("student")
	xsql.Userdb.Where("id = ?", st.Id).Take(&u)
	if u.Id == "" {
		message = "信息录入成功！请核对后付款"
		messagetype = 1
		if err := xsql.Userdb.Create(st).Error; err != nil {
			fmt.Println("失败", err)
		}
		fmt.Println(st)
		s := fmt.Sprintf("请仔细核对自己的信息:\n姓名:%s\n手机号:%s\n行程:%s\n日期:%s\n学校:%s\n家:%s\n\n同时，我们会发送一封包含你录入信息内容的邮件到你的邮箱内，请注意查收。",
			st.Name, st.Phone, st.Isback, myDate, st.School, st.City)
		c.HTML(http.StatusOK, "submit.html", gin.H{
			"title":       "请付款",
			"message":     message,
			"messagetype": messagetype,
			"check":       s,
			"id":          st.Id,
		})
		go email.CheckStudentInfo(st)

	} else {
		message = "不要刷新或者重复提交！请尝试重新登记后等待"
		messagetype = 2
		c.HTML(http.StatusOK, "student.html", gin.H{
			"message":     message,
			"messagetype": messagetype,
		})
	}
}
func data(c *gin.Context) {
	var s xsql.StudentS
	searchParams := c.Query("searchParams")
	jsoniter.Unmarshal([]byte(searchParams), &s)
	fmt.Println("开始获取数据")
	mydata := xsql.GetData(s)
	fmt.Println("获取数据结束")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	min := limit * (page - 1)
	max := limit * page
	if max > xsql.DataCount {
		max = xsql.DataCount
	}
	myData := mydata[min:max]
	check_ := checkecookie(c)
	fmt.Println(check_)
	if check_ == -1 {
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"msg":   "",
			"count": xsql.DataCount,
			"data":  myData,
		})
	} else if check_ == 0 {

		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"msg":   "请重新登陆！",
			"count": xsql.DataCount,
			"data":  myData,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"msg":   "登录超时，请尝试重新登录！",
			"count": xsql.DataCount,
			"data":  myData,
		})
	}
}

func checkecookie(c *gin.Context) int {
	nowtime, _ := c.Cookie("time")
	check, _ := c.Cookie("acc")
	check1 := nowtime + "tLRZBQk*b@OqubT^hNMo"
	timeNowInt := int(time.Now().Unix())
	timeOldInt, _ := strconv.Atoi(nowtime)
	i := timeNowInt - timeOldInt
	fmt.Println(timeNowInt, "   ", timeOldInt, " ", i)
	s := sha256.Sum256([]byte(check1))
	s2 := string(s[:])
	if s2 != check || nowtime == "" {
		return 0 //jiao yan shi bai
	} else if i < 0 || i >= 1800 {
		//chao shi
		return i
	} else {
		//cheng gong
		return -1
	}
}
func startEdit(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	c.HTML(http.StatusOK, "edit.html", gin.H{
		"id": id,
	})
}
func submitEdit(c *gin.Context) {
	var st xsql.Student
	id := c.PostForm("id")
	st.Year, _ = strconv.Atoi(c.PostForm("year"))
	st.Month, _ = strconv.Atoi(c.PostForm("month"))
	st.Day, _ = strconv.Atoi(c.PostForm("day"))
	st.Name = c.PostForm("name")
	st.Phone = c.PostForm("phone")
	st.Applydate = time.Now().Local().Format("2006-01-02-15:04:05")
	st.Id = st.Phone + strconv.Itoa(st.Month*100+st.Day) + st.Applydate
	st.School = c.PostForm("school")
	st.Isback = c.PostForm("isback")
	st.City = c.PostForm("city")
	xsql.Userdb.Table("student")
	xsql.Userdb.Where("id = ?", id).Update("id", "")
	if err := xsql.Userdb.Create(st).Error; err != nil {
		fmt.Println("失败", err)
	}
	fmt.Println(st)
	c.HTML(http.StatusOK, "checkwechat.html", gin.H{
		"message":     "修改成功！可以关闭本页面",
		"messagetype": 1,
	})
}
func Dele(c *gin.Context) {
	// Param
	check := checkecookie(c)
	if check == -1 {
		id := c.Param("id")
		xsql.DelData(id)
	}
}

func getStudent(c *gin.Context) {
	id := c.Param("id")
	var u xsql.Student
	xsql.Userdb = xsql.Userdb.Table("students")
	xsql.Userdb.Where("id = ?", id).Take(&u)
	fmt.Println(u, id)
	c.JSON(http.StatusOK, u)
	fmt.Println("ok")
}
