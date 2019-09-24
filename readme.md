# hatGo

> **gin**框架构建，linux下采用**endless**实现热更新，mysql数据库操作采用**xorm**

# 安装
go mod init hatgo

go build 拉取依赖包

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
