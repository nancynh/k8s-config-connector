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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/privateca/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePrivateCACAPools implements PrivateCACAPoolInterface
type FakePrivateCACAPools struct {
	Fake *FakePrivatecaV1beta1
	ns   string
}

var privatecacapoolsResource = v1beta1.SchemeGroupVersion.WithResource("privatecacapools")

var privatecacapoolsKind = v1beta1.SchemeGroupVersion.WithKind("PrivateCACAPool")

// Get takes name of the privateCACAPool, and returns the corresponding privateCACAPool object, and an error if there is any.
func (c *FakePrivateCACAPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.PrivateCACAPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(privatecacapoolsResource, c.ns, name), &v1beta1.PrivateCACAPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PrivateCACAPool), err
}

// List takes label and field selectors, and returns the list of PrivateCACAPools that match those selectors.
func (c *FakePrivateCACAPools) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PrivateCACAPoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(privatecacapoolsResource, privatecacapoolsKind, c.ns, opts), &v1beta1.PrivateCACAPoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PrivateCACAPoolList{ListMeta: obj.(*v1beta1.PrivateCACAPoolList).ListMeta}
	for _, item := range obj.(*v1beta1.PrivateCACAPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested privateCACAPools.
func (c *FakePrivateCACAPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(privatecacapoolsResource, c.ns, opts))

}

// Create takes the representation of a privateCACAPool and creates it.  Returns the server's representation of the privateCACAPool, and an error, if there is any.
func (c *FakePrivateCACAPools) Create(ctx context.Context, privateCACAPool *v1beta1.PrivateCACAPool, opts v1.CreateOptions) (result *v1beta1.PrivateCACAPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(privatecacapoolsResource, c.ns, privateCACAPool), &v1beta1.PrivateCACAPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PrivateCACAPool), err
}

// Update takes the representation of a privateCACAPool and updates it. Returns the server's representation of the privateCACAPool, and an error, if there is any.
func (c *FakePrivateCACAPools) Update(ctx context.Context, privateCACAPool *v1beta1.PrivateCACAPool, opts v1.UpdateOptions) (result *v1beta1.PrivateCACAPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(privatecacapoolsResource, c.ns, privateCACAPool), &v1beta1.PrivateCACAPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PrivateCACAPool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrivateCACAPools) UpdateStatus(ctx context.Context, privateCACAPool *v1beta1.PrivateCACAPool, opts v1.UpdateOptions) (*v1beta1.PrivateCACAPool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(privatecacapoolsResource, "status", c.ns, privateCACAPool), &v1beta1.PrivateCACAPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PrivateCACAPool), err
}

// Delete takes name of the privateCACAPool and deletes it. Returns an error if one occurs.
func (c *FakePrivateCACAPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(privatecacapoolsResource, c.ns, name, opts), &v1beta1.PrivateCACAPool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrivateCACAPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(privatecacapoolsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.PrivateCACAPoolList{})
	return err
}

// Patch applies the patch and returns the patched privateCACAPool.
func (c *FakePrivateCACAPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.PrivateCACAPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(privatecacapoolsResource, c.ns, name, pt, data, subresources...), &v1beta1.PrivateCACAPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PrivateCACAPool), err
}
