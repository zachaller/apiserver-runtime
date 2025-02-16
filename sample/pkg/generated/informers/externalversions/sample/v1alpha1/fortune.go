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
	samplev1alpha1 "github.com/zachaller/apiserver-runtime/sample/pkg/apis/sample/v1alpha1"
	versioned "github.com/zachaller/apiserver-runtime/sample/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/zachaller/apiserver-runtime/sample/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/zachaller/apiserver-runtime/sample/pkg/generated/listers/sample/v1alpha1"
)

// FortuneInformer provides access to a shared informer and lister for
// Fortunes.
type FortuneInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FortuneLister
}

type fortuneInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFortuneInformer constructs a new informer for Fortune type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFortuneInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFortuneInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFortuneInformer constructs a new informer for Fortune type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFortuneInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SampleV1alpha1().Fortunes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SampleV1alpha1().Fortunes(namespace).Watch(context.TODO(), options)
			},
		},
		&samplev1alpha1.Fortune{},
		resyncPeriod,
		indexers,
	)
}

func (f *fortuneInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFortuneInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *fortuneInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&samplev1alpha1.Fortune{}, f.defaultInformer)
}

func (f *fortuneInformer) Lister() v1alpha1.FortuneLister {
	return v1alpha1.NewFortuneLister(f.Informer().GetIndexer())
}
