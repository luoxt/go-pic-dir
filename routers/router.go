package routers

import (
	"go-pic-dir/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.TuController{}, "*:Index")

}
