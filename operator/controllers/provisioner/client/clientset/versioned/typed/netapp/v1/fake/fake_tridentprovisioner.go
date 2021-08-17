// Copyright 2020 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	netappv1 "github.com/netapp/trident/v21/operator/controllers/provisioner/apis/netapp/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTridentProvisioners implements TridentProvisionerInterface
type FakeTridentProvisioners struct {
	Fake *FakeTridentV1
	ns   string
}

var tridentprovisionersResource = schema.GroupVersionResource{Group: "trident.netapp.io", Version: "v1", Resource: "tridentprovisioners"}

var tridentprovisionersKind = schema.GroupVersionKind{Group: "trident.netapp.io", Version: "v1", Kind: "TridentProvisioner"}

// Get takes name of the tridentProvisioner, and returns the corresponding tridentProvisioner object, and an error if there is any.
func (c *FakeTridentProvisioners) Get(ctx context.Context, name string, options v1.GetOptions) (result *netappv1.TridentProvisioner, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tridentprovisionersResource, c.ns, name), &netappv1.TridentProvisioner{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentProvisioner), err
}

// List takes label and field selectors, and returns the list of TridentProvisioners that match those selectors.
func (c *FakeTridentProvisioners) List(ctx context.Context, opts v1.ListOptions) (result *netappv1.TridentProvisionerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tridentprovisionersResource, tridentprovisionersKind, c.ns, opts), &netappv1.TridentProvisionerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &netappv1.TridentProvisionerList{ListMeta: obj.(*netappv1.TridentProvisionerList).ListMeta}
	for _, item := range obj.(*netappv1.TridentProvisionerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tridentProvisioners.
func (c *FakeTridentProvisioners) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tridentprovisionersResource, c.ns, opts))

}

// Create takes the representation of a tridentProvisioner and creates it.  Returns the server's representation of the tridentProvisioner, and an error, if there is any.
func (c *FakeTridentProvisioners) Create(ctx context.Context, tridentProvisioner *netappv1.TridentProvisioner, opts v1.CreateOptions) (result *netappv1.TridentProvisioner, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tridentprovisionersResource, c.ns, tridentProvisioner), &netappv1.TridentProvisioner{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentProvisioner), err
}

// Update takes the representation of a tridentProvisioner and updates it. Returns the server's representation of the tridentProvisioner, and an error, if there is any.
func (c *FakeTridentProvisioners) Update(ctx context.Context, tridentProvisioner *netappv1.TridentProvisioner, opts v1.UpdateOptions) (result *netappv1.TridentProvisioner, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tridentprovisionersResource, c.ns, tridentProvisioner), &netappv1.TridentProvisioner{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentProvisioner), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTridentProvisioners) UpdateStatus(ctx context.Context, tridentProvisioner *netappv1.TridentProvisioner, opts v1.UpdateOptions) (*netappv1.TridentProvisioner, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tridentprovisionersResource, "status", c.ns, tridentProvisioner), &netappv1.TridentProvisioner{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentProvisioner), err
}

// Delete takes name of the tridentProvisioner and deletes it. Returns an error if one occurs.
func (c *FakeTridentProvisioners) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tridentprovisionersResource, c.ns, name), &netappv1.TridentProvisioner{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTridentProvisioners) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tridentprovisionersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &netappv1.TridentProvisionerList{})
	return err
}

// Patch applies the patch and returns the patched tridentProvisioner.
func (c *FakeTridentProvisioners) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *netappv1.TridentProvisioner, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tridentprovisionersResource, c.ns, name, pt, data, subresources...), &netappv1.TridentProvisioner{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentProvisioner), err
}
