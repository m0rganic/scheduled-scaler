/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	scaling_v1alpha1 "k8s.restdev.com/operators/pkg/apis/scaling/v1alpha1"
	versioned "k8s.restdev.com/operators/pkg/client/clientset/versioned"
	internalinterfaces "k8s.restdev.com/operators/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "k8s.restdev.com/operators/pkg/client/listers/scaling/v1alpha1"
)

// ScheduledScalerInformer provides access to a shared informer and lister for
// ScheduledScalers.
type ScheduledScalerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ScheduledScalerLister
}

type scheduledScalerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewScheduledScalerInformer constructs a new informer for ScheduledScaler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewScheduledScalerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredScheduledScalerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredScheduledScalerInformer constructs a new informer for ScheduledScaler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredScheduledScalerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ScalingV1alpha1().ScheduledScalers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ScalingV1alpha1().ScheduledScalers(namespace).Watch(options)
			},
		},
		&scaling_v1alpha1.ScheduledScaler{},
		resyncPeriod,
		indexers,
	)
}

func (f *scheduledScalerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredScheduledScalerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *scheduledScalerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&scaling_v1alpha1.ScheduledScaler{}, f.defaultInformer)
}

func (f *scheduledScalerInformer) Lister() v1alpha1.ScheduledScalerLister {
	return v1alpha1.NewScheduledScalerLister(f.Informer().GetIndexer())
}
