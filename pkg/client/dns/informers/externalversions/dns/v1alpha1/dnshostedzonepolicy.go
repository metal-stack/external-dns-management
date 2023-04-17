/*
Copyright (c) 2023 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

	dnsv1alpha1 "github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
	versioned "github.com/gardener/external-dns-management/pkg/client/dns/clientset/versioned"
	internalinterfaces "github.com/gardener/external-dns-management/pkg/client/dns/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gardener/external-dns-management/pkg/client/dns/listers/dns/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DNSHostedZonePolicyInformer provides access to a shared informer and lister for
// DNSHostedZonePolicies.
type DNSHostedZonePolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DNSHostedZonePolicyLister
}

type dNSHostedZonePolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDNSHostedZonePolicyInformer constructs a new informer for DNSHostedZonePolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDNSHostedZonePolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDNSHostedZonePolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDNSHostedZonePolicyInformer constructs a new informer for DNSHostedZonePolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDNSHostedZonePolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DnsV1alpha1().DNSHostedZonePolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DnsV1alpha1().DNSHostedZonePolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&dnsv1alpha1.DNSHostedZonePolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *dNSHostedZonePolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDNSHostedZonePolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dNSHostedZonePolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&dnsv1alpha1.DNSHostedZonePolicy{}, f.defaultInformer)
}

func (f *dNSHostedZonePolicyInformer) Lister() v1alpha1.DNSHostedZonePolicyLister {
	return v1alpha1.NewDNSHostedZonePolicyLister(f.Informer().GetIndexer())
}
