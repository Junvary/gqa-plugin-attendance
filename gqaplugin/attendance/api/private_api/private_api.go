package private_api

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/attendance/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/attendance/service/private_service"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InareaWithUserinfoList(c *gin.Context) {
	var inareaWithUserinfoDateRequest model.InareaWithUserinfoDateRequest
	if err := c.ShouldBindJSON(&inareaWithUserinfoDateRequest); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, list, total := private_service.InareaWithUserinfoList(inareaWithUserinfoDateRequest); err != nil {
		global.GqaLog.Error("获取考勤统计数据失败！", zap.Any("err", err))
		global.ErrorMessage("获取考勤统计数据失败！"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  list,
			Page:     inareaWithUserinfoDateRequest.Page,
			PageSize: inareaWithUserinfoDateRequest.PageSize,
			Total:    total,
		}, c)
	}
}

func InareaWithUserinfoYear(c *gin.Context) {
	var inareaWithUserinfoYearRequest model.InareaWithUserinfoYearRequest
	if err := c.ShouldBindJSON(&inareaWithUserinfoYearRequest); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, record := private_service.InareaWithUserinfoYear(inareaWithUserinfoYearRequest); err != nil {
		global.GqaLog.Error("获取考勤统计数据失败！", zap.Any("err", err))
		global.ErrorMessage("获取考勤统计数据失败！"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records: record,
		}, c)
	}
}
