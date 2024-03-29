package k8sconfig

import "k8s.io/apimachinery/pkg/runtime"

var SchemeBuilder = &Builder{}

type Builder struct {
	runtime.SchemeBuilder
}

func(b *Builder) AddSceme(scheme *runtime.Scheme) error{
	return b.AddToScheme(scheme)
}