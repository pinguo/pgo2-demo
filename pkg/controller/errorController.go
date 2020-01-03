package controller

import (
	"github.com/pinguo/pgo2"
)

func init(){
	container := pgo2.App().Container()
	pgo2.App().Router().SetErrorController(container.Bind(&ErrorController{}))
	// 设置是否覆盖 HTTP status code
	pgo2.App().Router().SetHttpStatus(true)
}

type ErrorController struct {
	pgo2.Controller
}

// 此函数必须有 ErrorController 遵循接口iface.IErrorController
func (e *ErrorController) Error(status int , message string){
	// e.Json(pgo2.EmptyObject,status, "Controller.Error " + message)
	// 可扩展
	switch status {
	case 404:
		e.error404(message)
	default:
		e.other(status,message)
	}
}

func (e *ErrorController) error404(message string){
	e.Json(pgo2.EmptyObject,404, "Controller.error404 " + message)
	// e.View("404.html",message)
}

func (e *ErrorController) other(status int, message string){
	e.Json(pgo2.EmptyObject,status, "Controller.other " + message)
	// e.View("other.html",message)
}