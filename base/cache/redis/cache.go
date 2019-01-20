package redis

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"time"
)

var bm cache.Cache

func Get(key string) interface{} {
	val := bm.Get(key)
	return val
}

func Put(key string, val interface{}, timeout time.Duration) error {
	bm.Delete(key)
	return bm.Put(key, val, timeout)
}

func Delete(key string) error {
	return bm.Delete(key)
}

func Incr(key string) error {
	return bm.Incr(key)
}

func Decr(key string) error {
	return bm.Decr(key)
}

func IsExist(key string) bool {
	return bm.IsExist(key)
}

func ClearAll() error {
	return bm.ClearAll()
}

func StartAndGC(config string) error {
	return bm.StartAndGC(config)
}

//连接redis
func Connect() cache.Cache  {
	cacheType := beego.AppConfig.String("CACHE_TYPE")
	cacheHost := beego.AppConfig.String("CACHE_HOST")
	cachePort := beego.AppConfig.String("CACHE_PORT")
	cacheKey := beego.AppConfig.String("CACHE_KEY")

	redis, err := cache.NewCache(cacheType, `{"conn":"`+cacheHost+`:`+cachePort+`", "key":"`+cacheKey+`"}`)
	if err != nil {
		beego.Error("【连接redis失败】", err)
	}
	return redis
}

//初始化redis
func Init() {
	beego.Debug("【初始化REDIS】")
	bm = Connect()
}


