package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginAttendanceSysMenu = new(sysMenu)

type sysMenu struct{}

func (s *sysMenu) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysMenu{}).Where("parent_code = ?", "GqaPluginAttendance").Or("name = ?", "GqaPluginAttendance").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_menu 表中attendance插件菜单已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_menu 表中attendance插件菜单已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> attendance插件初始数据进入 sys_menu 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> attendance插件初始数据进入 sys_menu 表成功！")
		return nil
	})
}

var sysMenuData = []gqaModel.SysMenu{
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 901, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是考勤统计插件",
	}}, IsPlugin: "yes", Name: "GqaPluginAttendance", Title: "考勤统计", Icon: "house", Path: "", Component: ""},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "考勤统计",
	}}, IsPlugin: "yes", Name: "plugin-attendance-all", Title: "考勤统计", Icon: "newspaper", Path: "/plugin-attendance/attendance/all", Component: "plugins/attendance/All/index", ParentCode: "GqaPluginAttendance"},
}
