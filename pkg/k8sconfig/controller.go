package k8sconfig

import (
	"context"
	"fmt"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// MyProxyController 控制器
type MyProxyController struct {
	client.Client	// 客户端
}

func NewMyProxyController() *MyProxyController {
	return &MyProxyController{}
}

// Reconcile 资源发生变更，会监听到，并且对现在状态与期望状态进行处理
func (r *MyProxyController) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {

	//obj := &v1.Ingress{}	// 对象
	//err := r.Get(ctx,req.NamespacedName,obj)
	//
	//if err != nil{
	//	return reconcile.Result{}, err
	//}
	//// 捞到符合的Annotations 资源对象
	//if v, ok := obj.Annotations["kubernetes.io/ingress.class"];ok && v=="jtthink" {
	//	fmt.Println(obj)
	//}

	obj := &Route{}
	err := r.Get(ctx, req.NamespacedName, obj)
	if err != nil {
		return reconcile.Result{}, err
	}
	fmt.Println(obj)


	return reconcile.Result{}, nil
}

// InjectClient 依赖注入参数
func (r *MyProxyController) InjectClient(c client.Client) error {
	r.Client = c
	return nil
}
