/*
Copyright the Velero contributors.

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

package v1

import (
	"context"
	time "time"

	backupdriverv1 "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/apis/backupdriver/v1"
	versioned "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/generated/listers/backupdriver/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DeleteSnapshotInformer provides access to a shared informer and lister for
// DeleteSnapshots.
type DeleteSnapshotInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DeleteSnapshotLister
}

type deleteSnapshotInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDeleteSnapshotInformer constructs a new informer for DeleteSnapshot type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDeleteSnapshotInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDeleteSnapshotInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDeleteSnapshotInformer constructs a new informer for DeleteSnapshot type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDeleteSnapshotInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BackupdriverV1().DeleteSnapshots(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BackupdriverV1().DeleteSnapshots(namespace).Watch(context.TODO(), options)
			},
		},
		&backupdriverv1.DeleteSnapshot{},
		resyncPeriod,
		indexers,
	)
}

func (f *deleteSnapshotInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDeleteSnapshotInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *deleteSnapshotInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&backupdriverv1.DeleteSnapshot{}, f.defaultInformer)
}

func (f *deleteSnapshotInformer) Lister() v1.DeleteSnapshotLister {
	return v1.NewDeleteSnapshotLister(f.Informer().GetIndexer())
}
