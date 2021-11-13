package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

//日志文件

func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log") //创建文件
	// gin.DefaultWriter = io.MultiWriter(f)            //日志写到文件中
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) // 同时将日志写入文件和控制台

	//================上面的就是日志文件需要的代码
	r := gin.Default()
	r.GET("/a", func(c *gin.Context) {
		c.String(200, "aaaaa")
	})
	r.Run(":8088")
}
