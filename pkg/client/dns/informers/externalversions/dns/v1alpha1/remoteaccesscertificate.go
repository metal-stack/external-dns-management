// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0
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

// RemoteAccessCertificateInformer provides access to a shared informer and lister for
// RemoteAccessCertificates.
type RemoteAccessCertificateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.RemoteAccessCertificateLister
}

type remoteAccessCertificateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRemoteAccessCertificateInformer constructs a new informer for RemoteAccessCertificate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRemoteAccessCertificateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRemoteAccessCertificateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRemoteAccessCertificateInformer constructs a new informer for RemoteAccessCertificate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRemoteAccessCertificateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DnsV1alpha1().RemoteAccessCertificates(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DnsV1alpha1().RemoteAccessCertificates(namespace).Watch(context.TODO(), options)
			},
		},
		&dnsv1alpha1.RemoteAccessCertificate{},
		resyncPeriod,
		indexers,
	)
}

func (f *remoteAccessCertificateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRemoteAccessCertificateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *remoteAccessCertificateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&dnsv1alpha1.RemoteAccessCertificate{}, f.defaultInformer)
}

func (f *remoteAccessCertificateInformer) Lister() v1alpha1.RemoteAccessCertificateLister {
	return v1alpha1.NewRemoteAccessCertificateLister(f.Informer().GetIndexer())
}
