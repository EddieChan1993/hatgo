#!/bin/bash
option=$1 #参数
appName=${2} #应用名称
appName=${appName:=main}

#编译开发应用
buildDevApp(){
    if read -p "please enter version:" version
    then
             if [ ! -z $version ]
            then
                    go build -ldflags "-X main._version_=$version" -o ${appName}
                    echo "Dev programe is success"
                    echo "App name is 【${appName}】"
                    echo "Version is 【$version】"
            else
                    echo "version is none"
            fi
    else
       echo "sorry. too slow!"
    fi
}

#编译生产应用
buildProdApp(){
    if read -p "please enter version:" version
    then
             if [ ! -z $version ]
            then
                    go build -ldflags "-X main._version_=${version}" -tags=prod -o ${appName}
                    echo "Prod programe is success"
                    echo "App name is 【${appName}】"
                    echo "Version is 【$version】"
            else
                    echo "version is none"
            fi
    else
       echo "sorry. too slow!"
    fi
}

#后台运行应用
nohupApp(){
    nohup ./${appName} &
}

#查看应用运行状态
showStatus(){
    tail -f nohup.out
}


#平滑重启应用
restartApp(){
    ps aux | grep "${appName}" | grep -v grep | awk '{print $2}' | xargs -i kill -1 {}
}

#关闭应用
stopApp(){
    ps aux | grep "${appName}" | grep -v grep | awk '{print $2}' | xargs -i kill {}
}

#帮助
help(){
    echo "Usage:"
    echo ""
    echo "          ./app command [arguments]"
    echo ""
    echo "The commands are:"
    echo ""
    echo "          dev   [appName]"
    echo "          prod  [appName]"
    echo "          start   [appName]"
    echo "          restart [appName]"
    echo "          stop    [appName]"
    echo "          status  [appName]"
    echo "          -h      [appName]"
}

case ${option} in
        "dev")buildDevApp
        ;;
        "prod")buildProdApp
        ;;
        "start") nohupApp ; showStatus
        ;;
        "restart")restartApp
         ;;
        "stop")stopApp
        ;;
        "status")showStatus
        ;;
        "-h")help
        ;;
        *)help
        exit 1 #退出
         ;;
esac
