// Copyright 2021 NetApp, Inc. All Rights Reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/netapp/trident/v21/persistent_store/crd/apis/netapp/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TridentTransactionLister helps list TridentTransactions.
type TridentTransactionLister interface {
	// List lists all TridentTransactions in the indexer.
	List(selector labels.Selector) (ret []*v1.TridentTransaction, err error)
	// TridentTransactions returns an object that can list and get TridentTransactions.
	TridentTransactions(namespace string) TridentTransactionNamespaceLister
	TridentTransactionListerExpansion
}

// tridentTransactionLister implements the TridentTransactionLister interface.
type tridentTransactionLister struct {
	indexer cache.Indexer
}

// NewTridentTransactionLister returns a new TridentTransactionLister.
func NewTridentTransactionLister(indexer cache.Indexer) TridentTransactionLister {
	return &tridentTransactionLister{indexer: indexer}
}

// List lists all TridentTransactions in the indexer.
func (s *tridentTransactionLister) List(selector labels.Selector) (ret []*v1.TridentTransaction, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TridentTransaction))
	})
	return ret, err
}

// TridentTransactions returns an object that can list and get TridentTransactions.
func (s *tridentTransactionLister) TridentTransactions(namespace string) TridentTransactionNamespaceLister {
	return tridentTransactionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TridentTransactionNamespaceLister helps list and get TridentTransactions.
type TridentTransactionNamespaceLister interface {
	// List lists all TridentTransactions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.TridentTransaction, err error)
	// Get retrieves the TridentTransaction from the indexer for a given namespace and name.
	Get(name string) (*v1.TridentTransaction, error)
	TridentTransactionNamespaceListerExpansion
}

// tridentTransactionNamespaceLister implements the TridentTransactionNamespaceLister
// interface.
type tridentTransactionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TridentTransactions in the indexer for a given namespace.
func (s tridentTransactionNamespaceLister) List(selector labels.Selector) (ret []*v1.TridentTransaction, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TridentTransaction))
	})
	return ret, err
}

// Get retrieves the TridentTransaction from the indexer for a given namespace and name.
func (s tridentTransactionNamespaceLister) Get(name string) (*v1.TridentTransaction, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("tridenttransaction"), name)
	}
	return obj.(*v1.TridentTransaction), nil
}
