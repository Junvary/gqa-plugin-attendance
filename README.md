# 欢迎使用Gin-Quasar-Admin！

<div align=center>
<img src="https://i.loli.net/2020/12/14/cnJoF9r1BXY7Da5.png" width=200" height="200" />
<h1>Gin-Quasar-Admin</h1>
</div>

### [Gin-Quasar-Admin](https://github.com/Junvary/gin-quasar-admin)的Attendance插件

***

### 插件引入：

> 项目运行参考：https://github.com/Junvary/gin-quasar-admin



> 前台：
>
> ```go
> 安装前台：clone本仓库，将 Plugin 文件夹中的内容拷贝至：GQA-FRONTEND/src/pages/Plugin/
> ```



> 后台：
>
> ```go
> 打开gin-quasar-admin项目的： GQA-BACKEND/gqaplugin/gqa_plugin.go 文件
> 1.github引入方式：
> import "github.com/Junvary/gqa-plugin-attendance/gqaplugin/attendance"
> 2.本地引入方式：
> clone本仓库，将 gqaplugin 文件夹中的内容拷贝至：GQA-BACKEND/gqaplugin/
> import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/attendance"
> 
> 找到 PluginList 切片，加入 attendance.PluginAttendance,
> ```
>

### License

Copyright (c) 2021-time.Now()    Junvary

[MIT License](https://github.com/Junvary/gin-quasar-admin/blob/main/LICENSE)