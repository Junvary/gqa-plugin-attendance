package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginAttendanceSysApi = new(sysApi)

type sysApi struct{}

var sysApiData = []system.SysApi{
	{GqaModel: global.GqaModel{Status: "on", Sort: 901, Remark: "插件Attendance：获取in-area-list", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-attendance", ApiPath: "/plugin-attendance/in-area-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 902, Remark: "插件Attendance：获取in-area-year", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "plugin-attendance", ApiPath: "/plugin-attendance/in-area-year", ApiMethod: "POST",
	},
}

func (s *sysApi) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysApi{}).Where("api_group = ?", "plugin-attendance").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_api 表中attendance插件数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[GQA-Plugin] --> sys_api 表中attendance插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> attendance插件初始数据进入 sys_api 表成功！")
		global.GqaLog.Info("[GQA-Plugin] --> attendance插件初始数据进入 sys_api 表成功！")
		return nil
	})
}
