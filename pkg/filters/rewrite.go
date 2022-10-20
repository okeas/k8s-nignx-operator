package filters

import (
	"github.com/valyala/fasthttp"
	"log"
	"regexp"
)
const RewriteAnnotation = AnnotationPrefix + "/rewrite-target"
//const RewriteAnnotation="nginx.ingress.kubernetes.io/rewrite-target"

func init() {
	registerFilter(RewriteAnnotation,(*RewriteFilter)(nil) )
}

type RewriteFilter struct {
	pathValue string
	target string  //注解  值
	path string
}
//可变参数。第1个是 rewrie-target:的值 如 /$1  第2个是 原始的path 值 ，如/jtthink/{(.*)}
//func(r *RewriteFilter) SetValue(values ...string){
//	if len(values) != 2 {
//		return
//	}
//	r.target = values[0]
//	r.pathValue = values[1]
//	r.pathValue = strings.Replace(r.pathValue,"{","",-1)
//	r.pathValue = strings.Replace(r.pathValue,"}","",-1)
//}

func(r *RewriteFilter) SetPath(value string){
	r.pathValue = value
}

//可变参数。第1个是 rewrie-target:的值 如 /$1
func(r *RewriteFilter) SetValue(values ...string){
	r.target = values[0]
}

func(r *RewriteFilter) Do(ctx *fasthttp.RequestCtx){
	getUrl := ctx.Request.URI().String() //获取 请求URL  譬如  /jtthink/users
	reg,err := regexp.Compile(r.pathValue)
	if err != nil {
		log.Println(err)
		return
	}

	getUrl = reg.ReplaceAllString(getUrl,r.target)

	ctx.Request.SetRequestURI(getUrl)
	if err != nil{
		log.Println(err)
		return
	}

}
