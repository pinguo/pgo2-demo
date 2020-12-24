package controller

import (
	"github.com/pinguo/pgo2"
	"github.com/pinguo/pgo2/adapter"
)

type Orm struct {
	pgo2.Controller
}

// curl -v http://127.0.0.1:8000/orm/exec
func (o *Orm) ActionExec() {
	// 获取db的上下文适配对象
	db:=o.GetObjBox(adapter.OrmClass).(adapter.IOrm)
	product := &Product{Name: o.Context().LogId(),Age:12}
	db.Create(&product)
	o.Context().PushLog("insertId",product.ID)

	result := &Product{}
	db.Model(&Product{}).Select("name").Where("id=?",product.ID).Take(&result)
	o.Context().PushLog("name",product.Name)

	o.JsonV2(nil,200)
}

type Product struct {
	ID           uint
	Name         string
	Age          uint8
}