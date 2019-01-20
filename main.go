package main

import (
	"github.com/astaxie/beego"
	"go-pic-dir/base/cache/redis"
	"go-pic-dir/base/commands"
	"go-pic-dir/models"
	_ "go-pic-dir/routers"
)

func main() {
   //日志配置
   beego.SetLevel(beego.LevelDebug)
   beego.SetLogger("file", `{"filename":"storage/logs/log.log"}`)

   //注册DB
   models.RegisterDB()

   //初始缓存
   redis.Init()

   //执行任务
   commands.RegisterTask()

   beego.Run()

}



