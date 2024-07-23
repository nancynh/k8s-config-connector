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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networkservices/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkServicesHTTPRoutes implements NetworkServicesHTTPRouteInterface
type FakeNetworkServicesHTTPRoutes struct {
	Fake *FakeNetworkservicesV1beta1
	ns   string
}

var networkserviceshttproutesResource = v1beta1.SchemeGroupVersion.WithResource("networkserviceshttproutes")

var networkserviceshttproutesKind = v1beta1.SchemeGroupVersion.WithKind("NetworkServicesHTTPRoute")

// Get takes name of the networkServicesHTTPRoute, and returns the corresponding networkServicesHTTPRoute object, and an error if there is any.
func (c *FakeNetworkServicesHTTPRoutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.NetworkServicesHTTPRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkserviceshttproutesResource, c.ns, name), &v1beta1.NetworkServicesHTTPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesHTTPRoute), err
}

// List takes label and field selectors, and returns the list of NetworkServicesHTTPRoutes that match those selectors.
func (c *FakeNetworkServicesHTTPRoutes) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.NetworkServicesHTTPRouteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkserviceshttproutesResource, networkserviceshttproutesKind, c.ns, opts), &v1beta1.NetworkServicesHTTPRouteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.NetworkServicesHTTPRouteList{ListMeta: obj.(*v1beta1.NetworkServicesHTTPRouteList).ListMeta}
	for _, item := range obj.(*v1beta1.NetworkServicesHTTPRouteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkServicesHTTPRoutes.
func (c *FakeNetworkServicesHTTPRoutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkserviceshttproutesResource, c.ns, opts))

}

// Create takes the representation of a networkServicesHTTPRoute and creates it.  Returns the server's representation of the networkServicesHTTPRoute, and an error, if there is any.
func (c *FakeNetworkServicesHTTPRoutes) Create(ctx context.Context, networkServicesHTTPRoute *v1beta1.NetworkServicesHTTPRoute, opts v1.CreateOptions) (result *v1beta1.NetworkServicesHTTPRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkserviceshttproutesResource, c.ns, networkServicesHTTPRoute), &v1beta1.NetworkServicesHTTPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesHTTPRoute), err
}

// Update takes the representation of a networkServicesHTTPRoute and updates it. Returns the server's representation of the networkServicesHTTPRoute, and an error, if there is any.
func (c *FakeNetworkServicesHTTPRoutes) Update(ctx context.Context, networkServicesHTTPRoute *v1beta1.NetworkServicesHTTPRoute, opts v1.UpdateOptions) (result *v1beta1.NetworkServicesHTTPRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkserviceshttproutesResource, c.ns, networkServicesHTTPRoute), &v1beta1.NetworkServicesHTTPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesHTTPRoute), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkServicesHTTPRoutes) UpdateStatus(ctx context.Context, networkServicesHTTPRoute *v1beta1.NetworkServicesHTTPRoute, opts v1.UpdateOptions) (*v1beta1.NetworkServicesHTTPRoute, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkserviceshttproutesResource, "status", c.ns, networkServicesHTTPRoute), &v1beta1.NetworkServicesHTTPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesHTTPRoute), err
}

// Delete takes name of the networkServicesHTTPRoute and deletes it. Returns an error if one occurs.
func (c *FakeNetworkServicesHTTPRoutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkserviceshttproutesResource, c.ns, name, opts), &v1beta1.NetworkServicesHTTPRoute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkServicesHTTPRoutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkserviceshttproutesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.NetworkServicesHTTPRouteList{})
	return err
}

// Patch applies the patch and returns the patched networkServicesHTTPRoute.
func (c *FakeNetworkServicesHTTPRoutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NetworkServicesHTTPRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkserviceshttproutesResource, c.ns, name, pt, data, subresources...), &v1beta1.NetworkServicesHTTPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesHTTPRoute), err
}
