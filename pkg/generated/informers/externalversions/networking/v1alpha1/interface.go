// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/karmada-io/karmada/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// MultiClusterIngresses returns a MultiClusterIngressInformer.
	MultiClusterIngresses() MultiClusterIngressInformer
	// MultiClusterServices returns a MultiClusterServiceInformer.
	MultiClusterServices() MultiClusterServiceInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// MultiClusterIngresses returns a MultiClusterIngressInformer.
func (v *version) MultiClusterIngresses() MultiClusterIngressInformer {
	return &multiClusterIngressInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MultiClusterServices returns a MultiClusterServiceInformer.
func (v *version) MultiClusterServices() MultiClusterServiceInformer {
	return &multiClusterServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
