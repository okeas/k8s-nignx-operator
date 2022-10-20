package filters

import (
	"github.com/valyala/fasthttp"
	"reflect"
)

const AnnotationPrefix = "jtthink.ingress.kubernetes.io"

//所有过滤器 的接口
type ProxyFilter interface {
	SetPath(path string)  //用来设置  path的设置（带正则支持)-----并不是所有过滤器都要用到
	SetValue(values ...string)	// 用来设置
	Do(ctx *fasthttp.RequestCtx)
}

type ProxyFilters []ProxyFilter

func(p ProxyFilters) SetPath(path string) {
	for _,filter := range p {
		filter.SetPath(path)
	}
}

func(p ProxyFilters) Do(ctx *fasthttp.RequestCtx){
	for _, filter := range p {
		filter.Do(ctx)
	}
}

//针对Request
var FilterList = map[string]ProxyFilter{}

//针对Response
var FilterListResponse = map[string]ProxyFilter{}

//注册过滤器(request)
func registerFilter(key string ,filter ProxyFilter)  {
	FilterList[key]= filter
}

//注册过滤器(response)
func registerResponseFilter(key  string ,filter ProxyFilter)  {
	FilterListResponse[key]= filter
}

func init() {

}
//检查注解是否 和预设的 过滤器 匹配
func CheckAnnotations(annos map[string]string, isrespone bool, exts ...string) []ProxyFilter{
	filters := []ProxyFilter{}

	var list map[string]ProxyFilter
	if isrespone {
		list = FilterListResponse
	} else {
		list = FilterList
	}


	for annoKey, annoValue := range annos{
		for filterKey, filterReflect := range list{
			if annoKey == filterKey{
				t := reflect.TypeOf(filterReflect)
				if t.Kind() == reflect.Ptr {
					t=t.Elem()
				}
				filter := reflect.New(t).Interface().(ProxyFilter)
				params := []string{annoValue}
				params = append(params,exts...)
				filter.SetValue(params...)
				filters = append(filters,filter)
			}
		}
	}
	return filters
}

