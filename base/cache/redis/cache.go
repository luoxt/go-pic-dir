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

	//if val == nil {
	//	return e, nil
	//}
	//if b, ok := val.([]byte); ok {
	//	buf := bytes.NewBuffer(b)
	//
	//	decoder := gob.NewDecoder(buf)
	//	err := decoder.Decode(e)
	//
	//	if err != nil {
	//		beego.Error("反序列化对象失败 ->", err)
	//	}
	//	return e, err
	//} else if s, ok := val.(string); ok && s != "" {
	//
	//	buf := bytes.NewBufferString(s)
	//
	//	decoder := gob.NewDecoder(buf)
	//
	//	err := decoder.Decode(e)
	//
	//	if err != nil {
	//		beego.Error("反序列化对象失败 ->", err)
	//	}
	//	return e, err
	//}
	//return e, nil
}


func Put(key string, val interface{}, timeout time.Duration) error {

	//var buf bytes.Buffer
	//
	//encoder := gob.NewEncoder(&buf)
	//
	//err := encoder.Encode(val)
	//if err != nil {
	//	beego.Error("序列化对象失败 ->", err)
	//	return err
	//}

	//return bm.Put(key, buf.String(), timeout)
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

	redis, err := cache.NewCache(cacheType, `{"conn":"`+cacheHost+`:`+cachePort+`", "key":"SSORedis"}`)
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


