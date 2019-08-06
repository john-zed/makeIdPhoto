package main

import (
	. "../staff_remote/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {

	router := gin.Default()

	//允许所有的跨域请求
	router.Use(cors.Default())

	router.Static("/index", "./view")
	router.Static("/upload", "./upload")
	//router.GET("/", IndexApi)

	router.POST("/addStaff", AddStaffTrainingApi)

	router.GET("/getStaffs", QueryStaffTrainingsApi)

	// 文件上传
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.POST("/upload", UploadFile)

	router.POST("/makePhoto", MakePhoto)

	// 主扫支付
	router.POST("/precreate", LiantuofuPrecreate)

	// 公众号支付
	router.POST("/openPay", LiantuofuOpenPay)

	// 订单查询
	router.POST("/queryOrder", LiantuofuQueryOrder)

	return router

}
