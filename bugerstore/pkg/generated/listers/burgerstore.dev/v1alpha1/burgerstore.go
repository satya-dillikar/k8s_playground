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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "satya.com/burgerstore/pkg/apis/burgerstore.dev/v1alpha1"
)

// BurgerStoreLister helps list BurgerStores.
// All objects returned here must be treated as read-only.
type BurgerStoreLister interface {
	// List lists all BurgerStores in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BurgerStore, err error)
	// BurgerStores returns an object that can list and get BurgerStores.
	BurgerStores(namespace string) BurgerStoreNamespaceLister
	BurgerStoreListerExpansion
}

// burgerStoreLister implements the BurgerStoreLister interface.
type burgerStoreLister struct {
	indexer cache.Indexer
}

// NewBurgerStoreLister returns a new BurgerStoreLister.
func NewBurgerStoreLister(indexer cache.Indexer) BurgerStoreLister {
	return &burgerStoreLister{indexer: indexer}
}

// List lists all BurgerStores in the indexer.
func (s *burgerStoreLister) List(selector labels.Selector) (ret []*v1alpha1.BurgerStore, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BurgerStore))
	})
	return ret, err
}

// BurgerStores returns an object that can list and get BurgerStores.
func (s *burgerStoreLister) BurgerStores(namespace string) BurgerStoreNamespaceLister {
	return burgerStoreNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BurgerStoreNamespaceLister helps list and get BurgerStores.
// All objects returned here must be treated as read-only.
type BurgerStoreNamespaceLister interface {
	// List lists all BurgerStores in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BurgerStore, err error)
	// Get retrieves the BurgerStore from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BurgerStore, error)
	BurgerStoreNamespaceListerExpansion
}

// burgerStoreNamespaceLister implements the BurgerStoreNamespaceLister
// interface.
type burgerStoreNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BurgerStores in the indexer for a given namespace.
func (s burgerStoreNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BurgerStore, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BurgerStore))
	})
	return ret, err
}

// Get retrieves the BurgerStore from the indexer for a given namespace and name.
func (s burgerStoreNamespaceLister) Get(name string) (*v1alpha1.BurgerStore, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("burgerstore"), name)
	}
	return obj.(*v1alpha1.BurgerStore), nil
}
