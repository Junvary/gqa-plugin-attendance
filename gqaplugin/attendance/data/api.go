package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginAttendanceSysApi = new(sysApi)

type sysApi struct{}

func (s *sysApi) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysApi{}).Where("api_group = ?", "plugin-attendance").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_api 表中attendance插件数据已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_api 表中attendance插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> attendance插件初始数据进入 sys_api 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> attendance插件初始数据进入 sys_api 表成功！")
		return nil
	})
}

var sysApiData = []gqaModel.SysApi{
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 901, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件Attendance：获取in-area-list",
	}}, ApiGroup: "plugin-attendance", ApiMethod: "POST", ApiPath: "/plugin-attendance/in-area-list"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 902, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件Attendance：获取in-area-year",
	}}, ApiGroup: "plugin-attendance", ApiMethod: "POST", ApiPath: "/plugin-attendance/in-area-year"},
}
