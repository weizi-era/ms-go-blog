package main

import (
	"ms-go-blog/common"
	"ms-go-blog/server"
)

func init() {
	// 模板加载
	//log.Println("执行main的init")
	println("执行main的init")
	common.LoadTemplate()
}

func main() {

	server.App.Start("127.0.0.1", "8080")
}
