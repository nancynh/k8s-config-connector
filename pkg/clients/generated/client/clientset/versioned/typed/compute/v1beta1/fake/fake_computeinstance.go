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

// FakeComputeInstances implements ComputeInstanceInterface
type FakeComputeInstances struct {
	Fake *FakeComputeV1beta1
	ns   string
}

var computeinstancesResource = v1beta1.SchemeGroupVersion.WithResource("computeinstances")

var computeinstancesKind = v1beta1.SchemeGroupVersion.WithKind("ComputeInstance")

// Get takes name of the computeInstance, and returns the corresponding computeInstance object, and an error if there is any.
func (c *FakeComputeInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computeinstancesResource, c.ns, name), &v1beta1.ComputeInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeInstance), err
}

// List takes label and field selectors, and returns the list of ComputeInstances that match those selectors.
func (c *FakeComputeInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computeinstancesResource, computeinstancesKind, c.ns, opts), &v1beta1.ComputeInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ComputeInstanceList{ListMeta: obj.(*v1beta1.ComputeInstanceList).ListMeta}
	for _, item := range obj.(*v1beta1.ComputeInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeInstances.
func (c *FakeComputeInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computeinstancesResource, c.ns, opts))

}

// Create takes the representation of a computeInstance and creates it.  Returns the server's representation of the computeInstance, and an error, if there is any.
func (c *FakeComputeInstances) Create(ctx context.Context, computeInstance *v1beta1.ComputeInstance, opts v1.CreateOptions) (result *v1beta1.ComputeInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computeinstancesResource, c.ns, computeInstance), &v1beta1.ComputeInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeInstance), err
}

// Update takes the representation of a computeInstance and updates it. Returns the server's representation of the computeInstance, and an error, if there is any.
func (c *FakeComputeInstances) Update(ctx context.Context, computeInstance *v1beta1.ComputeInstance, opts v1.UpdateOptions) (result *v1beta1.ComputeInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computeinstancesResource, c.ns, computeInstance), &v1beta1.ComputeInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeInstances) UpdateStatus(ctx context.Context, computeInstance *v1beta1.ComputeInstance, opts v1.UpdateOptions) (*v1beta1.ComputeInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computeinstancesResource, "status", c.ns, computeInstance), &v1beta1.ComputeInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeInstance), err
}

// Delete takes name of the computeInstance and deletes it. Returns an error if one occurs.
func (c *FakeComputeInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computeinstancesResource, c.ns, name, opts), &v1beta1.ComputeInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computeinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ComputeInstanceList{})
	return err
}

// Patch applies the patch and returns the patched computeInstance.
func (c *FakeComputeInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computeinstancesResource, c.ns, name, pt, data, subresources...), &v1beta1.ComputeInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeInstance), err
}
