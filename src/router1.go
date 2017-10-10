package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// 路由信息
type RouteInfo struct {
	Method   string
	Path     string
	Handlers gin.HandlersChain
}

// make route info
func Route(method string, path string, handlers ...gin.HandlerFunc) RouteInfo {
	var routeInfo RouteInfo
	routeInfo.Method = method
	routeInfo.Path = path
	routeInfo.Handlers = append(routeInfo.Handlers, RecordInfo)
	routeInfo.Handlers = append(routeInfo.Handlers, handlers...)

	return routeInfo
}

func RecordInfo(c *gin.Context) {
	log.Printf("Request URL:[%s] %s \n", c.Request.Method, c.Request.URL)
	c.Request.ParseMultipartForm(1024 * 1024)
	log.Println("Request Param:", c.Request.MultipartForm)
	c.Next()
}

var Routes = []RouteInfo{
	// 获取评论列表
	Route(
		"GET",
		"/hello",
		otherMw(),
		loginTest,
	),
	Route(
		"GET",
		"/hello1",
		otherMw(),
		loginTest,
	),
}

func startGin1() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	for _, route := range Routes {
		switch route.Method {
		case "GET":
			r.GET(route.Path, route.Handlers...)
		case "POST":
			r.POST(route.Path, route.Handlers...)
		}
	}
	err := r.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
