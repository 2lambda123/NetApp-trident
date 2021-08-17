// Copyright 2021 NetApp, Inc. All Rights Reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/netapp/trident/v21/persistent_store/crd/apis/netapp/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TridentBackendLister helps list TridentBackends.
type TridentBackendLister interface {
	// List lists all TridentBackends in the indexer.
	List(selector labels.Selector) (ret []*v1.TridentBackend, err error)
	// TridentBackends returns an object that can list and get TridentBackends.
	TridentBackends(namespace string) TridentBackendNamespaceLister
	TridentBackendListerExpansion
}

// tridentBackendLister implements the TridentBackendLister interface.
type tridentBackendLister struct {
	indexer cache.Indexer
}

// NewTridentBackendLister returns a new TridentBackendLister.
func NewTridentBackendLister(indexer cache.Indexer) TridentBackendLister {
	return &tridentBackendLister{indexer: indexer}
}

// List lists all TridentBackends in the indexer.
func (s *tridentBackendLister) List(selector labels.Selector) (ret []*v1.TridentBackend, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TridentBackend))
	})
	return ret, err
}

// TridentBackends returns an object that can list and get TridentBackends.
func (s *tridentBackendLister) TridentBackends(namespace string) TridentBackendNamespaceLister {
	return tridentBackendNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TridentBackendNamespaceLister helps list and get TridentBackends.
type TridentBackendNamespaceLister interface {
	// List lists all TridentBackends in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.TridentBackend, err error)
	// Get retrieves the TridentBackend from the indexer for a given namespace and name.
	Get(name string) (*v1.TridentBackend, error)
	TridentBackendNamespaceListerExpansion
}

// tridentBackendNamespaceLister implements the TridentBackendNamespaceLister
// interface.
type tridentBackendNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TridentBackends in the indexer for a given namespace.
func (s tridentBackendNamespaceLister) List(selector labels.Selector) (ret []*v1.TridentBackend, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TridentBackend))
	})
	return ret, err
}

// Get retrieves the TridentBackend from the indexer for a given namespace and name.
func (s tridentBackendNamespaceLister) Get(name string) (*v1.TridentBackend, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("tridentbackend"), name)
	}
	return obj.(*v1.TridentBackend), nil
}
