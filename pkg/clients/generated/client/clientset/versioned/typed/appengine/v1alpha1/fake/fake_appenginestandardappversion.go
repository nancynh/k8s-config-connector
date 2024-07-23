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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/appengine/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAppEngineStandardAppVersions implements AppEngineStandardAppVersionInterface
type FakeAppEngineStandardAppVersions struct {
	Fake *FakeAppengineV1alpha1
	ns   string
}

var appenginestandardappversionsResource = v1alpha1.SchemeGroupVersion.WithResource("appenginestandardappversions")

var appenginestandardappversionsKind = v1alpha1.SchemeGroupVersion.WithKind("AppEngineStandardAppVersion")

// Get takes name of the appEngineStandardAppVersion, and returns the corresponding appEngineStandardAppVersion object, and an error if there is any.
func (c *FakeAppEngineStandardAppVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AppEngineStandardAppVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appenginestandardappversionsResource, c.ns, name), &v1alpha1.AppEngineStandardAppVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppEngineStandardAppVersion), err
}

// List takes label and field selectors, and returns the list of AppEngineStandardAppVersions that match those selectors.
func (c *FakeAppEngineStandardAppVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AppEngineStandardAppVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appenginestandardappversionsResource, appenginestandardappversionsKind, c.ns, opts), &v1alpha1.AppEngineStandardAppVersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppEngineStandardAppVersionList{ListMeta: obj.(*v1alpha1.AppEngineStandardAppVersionList).ListMeta}
	for _, item := range obj.(*v1alpha1.AppEngineStandardAppVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appEngineStandardAppVersions.
func (c *FakeAppEngineStandardAppVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appenginestandardappversionsResource, c.ns, opts))

}

// Create takes the representation of a appEngineStandardAppVersion and creates it.  Returns the server's representation of the appEngineStandardAppVersion, and an error, if there is any.
func (c *FakeAppEngineStandardAppVersions) Create(ctx context.Context, appEngineStandardAppVersion *v1alpha1.AppEngineStandardAppVersion, opts v1.CreateOptions) (result *v1alpha1.AppEngineStandardAppVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appenginestandardappversionsResource, c.ns, appEngineStandardAppVersion), &v1alpha1.AppEngineStandardAppVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppEngineStandardAppVersion), err
}

// Update takes the representation of a appEngineStandardAppVersion and updates it. Returns the server's representation of the appEngineStandardAppVersion, and an error, if there is any.
func (c *FakeAppEngineStandardAppVersions) Update(ctx context.Context, appEngineStandardAppVersion *v1alpha1.AppEngineStandardAppVersion, opts v1.UpdateOptions) (result *v1alpha1.AppEngineStandardAppVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appenginestandardappversionsResource, c.ns, appEngineStandardAppVersion), &v1alpha1.AppEngineStandardAppVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppEngineStandardAppVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppEngineStandardAppVersions) UpdateStatus(ctx context.Context, appEngineStandardAppVersion *v1alpha1.AppEngineStandardAppVersion, opts v1.UpdateOptions) (*v1alpha1.AppEngineStandardAppVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(appenginestandardappversionsResource, "status", c.ns, appEngineStandardAppVersion), &v1alpha1.AppEngineStandardAppVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppEngineStandardAppVersion), err
}

// Delete takes name of the appEngineStandardAppVersion and deletes it. Returns an error if one occurs.
func (c *FakeAppEngineStandardAppVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(appenginestandardappversionsResource, c.ns, name, opts), &v1alpha1.AppEngineStandardAppVersion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppEngineStandardAppVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appenginestandardappversionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppEngineStandardAppVersionList{})
	return err
}

// Patch applies the patch and returns the patched appEngineStandardAppVersion.
func (c *FakeAppEngineStandardAppVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AppEngineStandardAppVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appenginestandardappversionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AppEngineStandardAppVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppEngineStandardAppVersion), err
}
