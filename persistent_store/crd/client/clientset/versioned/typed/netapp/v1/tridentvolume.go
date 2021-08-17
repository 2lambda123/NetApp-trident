// Copyright 2021 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/netapp/trident/v21/persistent_store/crd/apis/netapp/v1"
	scheme "github.com/netapp/trident/v21/persistent_store/crd/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TridentVolumesGetter has a method to return a TridentVolumeInterface.
// A group's client should implement this interface.
type TridentVolumesGetter interface {
	TridentVolumes(namespace string) TridentVolumeInterface
}

// TridentVolumeInterface has methods to work with TridentVolume resources.
type TridentVolumeInterface interface {
	Create(ctx context.Context, tridentVolume *v1.TridentVolume, opts metav1.CreateOptions) (*v1.TridentVolume, error)
	Update(ctx context.Context, tridentVolume *v1.TridentVolume, opts metav1.UpdateOptions) (*v1.TridentVolume, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.TridentVolume, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TridentVolumeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentVolume, err error)
	TridentVolumeExpansion
}

// tridentVolumes implements TridentVolumeInterface
type tridentVolumes struct {
	client rest.Interface
	ns     string
}

// newTridentVolumes returns a TridentVolumes
func newTridentVolumes(c *TridentV1Client, namespace string) *tridentVolumes {
	return &tridentVolumes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tridentVolume, and returns the corresponding tridentVolume object, and an error if there is any.
func (c *tridentVolumes) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.TridentVolume, err error) {
	result = &v1.TridentVolume{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridentvolumes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TridentVolumes that match those selectors.
func (c *tridentVolumes) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TridentVolumeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TridentVolumeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridentvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tridentVolumes.
func (c *tridentVolumes) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tridentvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tridentVolume and creates it.  Returns the server's representation of the tridentVolume, and an error, if there is any.
func (c *tridentVolumes) Create(ctx context.Context, tridentVolume *v1.TridentVolume, opts metav1.CreateOptions) (result *v1.TridentVolume, err error) {
	result = &v1.TridentVolume{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tridentvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentVolume).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tridentVolume and updates it. Returns the server's representation of the tridentVolume, and an error, if there is any.
func (c *tridentVolumes) Update(ctx context.Context, tridentVolume *v1.TridentVolume, opts metav1.UpdateOptions) (result *v1.TridentVolume, err error) {
	result = &v1.TridentVolume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tridentvolumes").
		Name(tridentVolume.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentVolume).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tridentVolume and deletes it. Returns an error if one occurs.
func (c *tridentVolumes) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridentvolumes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tridentVolumes) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridentvolumes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tridentVolume.
func (c *tridentVolumes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentVolume, err error) {
	result = &v1.TridentVolume{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tridentvolumes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
