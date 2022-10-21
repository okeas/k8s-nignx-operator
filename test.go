package main

import (
	"k8s-operator-practice/pkg/k8sconfig"
	v1 "k8s.io/api/networking/v1"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"

	"os"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)


/*
	manager 主要用来管理Controller Admission Webhook 包括：
	访问资源对象的client cache scheme 并提供依赖注入机制 优雅关闭机制

	operator = crd + controller + webhook
 */

func main() {
	logf.SetLogger(zap.New())
	var myLog = logf.Log.WithName("myProxy")


	mgr, err := manager.New(k8sconfig.K8sConfig(), manager.Options{})
	if err  != nil {
		myLog.Error(err, "unable to set up manager")
		os.Exit(1)
	}
	// 传入资源&v1.Ingress{}，也可以用crd
	err = builder.ControllerManagedBy(mgr).
		For(&v1.Ingress{}).Complete(k8sconfig.NewMyProxyController())

	//++ 注册进入序列化表
	err = k8sconfig.SchemeBuilder.AddToScheme(mgr.GetScheme())
	if err != nil {
		myLog.Error(err, "unable add schema")
		os.Exit(1)
	}

	if err != nil {
		myLog.Error(err, "unable to create manager")
		os.Exit(1)
	}

	if err = mgr.Start(signals.SetupSignalHandler()); err != nil {
		myLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

}
