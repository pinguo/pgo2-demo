package controller

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"

    "github.com/pinguo/pgo2"
    "github.com/pinguo/pgo2/adapter"
    "github.com/pinguo/pgo2/client/phttp"
)

type TestController struct {
    pgo2.Controller

    str string
    arr []int
}

// curl -v http://127.0.0.1:8000/test/index
func (t *TestController) ActionIndex() {
    client := t.GetObj(adapter.NewHttp()).(*adapter.Http)

    response := client.Get("http://127.0.0.1:8000/welcome/index", nil, &phttp.Option{Timeout: 2 * time.Second})
    if response !=nil{
        content, _ := ioutil.ReadAll(response.Body)
        response.Body.Close()
        fmt.Println("get response: ", string(content))
    }else{
        fmt.Println("get response: ")
    }

    t.str = "9999999"
    t.arr = []int{1, 2, 3}
    t.Json("call /test/index, str:"+t.str, http.StatusOK)
}
