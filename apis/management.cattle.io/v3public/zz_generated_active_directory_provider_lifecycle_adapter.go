package v3public

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type ActiveDirectoryProviderLifecycle interface {
	Create(obj *ActiveDirectoryProvider) (runtime.Object, error)
	Remove(obj *ActiveDirectoryProvider) (runtime.Object, error)
	Updated(obj *ActiveDirectoryProvider) (runtime.Object, error)
}

type activeDirectoryProviderLifecycleAdapter struct {
	lifecycle ActiveDirectoryProviderLifecycle
}

func (w *activeDirectoryProviderLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *activeDirectoryProviderLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *activeDirectoryProviderLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ActiveDirectoryProvider))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *activeDirectoryProviderLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ActiveDirectoryProvider))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *activeDirectoryProviderLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ActiveDirectoryProvider))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewActiveDirectoryProviderLifecycleAdapter(name string, clusterScoped bool, client ActiveDirectoryProviderInterface, l ActiveDirectoryProviderLifecycle) ActiveDirectoryProviderHandlerFunc {
	adapter := &activeDirectoryProviderLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ActiveDirectoryProvider) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
