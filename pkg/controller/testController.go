package controller

import "github.com/pinguo/pgo2"

type TestController struct {
	pgo2.Controller
}

func (t *TestController) ActionIndex() {
	t.Json("hello world", 200, "success")
}
