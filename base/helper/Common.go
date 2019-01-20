package helper

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"os"
	"reflect"
	"time"
)

func Gtime() int  {
	return time.Now().Second()
}

func Gdate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//创建目录
func Mkdir(path string) {

	codePath := path
	_, err := os.Stat(codePath)
	if os.IsNotExist(err) {
		os.Mkdir(codePath, os.ModePerm)
	}
}

func Gtype(a interface{}) interface{} {
	return reflect.TypeOf(a)
}

//组装JSON
type Rejson struct {
	Status string `json:"status"`
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data interface{}  `json:"data"`
}
func Gjson(status string, code string, msg string, data interface{}) *Rejson {

	re := &Rejson{
		status,
		code,
		msg,
		data,
	}
	//str, _ := json.Marshal(re)
	return re
}


//post
func Gpost(host string, header map[string]string, postdata map[string]interface{}) (map[string]interface{}, error)  {

	req := httplib.Post(host)

	for key, val := range header{
		req.Header(key, val)
	}

	for pkey, pval := range postdata{
		req.Param(pkey, pval.(string))
	}

	req.SetTimeout(100 * time.Second, 30 * time.Second)
	req.Response()

	var result map[string]interface{}
	req.ToJSON(&result)
	beego.Debug("【API接口调用】- URL:"+host, "REQUEST:", postdata, "RESPONE:", result)

	return result, nil

}

//调用平台接口
func Pfpost(ext_url string, postData map[string]interface{}) *Rejson {

	//url
	API_HOST := beego.AppConfig.String("API_HOST")
	host := API_HOST + ext_url

	//获取
	pfres, err:= Gpost(host, map[string]string{}, postData)
	if err != nil {
		return  Gjson("false", "400", "接口连接出错！", map[string]interface{}{})
	}
	var res map[string]interface{} = pfres

	////判断ret
	code := res["ret"].(float64)
	if code != 200 {
		return Gjson("false", "400", "接口数据出错！", res)
	}

	//判断data
	var data interface{}
	data = res["data"]

	//判断info
	info := data.(map[string]interface{})

	data_code := info["code"].(float64)
	data_msg := info["msg"].(string)

	if data_code != 0 {
		return Gjson("false", "401", data_msg, res)
	}
	return Gjson("true", "200", "获取成功", info["info"])

}

