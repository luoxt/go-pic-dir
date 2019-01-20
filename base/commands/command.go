package commands

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"go-pic-dir/models"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

//注册任务
func RegisterTask()  {
	Command:=beego.AppConfig.String("Command")
	if Command == "true" {

		 func() {
			//for {
				//收集数据
				Dotask("G:/MyData/")

				//处理数据
			 	//GetData()

			//}
		}()
	}
}

//调用任务
func Dotask(bpath string)  {
	//sec := 5*time.Second
	//time.Sleep(sec)
	filepath.Walk(bpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				fmt.Println("dir:", path)
				return nil
			}
			fmt.Println("file:", path)
			fmd5 := FileMd5(path)

			tu := models.NewTuPath()
			tu.Md5 = fmd5
			tu.Status = 0
			tu.TuName = filepath.Base(path)
			tu.TuPath = path
			id := tu.Inserts()

			fmt.Println(id)

			return nil
		})

}

func FileMd5(path string) string {
	fs, err := os.Open(path)
	if err != nil {
		return err.Error()
	}

	defer fs.Close()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, fs); err != nil {
		return err.Error()
	}

	return fmt.Sprintf("%x", md5hash.Sum(nil))
}

func GetData()  {
	tu := models.NewTuPath()
	tuList,_ := tu.Find("status", "1")

	for _, item :=range tuList {
		fmt.Println(item["tu_path"])

		oldPath := fmt.Sprintf("%s", item["tu_path"])

		basePath := "G:/pic/"
		tu_name := item["tu_name"]
		newPath := fmt.Sprintf("%s%s", basePath, tu_name)

		//fmt.Println(helper.Gtype(oldPath))
		//fmt.Println(oldPath)
		_, err := os.Stat(oldPath)
		if err == nil {
			os.Rename(oldPath, newPath)
		}

		ids := item["id"].(string)
		id, _ := strconv.ParseInt(ids, 10, 64)
		tu.Update(id)

	}
}

