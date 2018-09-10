// Copyright 2018 NetApp, Inc. All Rights Reserved.
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	netappv1 "github.com/netapp/trident/persistent_store/kubernetes/apis/netapp/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVolumes implements VolumeInterface
type FakeVolumes struct {
	Fake *FakeTridentV1
}

var volumesResource = schema.GroupVersionResource{Group: "trident.netapp.io", Version: "v1", Resource: "volumes"}

var volumesKind = schema.GroupVersionKind{Group: "trident.netapp.io", Version: "v1", Kind: "Volume"}

// Get takes name of the volume, and returns the corresponding volume object, and an error if there is any.
func (c *FakeVolumes) Get(name string, options v1.GetOptions) (result *netappv1.Volume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(volumesResource, name), &netappv1.Volume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.Volume), err
}

// List takes label and field selectors, and returns the list of Volumes that match those selectors.
func (c *FakeVolumes) List(opts v1.ListOptions) (result *netappv1.VolumeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(volumesResource, volumesKind, opts), &netappv1.VolumeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &netappv1.VolumeList{}
	for _, item := range obj.(*netappv1.VolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumes.
func (c *FakeVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(volumesResource, opts))
}

// Create takes the representation of a volume and creates it.  Returns the server's representation of the volume, and an error, if there is any.
func (c *FakeVolumes) Create(volume *netappv1.Volume) (result *netappv1.Volume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(volumesResource, volume), &netappv1.Volume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.Volume), err
}

// Update takes the representation of a volume and updates it. Returns the server's representation of the volume, and an error, if there is any.
func (c *FakeVolumes) Update(volume *netappv1.Volume) (result *netappv1.Volume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(volumesResource, volume), &netappv1.Volume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.Volume), err
}

// Delete takes name of the volume and deletes it. Returns an error if one occurs.
func (c *FakeVolumes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(volumesResource, name), &netappv1.Volume{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(volumesResource, listOptions)

	_, err := c.Fake.Invokes(action, &netappv1.VolumeList{})
	return err
}

// Patch applies the patch and returns the patched volume.
func (c *FakeVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *netappv1.Volume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(volumesResource, name, data, subresources...), &netappv1.Volume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.Volume), err
}
