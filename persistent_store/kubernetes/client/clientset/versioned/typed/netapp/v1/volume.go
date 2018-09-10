// Copyright 2018 NetApp, Inc. All Rights Reserved.
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/netapp/trident/persistent_store/kubernetes/apis/netapp/v1"
	scheme "github.com/netapp/trident/persistent_store/kubernetes/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VolumesGetter has a method to return a VolumeInterface.
// A group's client should implement this interface.
type VolumesGetter interface {
	Volumes() VolumeInterface
}

// VolumeInterface has methods to work with Volume resources.
type VolumeInterface interface {
	Create(*v1.Volume) (*v1.Volume, error)
	Update(*v1.Volume) (*v1.Volume, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Volume, error)
	List(opts metav1.ListOptions) (*v1.VolumeList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Volume, err error)
	VolumeExpansion
}

// volumes implements VolumeInterface
type volumes struct {
	client rest.Interface
}

// newVolumes returns a Volumes
func newVolumes(c *TridentV1Client) *volumes {
	return &volumes{
		client: c.RESTClient(),
	}
}

// Get takes name of the volume, and returns the corresponding volume object, and an error if there is any.
func (c *volumes) Get(name string, options metav1.GetOptions) (result *v1.Volume, err error) {
	result = &v1.Volume{}
	err = c.client.Get().
		Resource("volumes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Volumes that match those selectors.
func (c *volumes) List(opts metav1.ListOptions) (result *v1.VolumeList, err error) {
	result = &v1.VolumeList{}
	err = c.client.Get().
		Resource("volumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumes.
func (c *volumes) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("volumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a volume and creates it.  Returns the server's representation of the volume, and an error, if there is any.
func (c *volumes) Create(volume *v1.Volume) (result *v1.Volume, err error) {
	result = &v1.Volume{}
	err = c.client.Post().
		Resource("volumes").
		Body(volume).
		Do().
		Into(result)
	return
}

// Update takes the representation of a volume and updates it. Returns the server's representation of the volume, and an error, if there is any.
func (c *volumes) Update(volume *v1.Volume) (result *v1.Volume, err error) {
	result = &v1.Volume{}
	err = c.client.Put().
		Resource("volumes").
		Name(volume.Name).
		Body(volume).
		Do().
		Into(result)
	return
}

// Delete takes name of the volume and deletes it. Returns an error if one occurs.
func (c *volumes) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("volumes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumes) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Resource("volumes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched volume.
func (c *volumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Volume, err error) {
	result = &v1.Volume{}
	err = c.client.Patch(pt).
		Resource("volumes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
