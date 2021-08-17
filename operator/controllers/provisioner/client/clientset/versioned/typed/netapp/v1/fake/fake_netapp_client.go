// Copyright 2020 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/netapp/trident/v21/operator/controllers/provisioner/client/clientset/versioned/typed/netapp/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeTridentV1 struct {
	*testing.Fake
}

func (c *FakeTridentV1) TridentProvisioners(namespace string) v1.TridentProvisionerInterface {
	return &FakeTridentProvisioners{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeTridentV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
