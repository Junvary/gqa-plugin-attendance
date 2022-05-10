package privateapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/attendance/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/attendance/service/privateservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InareaWithUserinfoList(c *gin.Context) {
	var inareaWithUserinfoDateRequest model.InareaWithUserinfoDateRequest
	if err := gqaModel.RequestShouldBindJSON(c, &inareaWithUserinfoDateRequest); err != nil {
		return
	}
	if err, list, total := privateservice.InareaWithUserinfoList(inareaWithUserinfoDateRequest); err != nil {
		gqaGlobal.GqaLogger.Error("获取考勤统计数据失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取考勤统计数据失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  list,
			Page:     inareaWithUserinfoDateRequest.Page,
			PageSize: inareaWithUserinfoDateRequest.PageSize,
			Total:    total,
		}, c)
	}
}

func InareaWithUserinfoYear(c *gin.Context) {
	var inareaWithUserinfoYearRequest model.InareaWithUserinfoYearRequest
	if err := gqaModel.RequestShouldBindJSON(c, &inareaWithUserinfoYearRequest); err != nil {
		return
	}
	if err, record := privateservice.InareaWithUserinfoYear(inareaWithUserinfoYearRequest); err != nil {
		gqaGlobal.GqaLogger.Error("获取考勤统计数据失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取考勤统计数据失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records: record,
		}, c)
	}
}
