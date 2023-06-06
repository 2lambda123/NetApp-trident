// Copyright 2023 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	netappv1 "github.com/netapp/trident/operator/controllers/orchestrator/apis/netapp/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTridentOrchestrators implements TridentOrchestratorInterface
type FakeTridentOrchestrators struct {
	Fake *FakeTridentV1
}

var tridentorchestratorsResource = schema.GroupVersionResource{Group: "trident.netapp.io", Version: "v1", Resource: "tridentorchestrators"}

var tridentorchestratorsKind = schema.GroupVersionKind{Group: "trident.netapp.io", Version: "v1", Kind: "TridentOrchestrator"}

// Get takes name of the tridentOrchestrator, and returns the corresponding tridentOrchestrator object, and an error if there is any.
func (c *FakeTridentOrchestrators) Get(ctx context.Context, name string, options v1.GetOptions) (result *netappv1.TridentOrchestrator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(tridentorchestratorsResource, name), &netappv1.TridentOrchestrator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentOrchestrator), err
}

// List takes label and field selectors, and returns the list of TridentOrchestrators that match those selectors.
func (c *FakeTridentOrchestrators) List(ctx context.Context, opts v1.ListOptions) (result *netappv1.TridentOrchestratorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(tridentorchestratorsResource, tridentorchestratorsKind, opts), &netappv1.TridentOrchestratorList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &netappv1.TridentOrchestratorList{ListMeta: obj.(*netappv1.TridentOrchestratorList).ListMeta}
	for _, item := range obj.(*netappv1.TridentOrchestratorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tridentOrchestrators.
func (c *FakeTridentOrchestrators) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(tridentorchestratorsResource, opts))
}

// Create takes the representation of a tridentOrchestrator and creates it.  Returns the server's representation of the tridentOrchestrator, and an error, if there is any.
func (c *FakeTridentOrchestrators) Create(ctx context.Context, tridentOrchestrator *netappv1.TridentOrchestrator, opts v1.CreateOptions) (result *netappv1.TridentOrchestrator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(tridentorchestratorsResource, tridentOrchestrator), &netappv1.TridentOrchestrator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentOrchestrator), err
}

// Update takes the representation of a tridentOrchestrator and updates it. Returns the server's representation of the tridentOrchestrator, and an error, if there is any.
func (c *FakeTridentOrchestrators) Update(ctx context.Context, tridentOrchestrator *netappv1.TridentOrchestrator, opts v1.UpdateOptions) (result *netappv1.TridentOrchestrator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(tridentorchestratorsResource, tridentOrchestrator), &netappv1.TridentOrchestrator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentOrchestrator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTridentOrchestrators) UpdateStatus(ctx context.Context, tridentOrchestrator *netappv1.TridentOrchestrator, opts v1.UpdateOptions) (*netappv1.TridentOrchestrator, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(tridentorchestratorsResource, "status", tridentOrchestrator), &netappv1.TridentOrchestrator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentOrchestrator), err
}

// Delete takes name of the tridentOrchestrator and deletes it. Returns an error if one occurs.
func (c *FakeTridentOrchestrators) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(tridentorchestratorsResource, name), &netappv1.TridentOrchestrator{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTridentOrchestrators) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(tridentorchestratorsResource, listOpts)

	_, err := c.Fake.Invokes(action, &netappv1.TridentOrchestratorList{})
	return err
}

// Patch applies the patch and returns the patched tridentOrchestrator.
func (c *FakeTridentOrchestrators) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *netappv1.TridentOrchestrator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(tridentorchestratorsResource, name, pt, data, subresources...), &netappv1.TridentOrchestrator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentOrchestrator), err
}
