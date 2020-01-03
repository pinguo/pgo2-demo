package controller

import (
    "fmt"
    "time"

    "github.com/pinguo/pgo2"
    "github.com/pinguo/pgo2/adapter"
)

type MemoryController struct {
    pgo2.Controller
}

// curl -v http://127.0.0.1:8000/memory/set
func (m *MemoryController) ActionSet() {
    // 获取memory的上下文适配对象
    mm := m.GetObj(adapter.NewMemory()).(*adapter.Memory)

    // 设置用户输入值
    key := m.Context().ValidateParam("key", "test_key1").Do()
    val := m.Context().ValidateParam("val", "test_val1").Do()
    mm.Set(key, val)

    // 设置自定义过期时间
    mm.Set("test_key2", 100, 2*time.Minute)

    // 设置map值，会自动进行json序列化
    data := map[string]interface{}{"f1": 100, "f2": true, "f3": "hello"}
    mm.Set("test_key3", data)

    // 并行设置多个key
    items := map[string]interface{}{
        "test_key4": []int{1, 2, 3, 4},
        "test_key5": "test_val5",
        "test_key6": map[string]interface{}{"f61": 11, "f62": "hello"},
    }
    mm.MSet(items)
}

// curl -v http://127.0.0.1:8000/memory/get
func (m *MemoryController) ActionGet() {
    // 从对象池获取memory的上下文适配对象
    mm := m.GetObjPool(adapter.MemoryClass, adapter.NewMemoryPool).(*adapter.Memory)

    // 获取string
    if val := mm.Get("test_key1"); val != nil {
        fmt.Println("value of test_key1:", val.String())
    }

    // 获取int
    if val := mm.Get("test_key2"); val != nil {
        fmt.Println("value of test_key2:", val.Int())
    }

    // 获取序列化的数据
    if val := mm.Get("test_key3"); val != nil {
        var data map[string]interface{}
        val.Decode(&data)
        fmt.Println("value of test_key3:", data)
    }

    // 获取多个key
    if res := mm.MGet([]string{"test_key4", "test_key5", "test_key6"}); res != nil {
        for key, val := range res {
            fmt.Printf("value of %s: %v\n", key, val.String())
        }
    }
}
