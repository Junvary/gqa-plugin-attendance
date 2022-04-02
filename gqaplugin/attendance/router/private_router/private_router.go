package private_router

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/attendance/api/private_api"
	"github.com/gin-gonic/gin"
)

func InitPrivateRouter(privateGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	{
		//获取某日打卡记录
		privateGroup.POST("in-area-list", private_api.InareaWithUserinfoList)
		//年度打卡记录
		privateGroup.POST("in-area-year", private_api.InareaWithUserinfoYear)
	}
}
