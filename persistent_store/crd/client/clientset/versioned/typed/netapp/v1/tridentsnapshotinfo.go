// Copyright 2023 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	scheme "github.com/netapp/trident/persistent_store/crd/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TridentSnapshotInfosGetter has a method to return a TridentSnapshotInfoInterface.
// A group's client should implement this interface.
type TridentSnapshotInfosGetter interface {
	TridentSnapshotInfos(namespace string) TridentSnapshotInfoInterface
}

// TridentSnapshotInfoInterface has methods to work with TridentSnapshotInfo resources.
type TridentSnapshotInfoInterface interface {
	Create(ctx context.Context, tridentSnapshotInfo *v1.TridentSnapshotInfo, opts metav1.CreateOptions) (*v1.TridentSnapshotInfo, error)
	Update(ctx context.Context, tridentSnapshotInfo *v1.TridentSnapshotInfo, opts metav1.UpdateOptions) (*v1.TridentSnapshotInfo, error)
	UpdateStatus(ctx context.Context, tridentSnapshotInfo *v1.TridentSnapshotInfo, opts metav1.UpdateOptions) (*v1.TridentSnapshotInfo, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.TridentSnapshotInfo, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TridentSnapshotInfoList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentSnapshotInfo, err error)
	TridentSnapshotInfoExpansion
}

// tridentSnapshotInfos implements TridentSnapshotInfoInterface
type tridentSnapshotInfos struct {
	client rest.Interface
	ns     string
}

// newTridentSnapshotInfos returns a TridentSnapshotInfos
func newTridentSnapshotInfos(c *TridentV1Client, namespace string) *tridentSnapshotInfos {
	return &tridentSnapshotInfos{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tridentSnapshotInfo, and returns the corresponding tridentSnapshotInfo object, and an error if there is any.
func (c *tridentSnapshotInfos) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.TridentSnapshotInfo, err error) {
	result = &v1.TridentSnapshotInfo{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridentsnapshotinfos").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TridentSnapshotInfos that match those selectors.
func (c *tridentSnapshotInfos) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TridentSnapshotInfoList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TridentSnapshotInfoList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridentsnapshotinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tridentSnapshotInfos.
func (c *tridentSnapshotInfos) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tridentsnapshotinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tridentSnapshotInfo and creates it.  Returns the server's representation of the tridentSnapshotInfo, and an error, if there is any.
func (c *tridentSnapshotInfos) Create(ctx context.Context, tridentSnapshotInfo *v1.TridentSnapshotInfo, opts metav1.CreateOptions) (result *v1.TridentSnapshotInfo, err error) {
	result = &v1.TridentSnapshotInfo{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tridentsnapshotinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentSnapshotInfo).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tridentSnapshotInfo and updates it. Returns the server's representation of the tridentSnapshotInfo, and an error, if there is any.
func (c *tridentSnapshotInfos) Update(ctx context.Context, tridentSnapshotInfo *v1.TridentSnapshotInfo, opts metav1.UpdateOptions) (result *v1.TridentSnapshotInfo, err error) {
	result = &v1.TridentSnapshotInfo{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tridentsnapshotinfos").
		Name(tridentSnapshotInfo.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentSnapshotInfo).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *tridentSnapshotInfos) UpdateStatus(ctx context.Context, tridentSnapshotInfo *v1.TridentSnapshotInfo, opts metav1.UpdateOptions) (result *v1.TridentSnapshotInfo, err error) {
	result = &v1.TridentSnapshotInfo{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tridentsnapshotinfos").
		Name(tridentSnapshotInfo.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentSnapshotInfo).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tridentSnapshotInfo and deletes it. Returns an error if one occurs.
func (c *tridentSnapshotInfos) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridentsnapshotinfos").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tridentSnapshotInfos) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridentsnapshotinfos").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tridentSnapshotInfo.
func (c *tridentSnapshotInfos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentSnapshotInfo, err error) {
	result = &v1.TridentSnapshotInfo{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tridentsnapshotinfos").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
