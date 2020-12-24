package service

import (
    "fmt"

    "github.com/pinguo/pgo2"
)

// 提供给调用NewWelcomePool的时候使用
var WelcomeClass string
func init(){
    WelcomeClass = pgo2.App().Container().Bind(&Welcome{})
}

type Welcome struct {
    pgo2.Object
    id string
}

// 获取对象
func NewWelcome() *Welcome{

    fmt.Printf("call in Service/NewWelcome set name NewWelcome-id\n")
    return &Welcome{id:"NewWelcome-id"}
}

// 通过GetObjBox时自动调用
func (w *Welcome) Prepare(id string){
    w.id = id
    fmt.Printf("call in service/Prepare set name"+w.id+"\n")
}

func (w *Welcome) SayHello(name string, age int, sex string) {
    fmt.Printf("call in  service/Welcome.SayHello, name:%s age:%d sex:%s\n", name, age, sex)
}


func (w *Welcome) ShowId() {
    fmt.Printf("call in  service/Welcome.ShowId, id:%s\n", w.id)
}