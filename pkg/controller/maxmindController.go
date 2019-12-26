package controller

import (
    "net/http"

    "github.com/pinguo/pgo2"
    "github.com/pinguo/pgo2/adapter"
)

type MaxMindController struct {
    pgo2.Controller
}

// curl -v http://127.0.0.1:8000/max-mind/geo-by-ip
func (m *MaxMindController) ActionGeoByIp() {
    // 获取要解析的IP地址
    ip := m.Context().ValidateParam("ip", "182.150.28.13").Do()

    // 获取MaxMind的上下文件适配对象
    // mmd := m.GetObj(adapter.NewMaxMind()).(*adapter.MaxMind)
    // 从对象池获取MaxMind的上下文件适配对象
    mmd := m.GetObjPool(adapter.MaxMindClass, adapter.NewMaxMindPool).(*adapter.MaxMind)

    // 解析IP的geo信息
    geo := mmd.GeoByIp(ip)

    m.Json(geo, http.StatusOK)
}
