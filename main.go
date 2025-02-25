package main

import (
	"embed"
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed static
var staticFS embed.FS

var port = flag.String("port", ":9999", "网络端口，格式为【:9999】")
var static = flag.String("static", "", "静态目录")
var index = flag.String("index", "", "入口文件")

func init() {
	flag.Parse()
}

func main() {
	log.Println("server starting...")

	r := gin.Default()

	// 静态文件服务
	if *static != "" {
		r.StaticFS("/static/", gin.Dir(*static, true))
	}

	r.StaticFS("/fs/", http.FS(staticFS))

	if *index != "" {
		if *static != "" {
			r.GET("/", func(c *gin.Context) {
				http.ServeFile(c.Writer, c.Request, *static+"/"+*index)
			})
		} else {
			r.GET("/", func(c *gin.Context) {
				c.FileFromFS("static/"+*index, http.FileSystem(http.FS(staticFS)))
			})
		}
	}

	r.Run(*port)
}
