package k8sconfig

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// 第一步：实现crd对象

type RouteSpec struct {
	Version string `json:"version,omitempty"`
}
// Route 和crd.yaml中的kind 保持一致。
type Route struct {
	// k8s基本的obj的字段。
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec RouteSpec `json:"spec,omitempty"`
}
// RouteList obj基本要有的对象
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items []Route `json:"items,omitempty"`
}

func init() {
	// Scheme defines methods for serializing and deserializing API objects
	SchemeBuilder.Register(func(scheme *runtime.Scheme) error {

		gv:=schema.GroupVersion{
			Group:"extensions.jtthink.com",
			Version: "v1",
		}
		// 加入自定义的两个类型。
		scheme.AddKnownTypes(gv,&Route{},&RouteList{})
		metav1.AddToGroupVersion(scheme,gv)
		return nil
	})
}

