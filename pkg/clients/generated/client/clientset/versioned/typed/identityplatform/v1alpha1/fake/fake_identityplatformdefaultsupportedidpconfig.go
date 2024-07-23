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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/identityplatform/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIdentityPlatformDefaultSupportedIDPConfigs implements IdentityPlatformDefaultSupportedIDPConfigInterface
type FakeIdentityPlatformDefaultSupportedIDPConfigs struct {
	Fake *FakeIdentityplatformV1alpha1
	ns   string
}

var identityplatformdefaultsupportedidpconfigsResource = v1alpha1.SchemeGroupVersion.WithResource("identityplatformdefaultsupportedidpconfigs")

var identityplatformdefaultsupportedidpconfigsKind = v1alpha1.SchemeGroupVersion.WithKind("IdentityPlatformDefaultSupportedIDPConfig")

// Get takes name of the identityPlatformDefaultSupportedIDPConfig, and returns the corresponding identityPlatformDefaultSupportedIDPConfig object, and an error if there is any.
func (c *FakeIdentityPlatformDefaultSupportedIDPConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IdentityPlatformDefaultSupportedIDPConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(identityplatformdefaultsupportedidpconfigsResource, c.ns, name), &v1alpha1.IdentityPlatformDefaultSupportedIDPConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityPlatformDefaultSupportedIDPConfig), err
}

// List takes label and field selectors, and returns the list of IdentityPlatformDefaultSupportedIDPConfigs that match those selectors.
func (c *FakeIdentityPlatformDefaultSupportedIDPConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IdentityPlatformDefaultSupportedIDPConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(identityplatformdefaultsupportedidpconfigsResource, identityplatformdefaultsupportedidpconfigsKind, c.ns, opts), &v1alpha1.IdentityPlatformDefaultSupportedIDPConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IdentityPlatformDefaultSupportedIDPConfigList{ListMeta: obj.(*v1alpha1.IdentityPlatformDefaultSupportedIDPConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.IdentityPlatformDefaultSupportedIDPConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested identityPlatformDefaultSupportedIDPConfigs.
func (c *FakeIdentityPlatformDefaultSupportedIDPConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(identityplatformdefaultsupportedidpconfigsResource, c.ns, opts))

}

// Create takes the representation of a identityPlatformDefaultSupportedIDPConfig and creates it.  Returns the server's representation of the identityPlatformDefaultSupportedIDPConfig, and an error, if there is any.
func (c *FakeIdentityPlatformDefaultSupportedIDPConfigs) Create(ctx context.Context, identityPlatformDefaultSupportedIDPConfig *v1alpha1.IdentityPlatformDefaultSupportedIDPConfig, opts v1.CreateOptions) (result *v1alpha1.IdentityPlatformDefaultSupportedIDPConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(identityplatformdefaultsupportedidpconfigsResource, c.ns, identityPlatformDefaultSupportedIDPConfig), &v1alpha1.IdentityPlatformDefaultSupportedIDPConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityPlatformDefaultSupportedIDPConfig), err
}

// Update takes the representation of a identityPlatformDefaultSupportedIDPConfig and updates it. Returns the server's representation of the identityPlatformDefaultSupportedIDPConfig, and an error, if there is any.
func (c *FakeIdentityPlatformDefaultSupportedIDPConfigs) Update(ctx context.Context, identityPlatformDefaultSupportedIDPConfig *v1alpha1.IdentityPlatformDefaultSupportedIDPConfig, opts v1.UpdateOptions) (result *v1alpha1.IdentityPlatformDefaultSupportedIDPConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(identityplatformdefaultsupportedidpconfigsResource, c.ns, identityPlatformDefaultSupportedIDPConfig), &v1alpha1.IdentityPlatformDefaultSupportedIDPConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityPlatformDefaultSupportedIDPConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIdentityPlatformDefaultSupportedIDPConfigs) UpdateStatus(ctx context.Context, identityPlatformDefaultSupportedIDPConfig *v1alpha1.IdentityPlatformDefaultSupportedIDPConfig, opts v1.UpdateOptions) (*v1alpha1.IdentityPlatformDefaultSupportedIDPConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(identityplatformdefaultsupportedidpconfigsResource, "status", c.ns, identityPlatformDefaultSupportedIDPConfig), &v1alpha1.IdentityPlatformDefaultSupportedIDPConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityPlatformDefaultSupportedIDPConfig), err
}

// Delete takes name of the identityPlatformDefaultSupportedIDPConfig and deletes it. Returns an error if one occurs.
func (c *FakeIdentityPlatformDefaultSupportedIDPConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(identityplatformdefaultsupportedidpconfigsResource, c.ns, name, opts), &v1alpha1.IdentityPlatformDefaultSupportedIDPConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIdentityPlatformDefaultSupportedIDPConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(identityplatformdefaultsupportedidpconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IdentityPlatformDefaultSupportedIDPConfigList{})
	return err
}

// Patch applies the patch and returns the patched identityPlatformDefaultSupportedIDPConfig.
func (c *FakeIdentityPlatformDefaultSupportedIDPConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IdentityPlatformDefaultSupportedIDPConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(identityplatformdefaultsupportedidpconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IdentityPlatformDefaultSupportedIDPConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityPlatformDefaultSupportedIDPConfig), err
}
