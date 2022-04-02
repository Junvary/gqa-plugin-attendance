package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginAttendanceSysMenu = new(sysMenu)

type sysMenu struct{}

var sysMenuData = []system.SysMenu{
	{GqaModel: global.GqaModel{Status: "on", Sort: 901, Remark: "这是考勤统计插件", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "GqaPluginAttendance", Title: "考勤统计", Icon: "house",
		Path: "", Component: "",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 1, Remark: "这是考勤统计", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "plugin-attendance-all", ParentCode: "GqaPluginAttendance", Title: "考勤统计", Icon: "newspaper",
		Path: "/plugin-attendance/attendance/all", Component: "/Plugin/Attendance/All/index",
	},
}

func (s *sysMenu) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysMenu{}).Where("parent_code = ?", "GqaPluginAttendance").Or("name = ?", "GqaPluginAttendance").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_menu 表中attendance插件菜单已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[GQA-Plugin] --> sys_menu 表中attendance插件菜单已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> attendance插件初始数据进入 sys_menu 表成功！")
		global.GqaLog.Info("[GQA-Plugin] --> attendance插件初始数据进入 sys_menu 表成功！")
		return nil
	})
}
