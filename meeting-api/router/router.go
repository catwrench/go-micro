package router

import (
	"github.com/gin-gonic/gin"
	"meeting-api/handler"
	"meeting-api/lib"
)

func NewRouter() *gin.Engine {
	route := gin.Default()

	spaceController := handler.SpaceController{}
	deviceController := handler.DeviceController{}
	roomController := handler.RoomController{}
	appletController := handler.AppletController{}

	//jaeger中间件
	route.Use(lib.JaegerMiddleware())

	webGroup := route.Group("/api/meeting")
	{
		//查询所有地点列表
		webGroup.GET("/spacesAll", spaceController.GetAllSpaces)
		spaceGroup := webGroup.Group("/spaces")
		{
			//查询地点列表
			spaceGroup.GET("", spaceController.GetSpaces)
			//查询地点详情
			spaceGroup.GET("/:id", spaceController.GetSpace)
			//新增地点
			spaceGroup.POST("", spaceController.CreateSpace)
			//编辑地点
			spaceGroup.PUT("/:id", spaceController.UpdateSpace)
			//删除地点
			spaceGroup.DELETE("/:id", spaceController.DelSpace)
			//启用/禁用地点
			spaceGroup.PUT("/:id/status", spaceController.UpdateSpaceStatus)
		}
		deviceGroup := webGroup.Group("/devices")
		{
			//查询设备列表
			deviceGroup.GET("", deviceController.GetDevices)
			//查询设备详情
			deviceGroup.GET("/:id", deviceController.GetDevice)
			//新增设备
			deviceGroup.POST("", deviceController.CreateDevice)
			//编辑设备
			deviceGroup.PUT("/:id", deviceController.UpdateDevice)
			//删除设备
			deviceGroup.DELETE("/:id", deviceController.DelDevice)
		}

		roomGroup := webGroup.Group("/rooms")
		{
			//查询会议室列表
			roomGroup.GET("", roomController.GetRooms)
			//查询会议室详情
			roomGroup.GET("/:id", roomController.GetRoom)
			//新增会议室
			roomGroup.POST("", roomController.CreateRoom)
			//编辑会议室
			roomGroup.PUT("/:id", roomController.UpdateRoom)
			//删除会议室
			roomGroup.DELETE("/:id", roomController.DeleteRoom)
		}
	}

	appletGroup := route.Group("/api/meetingApplet")
	{
		//查询会议列表
		appletGroup.GET("/reservations", appletController.GetReservations)
		//查询会议详情
		appletGroup.GET("/reservations/:id", appletController.GetReservation)
		//新增会议预约
		appletGroup.POST("/reservations", appletController.CreateReservation)

		//查询会议室列表
		appletGroup.GET("/rooms", appletController.GetRooms)
		//查询会议室详情
		appletGroup.GET("/rooms/:id", appletController.GetRoom)

		//查询地点列表
		appletGroup.GET("/spaces", appletController.GetSpaces)
	}
	return route
}
