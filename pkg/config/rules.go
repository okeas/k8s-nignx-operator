package config

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/valyala/fasthttp"
	proxy "github.com/yeqown/fasthttp-reverse-proxy/v2"
	"golanglearning/new_project/k8s-operator-practice/pkg/filters"
	v1 "k8s.io/api/networking/v1"
	"net/http"
	"net/url"
)

type ProxyHandler struct {
	Proxy *proxy.ReverseProxy
	//Filters []filters.ProxyFilter
	RequestFilters []filters.ProxyFilter
	ResponseFilters []filters.ProxyFilter
}


func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

//var MyRouter *mux.Router
//
//func init() {
//	MyRouter = mux.NewRouter()
//}


// ParseRule 解析配置文建中的rules，初始化路由
func ParseRule() {
	//for _, rule := range SysConfig.Ingress.Rules {
	//
	//	for _, path := range rule.HTTP.Paths {
	//		// 构建出proxy 反向代理对象
	//		rProxy := proxy.NewReverseProxy(
	//			fmt.Sprintf("%s", path.Backend.Service.Name))
	//		fmt.Println(fmt.Sprintf("%s", path.Backend.Service.Name))
	//		// 把path加入router中
	//		if path.PathType != nil && *path.PathType == v1.PathTypeExact {
	//			// 精确匹配
	//			MyRouter.NewRoute().Path(path.Path).Methods("GET", "POST", "DELETE", "OPTIONS").
	//				Handler(&ProxyHandler{Proxy: rProxy})
	//		} else {
	//			MyRouter.NewRoute().Path(path.Path).Methods("GET", "POST", "DELETE", "OPTIONS").
	//				Handler(&ProxyHandler{Proxy: rProxy})
	//		}
	//
	//	}
	//}
	// 遍历
	for _, ingress := range SysConfig.Ingress {
		for _, rule := range ingress.Spec.Rules {

			for _, path := range rule.HTTP.Paths {
				// 构建出proxy 反向代理对象
				rProxy := proxy.NewReverseProxy(
					fmt.Sprintf("%s", path.Backend.Service.Name))
				fmt.Println(fmt.Sprintf("%d", path.Backend.Service.Port.Number))
				// 把path加入router中
				routeBud:=NewRouteBuilder()

				routeBud.
					SetPath(path.Path,path.PathType!=nil && *path.PathType==v1.PathTypeExact).
					SetHost(rule.Host,rule.Host!="").
					Build(&ProxyHandler{
						Proxy: rProxy,
						//Filters:filters.CheckAnnotations(ingress.Annotations,path.Path),
						RequestFilters:filters.CheckAnnotations(ingress.Annotations,false),
						ResponseFilters:filters.CheckAnnotations(ingress.Annotations,true),
					})
				//if path.PathType != nil && *path.PathType == v1.PathTypeExact {
				//	// 精确匹配
				//	MyRouter.NewRoute().Path(path.Path).Methods("GET", "POST", "DELETE", "OPTIONS").
				//		Handler(&ProxyHandler{Proxy: rProxy})
				//} else {
				//	MyRouter.NewRoute().Path(path.Path).Methods("GET", "POST", "DELETE", "OPTIONS").
				//		Handler(&ProxyHandler{Proxy: rProxy})
				//}

			}

		}
	}
}

// 获取路由， (先匹配 请求path，如果匹配到，会返回 对应的proxy对象)
func GetRoute(request fasthttp.Request) *ProxyHandler {

	match := &mux.RouteMatch{}
	httpRequest := &http.Request{
		URL: &url.URL{Path: string(request.URI().Path())},
		Method: string(request.Header.Method()),
		Host: string(request.Header.Host()),
	}

	if MyRouter.Match(httpRequest, match) {	// 一旦匹配到

		proxyHandler:=match.Handler.(*ProxyHandler)
		pathExp,err:=match.Route.GetPathRegexp()  //对过滤器 塞值：  path
		// 譬如这样：^/users/(?P<v0>[^/]+)

		if err == nil {
			//filters.ProxyFilters(proxyHandler.Filters).SetPath(pathExp)
			filters.ProxyFilters(proxyHandler.RequestFilters).SetPath(pathExp)
			filters.ProxyFilters(proxyHandler.ResponseFilters).SetPath(pathExp)
		}
		return proxyHandler
	}

	return nil

}
