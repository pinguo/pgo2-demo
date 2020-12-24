package service

import (
	"fmt"

	"github.com/pinguo/pgo2"
	"github.com/pinguo/pgo2/iface"
)

func NewLogicData() *LogicData{
	return &LogicData{}
}

type LogicData struct {
	pgo2.Object
	data string
}

func (l *LogicData) Show(){
	fmt.Println("LogicData show data:" + l.data)
}



type LogicData1 struct {
	pgo2.Object
	data string
}

func (l *LogicData1) Show(){
	fmt.Println("LogicData1 show data:" + l.data)
}

type LogicData2 struct {
	pgo2.Object
	data string
}

// 通过GetObjBox时自动调用
func (l *LogicData2) Prepare( params ...string){
	if len(params)>0 {
		l.data = params[0]
	}
}

func (l *LogicData2) Show(){
	fmt.Println("LogicData2 show data:" + l.data)
}

func NewLogicData3Single(params ...interface{}) iface.IObject{
	data := ""
	if len(params) > 0{
		data = params[0].(string)
	}
	return &LogicData3{data:data}
}

type LogicData3 struct {
	pgo2.Object
	data string
}

func (l *LogicData3) Show(){
	fmt.Println("LogicData3 show data:" + l.data)
}

type Demo struct {
	pgo2.Object
}

var LogicData1Class string
var LogicData2Class string
var LogicData3Class string
// 如果需要用到对象池，需要先绑定对象
func init (){
	container := pgo2.App().Container()
	LogicData1Class = container.Bind(&LogicData1{})
	LogicData2Class = container.Bind(&LogicData2{})
	LogicData3Class = container.Bind(&LogicData3{})
}

func (d *Demo) Index(){
	// 获取对象并注入上下文
	l:=d.GetObj(NewLogicData()).(*LogicData)
	l.Show()
	// 获取对象并注入自定义上下文
	l0:=d.GetObjCtx(d.Context().Copy(), NewLogicData()).(*LogicData)
	l0.Show()

	// 从对象池获取对象并注入上下文
	l1 := d.GetObjBox(LogicData1Class).(*LogicData1)
	l1.Show()
	// 从对象池获取对象并注入上下文,并初始化函数
	l2 := d.GetObjBox(LogicData2Class).(*LogicData2)
	l2.Show()
	// 从对象池获取对象并注入上下文,并初始化函数
	l3 := d.GetObjBoxCtx(d.Context().Copy(), LogicData2Class).(*LogicData2)
	l3.Show()

	// 获取单例对象并注入上下文
	l4:=d.GetObjSingle("logicData3",NewLogicData3Single,"dataStr").(*LogicData3)
	l4.Show()

	// 获取单例对象并注入自定义上下文
	l5:=d.GetObjSingleCtx(d.Context().Copy(), "logicData3-1",NewLogicData3Single,"dataStr").(*LogicData3)
	l5.Show()
}

func NewDemo() *Demo{
	return &Demo{}
}
