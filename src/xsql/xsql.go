package xsql

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Userdb *gorm.DB
var DataCount int
var myData []Student

type StudentS struct {
	Id        string
	Name      string `form:"name"`
	Phone     string `form:"phone" `
	School    string `form:"school" `
	Isback    string `form:"isback"`
	City      string `form:"city"`
	Order     string `form:"price" `
	Email     string `form:"email" `
	Year      string
	Month     string
	Day       string
	Applydate string `form:"user" `
	Ispayed   string
}
type Student struct {
	Id        string
	Name      string `form:"name"`
	Phone     string `form:"phone" `
	School    string `form:"school" `
	Isback    string `form:"isback"`
	City      string `form:"city"`
	Email     string `form:"email" `
	Order     string `form:"price" `
	Year      int
	Month     int
	Day       int
	Applydate string `form:"user" `
	Ispayed   string
}

var username = "root"
var password = "123456"

const (
	host    = "localhost"
	port    = 3306
	Dbname  = "xcx"
	timeout = "10s"
)

func Initxsql(isUsed bool) {
	if isUsed {
		username = "xcx"
		password = "Yys6pM4dCSAZd2SF"
	}
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	Userdb, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("mysql connected success")
	}
}
func GetData(s StudentS) []Student {
	MyData := make([]Student, 0, 100)
	Userdb = Userdb.Table("students")
	y, _ := strconv.Atoi(s.Year)
	m, _ := strconv.Atoi(s.Month)
	d, _ := strconv.Atoi(s.Day)
	s2 := Student{
		Id:        s.Id,
		Name:      s.Name,
		Phone:     s.Phone,
		School:    s.School,
		Isback:    s.Isback,
		City:      s.City,
		Order:     s.Order,
		Email:     s.Email,
		Year:      y,
		Month:     m,
		Day:       d,
		Applydate: s.Applydate,
		Ispayed:   s.Ispayed,
	}
	Userdb.Where("id != ? ", "").Find(&myData)
	for _, i := range myData {
		if s.Id == "" || i.Id == s.Id {
			if s.Name == "" || i.Name == s.Name {
				if s.Phone == "" || i.Phone == s.Phone {
					if s.School == "" || i.School == s.School {
						if s.Isback == "" || i.Isback == s.Isback {
							if s.City == "" || i.City == s.City {
								if s.Order == "" || i.Order == s2.Order {
									if s.Year == "" || i.Year == s2.Year {
										if s.Month == "" || i.Month == s2.Month {
											if s.Day == "" || i.Day == s2.Day {
												if s.Applydate == "" || i.Applydate == s.Applydate {
													if s.Ispayed == "" || i.Ispayed == s.Ispayed {
														if s.Email == "" || i.Email == s.Email {
															MyData = append(MyData, i)
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	DataCount = len(MyData[:])
	return MyData
}
func DelData(s string) bool {
	Userdb := Userdb.Table("students")
	fmt.Printf("!!!s=: %s \n", s)
	Userdb.Where("id = ?", s).Update("id", "")
	return true
}

type city struct {
	City  string
	Price string
}

// func GetCity() (citys []city) {
// 	Userdb = Userdb.Table("citys")
// 	Userdb.Find(&citys)
// 	return
// }
// func GetCityS() (s string) {
// 	var citys []city
// 	Userdb = Userdb.Table("citys")
// 	Userdb.Find(&citys)
// 	for _, c := range citys {
// 		s = s + "<option value=\"" + c.City + "\">" + c.City + " -- " + c.Price + "</option>"
// 	}
// 	return
// }
