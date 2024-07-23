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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/datacatalog/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDataCatalogTags implements DataCatalogTagInterface
type FakeDataCatalogTags struct {
	Fake *FakeDatacatalogV1alpha1
	ns   string
}

var datacatalogtagsResource = v1alpha1.SchemeGroupVersion.WithResource("datacatalogtags")

var datacatalogtagsKind = v1alpha1.SchemeGroupVersion.WithKind("DataCatalogTag")

// Get takes name of the dataCatalogTag, and returns the corresponding dataCatalogTag object, and an error if there is any.
func (c *FakeDataCatalogTags) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DataCatalogTag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datacatalogtagsResource, c.ns, name), &v1alpha1.DataCatalogTag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataCatalogTag), err
}

// List takes label and field selectors, and returns the list of DataCatalogTags that match those selectors.
func (c *FakeDataCatalogTags) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DataCatalogTagList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datacatalogtagsResource, datacatalogtagsKind, c.ns, opts), &v1alpha1.DataCatalogTagList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataCatalogTagList{ListMeta: obj.(*v1alpha1.DataCatalogTagList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataCatalogTagList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataCatalogTags.
func (c *FakeDataCatalogTags) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datacatalogtagsResource, c.ns, opts))

}

// Create takes the representation of a dataCatalogTag and creates it.  Returns the server's representation of the dataCatalogTag, and an error, if there is any.
func (c *FakeDataCatalogTags) Create(ctx context.Context, dataCatalogTag *v1alpha1.DataCatalogTag, opts v1.CreateOptions) (result *v1alpha1.DataCatalogTag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datacatalogtagsResource, c.ns, dataCatalogTag), &v1alpha1.DataCatalogTag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataCatalogTag), err
}

// Update takes the representation of a dataCatalogTag and updates it. Returns the server's representation of the dataCatalogTag, and an error, if there is any.
func (c *FakeDataCatalogTags) Update(ctx context.Context, dataCatalogTag *v1alpha1.DataCatalogTag, opts v1.UpdateOptions) (result *v1alpha1.DataCatalogTag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datacatalogtagsResource, c.ns, dataCatalogTag), &v1alpha1.DataCatalogTag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataCatalogTag), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataCatalogTags) UpdateStatus(ctx context.Context, dataCatalogTag *v1alpha1.DataCatalogTag, opts v1.UpdateOptions) (*v1alpha1.DataCatalogTag, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datacatalogtagsResource, "status", c.ns, dataCatalogTag), &v1alpha1.DataCatalogTag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataCatalogTag), err
}

// Delete takes name of the dataCatalogTag and deletes it. Returns an error if one occurs.
func (c *FakeDataCatalogTags) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(datacatalogtagsResource, c.ns, name, opts), &v1alpha1.DataCatalogTag{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataCatalogTags) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datacatalogtagsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataCatalogTagList{})
	return err
}

// Patch applies the patch and returns the patched dataCatalogTag.
func (c *FakeDataCatalogTags) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DataCatalogTag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datacatalogtagsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataCatalogTag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataCatalogTag), err
}
