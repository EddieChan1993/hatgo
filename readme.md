# hatGo

> **gin**框架构建，linux下采用**endless**实现热更新，mysql数据库操作采用**xorm**

## 使用
### 安装Golang官方依赖管理工具：
**dep**,无奈速度太慢
```
go get -u github.com/golang/dep/cmd/dep
```
**govendor**
```
go get -u github.com/kardianos/govendor
#初始化
govendor init
#从本地添加项目所需依赖到vendor
govendor add +e
#同步vendor.json
govendor sync
#删除不用的包
govendor remove +u
#添加依赖，前提时本地已经存在
governdor add xx
#添加或更新包到本地 vendor 目录
govendor fetch urlpath
#类似 go get 目录，拉取依赖包到 vendor 目录
govendor get urlpath
```

### 下载本地
选择需要的[版本](https://github.com/EddieChan1993/hatgo/releases),将其部署到 **$GOPATH/src**
## 日志
采用**beego日志模块**，具有文件大小及其个数配置，同时**sql日志**和**请求日志**分开存储，高并发下，
不会乱序，其次具有自定义日志功能，可以独立的记录其他内容

## 快速部署
### 通过app.sh快速部署
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
#添加或更新包到本地 vendor 目录
govendor fetch urlpath
#类似 go get 目录，拉取依赖包到 vendor 目录
govendor get urlpath
```

### 通过辅助工具hat部署
[hat工具](https://github.com/EddieChan1993/hat)
```
$ hat
Usage:

        hat [arguments] command

The commands are:

        -v [version_code] -n [app_name] dev                create dev's program and eg version_code=1.0
        -v [version_code] -n [app_name] prod               create prod's program and eg version_code=1.0
        -n [app_name] start                                start program and default app_name=basename $PWD,next eq
        -n [app_name] restart                              restart program
        -n [app_name] stop                                 stop program
        -n [app_name] status                               status program
        help                                               look up help
        ver_dev                                            look up dev's version log
        ver_prod                                           look up prod's version log

```
