package main

import (
	"fmt"
	"os"
	"xcxc/xsql"
	"xcxc/xweb"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("use local passwd")
		go xsql.Initxsql(false)
	} else {
		go xsql.Initxsql(true)
		fmt.Println("use aliyun passwd")
	}
	xweb.InitWeb()

}
