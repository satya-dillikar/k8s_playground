/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	burgerstoredevv1alpha1 "satya.com/burgerstore/pkg/apis/burgerstore.dev/v1alpha1"
	versioned "satya.com/burgerstore/pkg/generated/clientset/versioned"
	internalinterfaces "satya.com/burgerstore/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "satya.com/burgerstore/pkg/generated/listers/burgerstore.dev/v1alpha1"
)

// BurgerStoreInformer provides access to a shared informer and lister for
// BurgerStores.
type BurgerStoreInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BurgerStoreLister
}

type burgerStoreInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBurgerStoreInformer constructs a new informer for BurgerStore type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBurgerStoreInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBurgerStoreInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBurgerStoreInformer constructs a new informer for BurgerStore type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBurgerStoreInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BurgerstoreV1alpha1().BurgerStores(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BurgerstoreV1alpha1().BurgerStores(namespace).Watch(context.TODO(), options)
			},
		},
		&burgerstoredevv1alpha1.BurgerStore{},
		resyncPeriod,
		indexers,
	)
}

func (f *burgerStoreInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBurgerStoreInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *burgerStoreInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&burgerstoredevv1alpha1.BurgerStore{}, f.defaultInformer)
}

func (f *burgerStoreInformer) Lister() v1alpha1.BurgerStoreLister {
	return v1alpha1.NewBurgerStoreLister(f.Informer().GetIndexer())
}
