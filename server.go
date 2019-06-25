package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"go-gin-demo/routes"
	"log"
	"net/http"
)

func Recovery() {
	if err := recover(); err != nil {
		glog.Errorf("[Recovery] %s", err)
		glog.Flush()
	}
}

func main()  {
	defer glog.Flush()
	defer Recovery()

	flag.Parse()

	//gin.SetMode 全局设置环境 gin.DebugMode为开发环境 线上环境为gin.ReleaseMode
	gin.SetMode(gin.DebugMode)

	e := gin.New()
	e.Use()

	routes.LoadApiRouter(e)

	httpPort := 8080

	endPoint := fmt.Sprintf(":%d", httpPort)

	//log.Printf("Listening and serving HTTP on %s\n", endPoint)

	glog.Infof("Listening and serving HTTP on %s\n", endPoint)



	//glog.Flush()
	err := http.ListenAndServe(endPoint, e)

	if err != nil {
		log.Printf("Listening and serving HTTP on %s\n", err)
	}


}

