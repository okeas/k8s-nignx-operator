package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"golanglearning/new_project/k8s-operator-practice/pkg/config"
	"golanglearning/new_project/k8s-operator-practice/pkg/filters"
)

//实现简易反向代理。
func ProxyHandler(ctx *fasthttp.RequestCtx) {

	// 匹配到path
	if getProxy := config.GetRoute(ctx.Request) ; getProxy != nil{
		filters.ProxyFilters(getProxy.RequestFilters).Do(ctx) //过滤
		getProxy.Proxy.ServeHTTP(ctx) //反代
		filters.ProxyFilters(getProxy.ResponseFilters).Do(ctx) //过滤
	} else {
		ctx.Response.SetStatusCode(404)
		ctx.Response.SetBodyString("404....")
	}

	// 测试用
	//jtthink.ServeHTTP(ctx)
	//// 可以修改响应值，更改头部
	//ctx.Response.Header.Add("myname", "test")
}

// 实际用的网址
//var jtthink = proxy.NewReverseProxy("www.jtthink.com")

func main() {

	config.InitConfig()


	// http.ListenAndServe
	//fasthttp.ListenAndServe(":80", ProxyHandler)	// 反向代理服务器
	fasthttp.ListenAndServe(fmt.Sprintf(":%d", config.SysConfig.Server.Port), ProxyHandler)

}