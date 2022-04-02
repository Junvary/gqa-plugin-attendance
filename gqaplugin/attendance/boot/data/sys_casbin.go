package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var PluginAttendanceSysCasbin = new(sysCasbin)

type sysCasbin struct{}

var sysCasbinData = []gormadapter.CasbinRule{
	// plugin-attendance组
	{Ptype: "p", V0: "super-admin", V1: "/plugin-attendance/in-area-list", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-attendance/in-area-year", V2: "POST"},
}

func (s *sysCasbin) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&[]gormadapter.CasbinRule{}).Where("v1 like ?", "/plugin-attendance%").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> casbin_rule 表中attendance插件数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[GQA-Plugin] --> casbin_rule 表中attendance插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysCasbinData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> attendance插件初始数据进入 casbin_rule 表成功！")
		global.GqaLog.Info("[GQA-Plugin] --> attendance插件初始数据进入 casbin_rule 表成功！")
		return nil
	})
}