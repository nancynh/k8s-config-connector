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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networkservices/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkServicesEdgeCacheOrigins implements NetworkServicesEdgeCacheOriginInterface
type FakeNetworkServicesEdgeCacheOrigins struct {
	Fake *FakeNetworkservicesV1alpha1
	ns   string
}

var networkservicesedgecacheoriginsResource = v1alpha1.SchemeGroupVersion.WithResource("networkservicesedgecacheorigins")

var networkservicesedgecacheoriginsKind = v1alpha1.SchemeGroupVersion.WithKind("NetworkServicesEdgeCacheOrigin")

// Get takes name of the networkServicesEdgeCacheOrigin, and returns the corresponding networkServicesEdgeCacheOrigin object, and an error if there is any.
func (c *FakeNetworkServicesEdgeCacheOrigins) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkServicesEdgeCacheOrigin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkservicesedgecacheoriginsResource, c.ns, name), &v1alpha1.NetworkServicesEdgeCacheOrigin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkServicesEdgeCacheOrigin), err
}

// List takes label and field selectors, and returns the list of NetworkServicesEdgeCacheOrigins that match those selectors.
func (c *FakeNetworkServicesEdgeCacheOrigins) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkServicesEdgeCacheOriginList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkservicesedgecacheoriginsResource, networkservicesedgecacheoriginsKind, c.ns, opts), &v1alpha1.NetworkServicesEdgeCacheOriginList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkServicesEdgeCacheOriginList{ListMeta: obj.(*v1alpha1.NetworkServicesEdgeCacheOriginList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkServicesEdgeCacheOriginList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkServicesEdgeCacheOrigins.
func (c *FakeNetworkServicesEdgeCacheOrigins) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkservicesedgecacheoriginsResource, c.ns, opts))

}

// Create takes the representation of a networkServicesEdgeCacheOrigin and creates it.  Returns the server's representation of the networkServicesEdgeCacheOrigin, and an error, if there is any.
func (c *FakeNetworkServicesEdgeCacheOrigins) Create(ctx context.Context, networkServicesEdgeCacheOrigin *v1alpha1.NetworkServicesEdgeCacheOrigin, opts v1.CreateOptions) (result *v1alpha1.NetworkServicesEdgeCacheOrigin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkservicesedgecacheoriginsResource, c.ns, networkServicesEdgeCacheOrigin), &v1alpha1.NetworkServicesEdgeCacheOrigin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkServicesEdgeCacheOrigin), err
}

// Update takes the representation of a networkServicesEdgeCacheOrigin and updates it. Returns the server's representation of the networkServicesEdgeCacheOrigin, and an error, if there is any.
func (c *FakeNetworkServicesEdgeCacheOrigins) Update(ctx context.Context, networkServicesEdgeCacheOrigin *v1alpha1.NetworkServicesEdgeCacheOrigin, opts v1.UpdateOptions) (result *v1alpha1.NetworkServicesEdgeCacheOrigin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkservicesedgecacheoriginsResource, c.ns, networkServicesEdgeCacheOrigin), &v1alpha1.NetworkServicesEdgeCacheOrigin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkServicesEdgeCacheOrigin), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkServicesEdgeCacheOrigins) UpdateStatus(ctx context.Context, networkServicesEdgeCacheOrigin *v1alpha1.NetworkServicesEdgeCacheOrigin, opts v1.UpdateOptions) (*v1alpha1.NetworkServicesEdgeCacheOrigin, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkservicesedgecacheoriginsResource, "status", c.ns, networkServicesEdgeCacheOrigin), &v1alpha1.NetworkServicesEdgeCacheOrigin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkServicesEdgeCacheOrigin), err
}

// Delete takes name of the networkServicesEdgeCacheOrigin and deletes it. Returns an error if one occurs.
func (c *FakeNetworkServicesEdgeCacheOrigins) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkservicesedgecacheoriginsResource, c.ns, name, opts), &v1alpha1.NetworkServicesEdgeCacheOrigin{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkServicesEdgeCacheOrigins) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkservicesedgecacheoriginsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkServicesEdgeCacheOriginList{})
	return err
}

// Patch applies the patch and returns the patched networkServicesEdgeCacheOrigin.
func (c *FakeNetworkServicesEdgeCacheOrigins) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkServicesEdgeCacheOrigin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkservicesedgecacheoriginsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkServicesEdgeCacheOrigin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkServicesEdgeCacheOrigin), err
}
