// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeFirewalls implements ComputeFirewallInterface
type FakeComputeFirewalls struct {
	Fake *FakeComputeV1beta1
	ns   string
}

var computefirewallsResource = v1beta1.SchemeGroupVersion.WithResource("computefirewalls")

var computefirewallsKind = v1beta1.SchemeGroupVersion.WithKind("ComputeFirewall")

// Get takes name of the computeFirewall, and returns the corresponding computeFirewall object, and an error if there is any.
func (c *FakeComputeFirewalls) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeFirewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computefirewallsResource, c.ns, name), &v1beta1.ComputeFirewall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeFirewall), err
}

// List takes label and field selectors, and returns the list of ComputeFirewalls that match those selectors.
func (c *FakeComputeFirewalls) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeFirewallList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computefirewallsResource, computefirewallsKind, c.ns, opts), &v1beta1.ComputeFirewallList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ComputeFirewallList{ListMeta: obj.(*v1beta1.ComputeFirewallList).ListMeta}
	for _, item := range obj.(*v1beta1.ComputeFirewallList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeFirewalls.
func (c *FakeComputeFirewalls) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computefirewallsResource, c.ns, opts))

}

// Create takes the representation of a computeFirewall and creates it.  Returns the server's representation of the computeFirewall, and an error, if there is any.
func (c *FakeComputeFirewalls) Create(ctx context.Context, computeFirewall *v1beta1.ComputeFirewall, opts v1.CreateOptions) (result *v1beta1.ComputeFirewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computefirewallsResource, c.ns, computeFirewall), &v1beta1.ComputeFirewall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeFirewall), err
}

// Update takes the representation of a computeFirewall and updates it. Returns the server's representation of the computeFirewall, and an error, if there is any.
func (c *FakeComputeFirewalls) Update(ctx context.Context, computeFirewall *v1beta1.ComputeFirewall, opts v1.UpdateOptions) (result *v1beta1.ComputeFirewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computefirewallsResource, c.ns, computeFirewall), &v1beta1.ComputeFirewall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeFirewall), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeFirewalls) UpdateStatus(ctx context.Context, computeFirewall *v1beta1.ComputeFirewall, opts v1.UpdateOptions) (*v1beta1.ComputeFirewall, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computefirewallsResource, "status", c.ns, computeFirewall), &v1beta1.ComputeFirewall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeFirewall), err
}

// Delete takes name of the computeFirewall and deletes it. Returns an error if one occurs.
func (c *FakeComputeFirewalls) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computefirewallsResource, c.ns, name, opts), &v1beta1.ComputeFirewall{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeFirewalls) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computefirewallsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ComputeFirewallList{})
	return err
}

// Patch applies the patch and returns the patched computeFirewall.
func (c *FakeComputeFirewalls) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeFirewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computefirewallsResource, c.ns, name, pt, data, subresources...), &v1beta1.ComputeFirewall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeFirewall), err
}
