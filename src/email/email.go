package email

import (
	"fmt"
	"net/smtp"
	"xcxc/xsql"
)

const (
	SMTP_MAIL_HOST     = "smtp.163.com"
	SMTP_MAIL_PORT     = "25"
	SMTP_MAIL_USER     = "Sxy13214587977@163.com"
	SMTP_MAIL_PWD      = "ROJQWSXAIWMJKKCX"
	SMTP_MAIL_NICKNAME = "【明天客运】"
	muban              = `<!DOCTYPE html>
		<html>
		  <head>
			<meta charset=\"UTF-8\" />
			<meta name=\"viewport\" content=\"width=device-width\" />
			<title>回执</title>
		  </head>
		  <body>
			<div style=\"font-size: 2rem\">
			  %s 你好， 欢迎使用本系统的服务，请确认录入信息无误：
			</div>
			<div style=\"margin-left: 3rem;font-size: 1.2rem;\">
			  <div>姓名：%s</div>
			  <div>手机号：%s</div>
			  <div>学校：%s</div>
			  <div>返校/回家：%s</div>
			  <div>%s：%s</div>
			  <div>预约时间：%s</div>
			</div>
			<div style=\"font-size:1.3rem;text-align:right\">
				提示：如果录入信息有误且未付款，可以重新从主页录入；如果已付款，请联系管理人员修改。	</br>【明天客运】</div>
		  </body>
		</html>`
)

func CheckStudentInfo(info xsql.Student) {
	fmt.Println("开始准备发送")
	var home string
	if info.Isback == "回家" {
		home = "回家重点"
	} else {
		home = "返校起点"
	}
	body := fmt.Sprintf(muban, info.Name, info.Name, info.Phone, info.School, info.Isback, home, info.City, info.Applydate)
	fmt.Println("内容：\n", body)
	add := []string{info.Email}
	err := SendMail(add, "明天客运回执", body)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}

}

func SendMail(address []string, subject, body string) (err error) {
	// 认证, content-type设置
	auth := smtp.PlainAuth("", SMTP_MAIL_USER, SMTP_MAIL_PWD, SMTP_MAIL_HOST)
	contentType := "Content-Type: text/html; charset=UTF-8"

	// 因为收件人，即address有多个，所以需要遍历,挨个发送
	for _, to := range address {
		s := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s", to, SMTP_MAIL_NICKNAME, SMTP_MAIL_USER, subject, contentType, body)
		msg := []byte(s)
		addr := fmt.Sprintf("%s:%s", SMTP_MAIL_HOST, SMTP_MAIL_PORT)
		err = smtp.SendMail(addr, auth, SMTP_MAIL_USER, []string{to}, msg)
		if err != nil {
			return err
		}
	}
	return err
}
