package controller

import "github.com/pinguo/pgo2"

func init() {
	container := pgo2.App().Container()
	container.Bind(&TestController{})
}
