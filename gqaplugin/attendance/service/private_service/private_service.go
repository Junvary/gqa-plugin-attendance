package private_service

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	pluginGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/attendance/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/attendance/model"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func InareaWithUserinfoList(inareaWithUserinfoDateRequest model.InareaWithUserinfoDateRequest) (err error, result []model.InareaWithUserinfo, total int64) {
	pageSize := inareaWithUserinfoDateRequest.PageSize
	offset := inareaWithUserinfoDateRequest.PageSize * (inareaWithUserinfoDateRequest.Page - 1)
	var resultList []model.InareaWithUserinfo
	var db *gorm.DB
	db = pluginGlobal.AttendanceDb.Model(&model.Inarea{})
	if inareaWithUserinfoDateRequest.InAreaTime != "" {
		db = db.Where("InAreaTime like ?", inareaWithUserinfoDateRequest.InAreaTime+"%")
	}
	db = db.Order(global.OrderByColumn("InAreaTime", false))
	db = db.Select("*").Joins("left join userinfo on userinfo.`GUID` = inarea.`UserGUID`").Scan(&resultList)
	if inareaWithUserinfoDateRequest.WorkNumber != "" {
		db = db.Where("WorkNumber = ?", inareaWithUserinfoDateRequest.WorkNumber)
	}
	if inareaWithUserinfoDateRequest.UserName != "" {
		db = db.Where("UserName = ?", inareaWithUserinfoDateRequest.UserName)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Find(&resultList).Error
	return err, resultList, total
}

func InareaWithUserinfoYear(inareaWithUserinfoYearRequest model.InareaWithUserinfoYearRequest) (err error, record []model.InareaWithUserinfo) {
	var resultList []model.InareaWithUserinfo
	var db *gorm.DB
	db = pluginGlobal.AttendanceDb.Model(&model.Inarea{})
	db = db.Where("InAreaTime like ? ", strconv.Itoa(time.Now().Year())+"-%")
	db = db.Order("substr(InAreaTime, 12)")
	db = db.Select("*").Joins("left join userinfo on userinfo.`GUID` = inarea.`UserGUID`").Scan(&resultList)
	if inareaWithUserinfoYearRequest.WorkNumber != "" {
		db = db.Where("WorkNumber = ?", inareaWithUserinfoYearRequest.WorkNumber)
	}
	err = db.First(&resultList).Error
	return err, resultList
}
