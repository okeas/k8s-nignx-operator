package test

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"testing"
)

// 测试正则匹配的路由功能

func TestProxyMatch(t *testing.T) {

	// 开启新的路由
	r := mux.NewRouter()
	// 加入两个路由规则
	// "/users/2334" => 通过
	// "/users/abc" => 不通过
	r.NewRoute().Path("/").Methods("GET")
	r.NewRoute().Path("/users/{id:\\d+}").Methods("GET", "POST", "DELETE", "OPTIONS")

	// 开始要匹配
	match := &mux.RouteMatch{}

	// 手动创建一个request对象
	req := &http.Request{
		URL: &url.URL{Path: "/"},
		Method: "GET",
	}
	fmt.Println(r.Match(req, match))

}
