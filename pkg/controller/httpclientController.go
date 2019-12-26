package controller

import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "mime/multipart"
    "net/http"
    "os"

    "github.com/pinguo/pgo2"
    "github.com/pinguo/pgo2/adapter"
    "github.com/pinguo/pgo2/client/phttp"
)

type HttpClientController struct {
    pgo2.Controller
}

// curl -v http://127.0.0.1:8000/http-client/send-query
func (h *HttpClientController) ActionSendQuery() {
    // 获取http的上下文适配对象
    httpClient := h.GetObj(adapter.NewHttp()).(*adapter.Http)

    // 简单GET请求
    url := "http://127.0.0.1:8000/welcome/index"
    if res := httpClient.Get(url, nil); res != nil {
        defer res.Body.Close()
        content, _ := ioutil.ReadAll(res.Body)
        fmt.Println("content 1:", string(content))
    }

    // 带参数GET请求
    params := map[string]interface{}{"p1": "v1", "p2": 10, "p3": 9.9}
    if res := httpClient.Get(url, params); res != nil {
        defer res.Body.Close()
        content, _ := ioutil.ReadAll(res.Body)
        fmt.Println("content 2:", string(content))
    }

    // 自定义cookie和header GET请求
    option := phttp.Option{}
    option.SetCookie("c1", "cv1")
    option.SetHeader("h1", "hv1")
    if res := httpClient.Get(url, nil, &option); res != nil {
        defer res.Body.Close()
        content, _ := ioutil.ReadAll(res.Body)
        fmt.Println("content 3:", string(content))
    }
}

// curl -v http://127.0.0.1:8000/http-client/send-form
func (h *HttpClientController) ActionSendForm() {
    // 从对象池获取http的上下文适配对象
    httpClient := h.GetObjPool(adapter.HttpClass, adapter.NewHttpPool).(*adapter.Http)

    // 发送POST请求
    url := "http://127.0.0.1:8000/welcome/index"
    form := map[string]interface{}{"p1": "v1", "p2": 10, "p3": 9.9}
    if res := httpClient.Post(url, form); res != nil {
        defer res.Body.Close()
        content, _ := ioutil.ReadAll(res.Body)
        fmt.Println("content 1:", string(content))
    }
}

// curl -v http://127.0.0.1:8000/http-client/send-file
func (h *HttpClientController) ActionSendFile() {
    // 获取http的自定义上下文适配对象
    newCtx := h.Context().Copy()
    defer newCtx.FinishGoLog() // 刷新新上下文的日志
    httpClient := h.GetObjCtx(newCtx, adapter.NewHttp()).(*adapter.Http)

    // 上传文件的POST请求
    url := "http://127.0.0.1:8000/welcome/index"
    body := bytes.Buffer{}
    writer := multipart.NewWriter(&body)

    // 创建文件form
    formFile, _ := writer.CreateFormFile("form_file", "test.png")

    // 读取文件内空填充表单
    fileHandle, _ := os.Open("test.png")
    io.Copy(formFile, fileHandle)
    fileHandle.Close()

    // 结束表单构造
    option := phttp.Option{}
    option.SetHeader("Content-Type", writer.FormDataContentType())
    writer.Close() // 发送前一定要关闭writer以写入结尾

    // 发送表单，接收响应
    if res := httpClient.Post(url, &body, &option); res != nil {
        defer res.Body.Close()
        content, _ := ioutil.ReadAll(res.Body)
        fmt.Println("content 1:", string(content))
    }
}

// curl -v http://127.0.0.1:8000/http-client/multi-request
func (h *HttpClientController) ActionMultiRequest() {
    // 获取http的上下文适配对象
    newCtx := h.Context().Copy()
    defer newCtx.FinishGoLog() // 刷新新上下文的日志
    httpClient := h.GetObjPoolCtx(newCtx,adapter.HttpClass,adapter.NewHttpPool ).(*adapter.Http)

    req1, _ := http.NewRequest("GET", "http://127.0.0.1:8000/welcome/index", nil)
    req2, _ := http.NewRequest("GET", "http://127.0.0.1:8000/welcome/index", nil)
    req3, _ := http.NewRequest("GET", "http://127.0.0.1:8000/welcome/index", nil)
    req4, _ := http.NewRequest("GET", "http://127.0.0.1:8000/welcome/index", nil)

    // 并行请求多个url
    requests := []*http.Request{req1, req2, req3, req4}
    responses := httpClient.DoMulti(requests)

    for k, res := range responses {
        content, _ := ioutil.ReadAll(res.Body)
        fmt.Printf("content of response %d: %s\n", k, string(content))
    }
}
