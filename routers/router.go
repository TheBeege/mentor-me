package routers

import (
	"github.com/TheBeege/mentor-me/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
