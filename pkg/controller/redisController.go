package controller

import (
    "fmt"
    "time"

    "github.com/pinguo/pgo2"
    "github.com/pinguo/pgo2/adapter"
)

type RedisController struct {
    pgo2.Controller
}

// curl -v http://127.0.0.1:8000/redis/set
func (r *RedisController) ActionSet() {
    // 获取redis的上下文适配对象
    redis := r.GetObj(adapter.NewRedis()).(*adapter.Redis)

    // 设置用户输入值
    key := r.Context().ValidateParam("key", "test_key1").Do()
    val := r.Context().ValidateParam("val", "test_val1").Do()
    ok := redis.Set(key, val)
    fmt.Println("redis set test_key1=test_val1:", ok)

    // 设置自定义过期时间
    ok = redis.Set("test_key2", 100, 2*time.Minute)
    fmt.Println("redis set test_key2=100:", ok)

    // 设置map值，会自动进行json序列化
    data := map[string]interface{}{"f1": 100, "f2": true, "f3": "hello"}
    ok = redis.Set("test_key3", data)
    fmt.Println("redis set test_key3=<map>:", ok)

    // 并行设置多个key
    items := map[string]interface{}{
        "test_key4": []int{1, 2, 3, 4},
        "test_key5": "test_val5",
        "test_key6": map[string]interface{}{"f61": 11, "f62": "hello"},
    }
    ok = redis.MSet(items)
    fmt.Println("redis mset test_key[4-6]:", ok)
}

// curl -v http://127.0.0.1:8000/redis/get
func (r *RedisController) ActionGet() {
    // 对象池获取redis的上下文适配对象
    redis := r.GetObjBox(adapter.RedisClass).(*adapter.Redis)

    // 获取string
    if val := redis.Get("test_key1"); val != nil {
        fmt.Println("value of test_key1:", val.String())
    }

    // 获取int
    if val := redis.Get("test_key2"); val != nil {
        fmt.Println("value of test_key2:", val.Int())
    }

    // 获取序列化的数据
    if val := redis.Get("test_key3"); val != nil {
        var data map[string]interface{}
        val.Decode(&data)
        fmt.Println("value of test_key3:", data)
    }

    // 获取多个key
    if res := redis.MGet([]string{"test_key4", "test_key5", "test_key6"}); res != nil {
        for key, val := range res {
            fmt.Printf("value of %s: %v\n", key, val.String())
        }
    }
}
