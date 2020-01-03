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
	e.Json(pgo2.EmptyObject,status, "Controller.Error " + message)
}