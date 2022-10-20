package k8sconfig

import (
	"context"
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

	return reconcile.Result{}, nil
}

// InjectClient 依赖注入参数
func (r *MyProxyController) InjectClient(c client.Client) error {
	r.Client = c
	return nil
}
