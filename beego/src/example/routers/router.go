package routers

import (
	"example/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.CompanyRead{})
}
