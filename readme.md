# hatGo

标签（空格分隔）： 未分类

---

> **gin**框架构建，linux下采用**endless**实现热更新，mysql数据库操作采用**xorm**

## 使用
### 安装Golang官方依赖管理工具：dep
```
go get -u github.com/golang/dep/cmd/dep
```
### git项目到本地
```
git clone https://github.com/EddieChan1993/hatGo
```
## 日志
采用**beego日志模块**，具有文件大小及其个数配置，同时**sql日志**和**请求日志**分开存储，高并发下，
不会乱序，其次具有自定义日志功能，可以独立的记录其他内容

##快速部署
###通过app.sh快速部署
```
# 编译开发环境应用
./app.sh dev
# 编译生产环境应用
./app.sh prod
#启动应用(仅linux)
./app.sh start
`CTRL+Z`退出显示(任务中断，进程挂起)
#平滑重启(仅linux)
./app.sh restart
#停止程序(仅linux)
./app.sh stop
#获取程序执行状态(仅linux)
./app.sh status
```
### 通过辅助工具hat部署
[hat工具](https://github.com/EddieChan1993/hat)
