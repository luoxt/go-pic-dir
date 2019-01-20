#!/bin/bash
#检查箱码生成二维码服务是否正常
NUM=`ps -ef | grep go-pic-dir | grep -v grep | wc -l`
echo $NUM
if [ $NUM -ne 0 ]
    then
    echo "`date '+%Y-%m-%d %H:%M:%S'` go-pic-dir Server Is Runing!"
    curl localhost:8080/box/task
    echo "\n"
else
    echo "`date '+%Y-%m-%d %H:%M:%S'` go-pic-dir Server Is Starting!"
    setsid /data/wwwroot/qrcode.cha.ttyun.com/go-pic-dir &
fi

##打包
#bee pack -be GOOS=linux