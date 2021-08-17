// Copyright 2021 NetApp, Inc. All Rights Reserved.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	netappv1 "github.com/netapp/trident/v21/operator/controllers/orchestrator/apis/netapp/v1"
	versioned "github.com/netapp/trident/v21/operator/controllers/orchestrator/client/clientset/versioned"
	internalinterfaces "github.com/netapp/trident/v21/operator/controllers/orchestrator/client/informers/externalversions/internalinterfaces"
	v1 "github.com/netapp/trident/v21/operator/controllers/orchestrator/client/listers/netapp/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TridentOrchestratorInformer provides access to a shared informer and lister for
// TridentOrchestrators.
type TridentOrchestratorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.TridentOrchestratorLister
}

type tridentOrchestratorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewTridentOrchestratorInformer constructs a new informer for TridentOrchestrator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTridentOrchestratorInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTridentOrchestratorInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredTridentOrchestratorInformer constructs a new informer for TridentOrchestrator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTridentOrchestratorInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TridentV1().TridentOrchestrators().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TridentV1().TridentOrchestrators().Watch(context.TODO(), options)
			},
		},
		&netappv1.TridentOrchestrator{},
		resyncPeriod,
		indexers,
	)
}

func (f *tridentOrchestratorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTridentOrchestratorInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tridentOrchestratorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&netappv1.TridentOrchestrator{}, f.defaultInformer)
}

func (f *tridentOrchestratorInformer) Lister() v1.TridentOrchestratorLister {
	return v1.NewTridentOrchestratorLister(f.Informer().GetIndexer())
}
