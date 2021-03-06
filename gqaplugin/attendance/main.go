package attendance

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/attendance/boot"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/attendance/data"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/attendance/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/attendance/router/privaterouter"
	"github.com/gin-gonic/gin"
)

var PluginAttendance = new(attendance)

type attendance struct{}

func (f *attendance) PluginCode() string {
	return "plugin-attendance"
}

func (f *attendance) PluginName() string {
	return "考勤统计"
}

func (f *attendance) PluginVersion() string {
	return "v0.0.1"
}

func (f *attendance) PluginMemo() string {
	return "这是考勤统计插件"
}

func (f *attendance) PluginRouterPublic(publicGroup *gin.RouterGroup) {
}

func (f *attendance) PluginRouterPrivate(privateGroup *gin.RouterGroup) {
	privaterouter.InitPrivateRouter(privateGroup)
}

func (f *attendance) PluginMigrate() []interface{} {
	return nil
}

func (f *attendance) PluginData() []interface{ LoadData() (err error) } {
	var DataList = []interface{ LoadData() (err error) }{
		data.PluginAttendanceSysApi,
		data.PluginAttendanceSysRoleApi,
		data.PluginAttendanceSysMenu,
		data.PluginAttendanceSysRoleMenu,
	}
	return DataList
}

func init() {
	global.AttendanceDb = boot.Mysql()
	if global.AttendanceDb != nil {
		fmt.Println("插件Attendance连接数据库成功！")
	} else {
		fmt.Println("插件Attendance连接数据库失败！")
	}
}
