// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/sdminonne/kubebuilt-crd-webhook-validated/pkg/client/clientset/versioned/typed/mygroup/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMygroupV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMygroupV1alpha1) Myresources(namespace string) v1alpha1.MyresourceInterface {
	return &FakeMyresources{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMygroupV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
