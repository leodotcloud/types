package v3public

import (
	"context"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	ActiveDirectoryProviderGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ActiveDirectoryProvider",
	}
	ActiveDirectoryProviderResource = metav1.APIResource{
		Name:         "activedirectoryproviders",
		SingularName: "activedirectoryprovider",
		Namespaced:   false,
		Kind:         ActiveDirectoryProviderGroupVersionKind.Kind,
	}
)

func NewActiveDirectoryProvider(namespace, name string, obj ActiveDirectoryProvider) *ActiveDirectoryProvider {
	obj.APIVersion, obj.Kind = ActiveDirectoryProviderGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type ActiveDirectoryProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActiveDirectoryProvider
}

type ActiveDirectoryProviderHandlerFunc func(key string, obj *ActiveDirectoryProvider) (runtime.Object, error)

type ActiveDirectoryProviderChangeHandlerFunc func(obj *ActiveDirectoryProvider) (runtime.Object, error)

type ActiveDirectoryProviderLister interface {
	List(namespace string, selector labels.Selector) (ret []*ActiveDirectoryProvider, err error)
	Get(namespace, name string) (*ActiveDirectoryProvider, error)
}

type ActiveDirectoryProviderController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() ActiveDirectoryProviderLister
	AddHandler(ctx context.Context, name string, handler ActiveDirectoryProviderHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler ActiveDirectoryProviderHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ActiveDirectoryProviderInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*ActiveDirectoryProvider) (*ActiveDirectoryProvider, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ActiveDirectoryProvider, error)
	Get(name string, opts metav1.GetOptions) (*ActiveDirectoryProvider, error)
	Update(*ActiveDirectoryProvider) (*ActiveDirectoryProvider, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ActiveDirectoryProviderList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ActiveDirectoryProviderController
	AddHandler(ctx context.Context, name string, sync ActiveDirectoryProviderHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle ActiveDirectoryProviderLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync ActiveDirectoryProviderHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle ActiveDirectoryProviderLifecycle)
}

type activeDirectoryProviderLister struct {
	controller *activeDirectoryProviderController
}

func (l *activeDirectoryProviderLister) List(namespace string, selector labels.Selector) (ret []*ActiveDirectoryProvider, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*ActiveDirectoryProvider))
	})
	return
}

func (l *activeDirectoryProviderLister) Get(namespace, name string) (*ActiveDirectoryProvider, error) {
	var key string
	if namespace != "" {
		key = namespace + "/" + name
	} else {
		key = name
	}
	obj, exists, err := l.controller.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schema.GroupResource{
			Group:    ActiveDirectoryProviderGroupVersionKind.Group,
			Resource: "activeDirectoryProvider",
		}, key)
	}
	return obj.(*ActiveDirectoryProvider), nil
}

type activeDirectoryProviderController struct {
	controller.GenericController
}

func (c *activeDirectoryProviderController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *activeDirectoryProviderController) Lister() ActiveDirectoryProviderLister {
	return &activeDirectoryProviderLister{
		controller: c,
	}
}

func (c *activeDirectoryProviderController) AddHandler(ctx context.Context, name string, handler ActiveDirectoryProviderHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*ActiveDirectoryProvider); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *activeDirectoryProviderController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler ActiveDirectoryProviderHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*ActiveDirectoryProvider); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type activeDirectoryProviderFactory struct {
}

func (c activeDirectoryProviderFactory) Object() runtime.Object {
	return &ActiveDirectoryProvider{}
}

func (c activeDirectoryProviderFactory) List() runtime.Object {
	return &ActiveDirectoryProviderList{}
}

func (s *activeDirectoryProviderClient) Controller() ActiveDirectoryProviderController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.activeDirectoryProviderControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ActiveDirectoryProviderGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &activeDirectoryProviderController{
		GenericController: genericController,
	}

	s.client.activeDirectoryProviderControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type activeDirectoryProviderClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   ActiveDirectoryProviderController
}

func (s *activeDirectoryProviderClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *activeDirectoryProviderClient) Create(o *ActiveDirectoryProvider) (*ActiveDirectoryProvider, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ActiveDirectoryProvider), err
}

func (s *activeDirectoryProviderClient) Get(name string, opts metav1.GetOptions) (*ActiveDirectoryProvider, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ActiveDirectoryProvider), err
}

func (s *activeDirectoryProviderClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ActiveDirectoryProvider, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*ActiveDirectoryProvider), err
}

func (s *activeDirectoryProviderClient) Update(o *ActiveDirectoryProvider) (*ActiveDirectoryProvider, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ActiveDirectoryProvider), err
}

func (s *activeDirectoryProviderClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *activeDirectoryProviderClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *activeDirectoryProviderClient) List(opts metav1.ListOptions) (*ActiveDirectoryProviderList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ActiveDirectoryProviderList), err
}

func (s *activeDirectoryProviderClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *activeDirectoryProviderClient) Patch(o *ActiveDirectoryProvider, patchType types.PatchType, data []byte, subresources ...string) (*ActiveDirectoryProvider, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*ActiveDirectoryProvider), err
}

func (s *activeDirectoryProviderClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *activeDirectoryProviderClient) AddHandler(ctx context.Context, name string, sync ActiveDirectoryProviderHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *activeDirectoryProviderClient) AddLifecycle(ctx context.Context, name string, lifecycle ActiveDirectoryProviderLifecycle) {
	sync := NewActiveDirectoryProviderLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *activeDirectoryProviderClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync ActiveDirectoryProviderHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *activeDirectoryProviderClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle ActiveDirectoryProviderLifecycle) {
	sync := NewActiveDirectoryProviderLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

type ActiveDirectoryProviderIndexer func(obj *ActiveDirectoryProvider) ([]string, error)

type ActiveDirectoryProviderClientCache interface {
	Get(namespace, name string) (*ActiveDirectoryProvider, error)
	List(namespace string, selector labels.Selector) ([]*ActiveDirectoryProvider, error)

	Index(name string, indexer ActiveDirectoryProviderIndexer)
	GetIndexed(name, key string) ([]*ActiveDirectoryProvider, error)
}

type ActiveDirectoryProviderClient interface {
	Create(*ActiveDirectoryProvider) (*ActiveDirectoryProvider, error)
	Get(namespace, name string, opts metav1.GetOptions) (*ActiveDirectoryProvider, error)
	Update(*ActiveDirectoryProvider) (*ActiveDirectoryProvider, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*ActiveDirectoryProviderList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() ActiveDirectoryProviderClientCache

	OnCreate(ctx context.Context, name string, sync ActiveDirectoryProviderChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync ActiveDirectoryProviderChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync ActiveDirectoryProviderChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() ActiveDirectoryProviderInterface
}

type activeDirectoryProviderClientCache struct {
	client *activeDirectoryProviderClient2
}

type activeDirectoryProviderClient2 struct {
	iface      ActiveDirectoryProviderInterface
	controller ActiveDirectoryProviderController
}

func (n *activeDirectoryProviderClient2) Interface() ActiveDirectoryProviderInterface {
	return n.iface
}

func (n *activeDirectoryProviderClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *activeDirectoryProviderClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *activeDirectoryProviderClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *activeDirectoryProviderClient2) Create(obj *ActiveDirectoryProvider) (*ActiveDirectoryProvider, error) {
	return n.iface.Create(obj)
}

func (n *activeDirectoryProviderClient2) Get(namespace, name string, opts metav1.GetOptions) (*ActiveDirectoryProvider, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *activeDirectoryProviderClient2) Update(obj *ActiveDirectoryProvider) (*ActiveDirectoryProvider, error) {
	return n.iface.Update(obj)
}

func (n *activeDirectoryProviderClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *activeDirectoryProviderClient2) List(namespace string, opts metav1.ListOptions) (*ActiveDirectoryProviderList, error) {
	return n.iface.List(opts)
}

func (n *activeDirectoryProviderClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *activeDirectoryProviderClientCache) Get(namespace, name string) (*ActiveDirectoryProvider, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *activeDirectoryProviderClientCache) List(namespace string, selector labels.Selector) ([]*ActiveDirectoryProvider, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *activeDirectoryProviderClient2) Cache() ActiveDirectoryProviderClientCache {
	n.loadController()
	return &activeDirectoryProviderClientCache{
		client: n,
	}
}

func (n *activeDirectoryProviderClient2) OnCreate(ctx context.Context, name string, sync ActiveDirectoryProviderChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &activeDirectoryProviderLifecycleDelegate{create: sync})
}

func (n *activeDirectoryProviderClient2) OnChange(ctx context.Context, name string, sync ActiveDirectoryProviderChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &activeDirectoryProviderLifecycleDelegate{update: sync})
}

func (n *activeDirectoryProviderClient2) OnRemove(ctx context.Context, name string, sync ActiveDirectoryProviderChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &activeDirectoryProviderLifecycleDelegate{remove: sync})
}

func (n *activeDirectoryProviderClientCache) Index(name string, indexer ActiveDirectoryProviderIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*ActiveDirectoryProvider); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *activeDirectoryProviderClientCache) GetIndexed(name, key string) ([]*ActiveDirectoryProvider, error) {
	var result []*ActiveDirectoryProvider
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*ActiveDirectoryProvider); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *activeDirectoryProviderClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type activeDirectoryProviderLifecycleDelegate struct {
	create ActiveDirectoryProviderChangeHandlerFunc
	update ActiveDirectoryProviderChangeHandlerFunc
	remove ActiveDirectoryProviderChangeHandlerFunc
}

func (n *activeDirectoryProviderLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *activeDirectoryProviderLifecycleDelegate) Create(obj *ActiveDirectoryProvider) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *activeDirectoryProviderLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *activeDirectoryProviderLifecycleDelegate) Remove(obj *ActiveDirectoryProvider) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *activeDirectoryProviderLifecycleDelegate) Updated(obj *ActiveDirectoryProvider) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
