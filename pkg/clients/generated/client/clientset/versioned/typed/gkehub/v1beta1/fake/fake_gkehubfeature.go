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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/gkehub/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGKEHubFeatures implements GKEHubFeatureInterface
type FakeGKEHubFeatures struct {
	Fake *FakeGkehubV1beta1
	ns   string
}

var gkehubfeaturesResource = v1beta1.SchemeGroupVersion.WithResource("gkehubfeatures")

var gkehubfeaturesKind = v1beta1.SchemeGroupVersion.WithKind("GKEHubFeature")

// Get takes name of the gKEHubFeature, and returns the corresponding gKEHubFeature object, and an error if there is any.
func (c *FakeGKEHubFeatures) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.GKEHubFeature, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gkehubfeaturesResource, c.ns, name), &v1beta1.GKEHubFeature{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GKEHubFeature), err
}

// List takes label and field selectors, and returns the list of GKEHubFeatures that match those selectors.
func (c *FakeGKEHubFeatures) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.GKEHubFeatureList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gkehubfeaturesResource, gkehubfeaturesKind, c.ns, opts), &v1beta1.GKEHubFeatureList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.GKEHubFeatureList{ListMeta: obj.(*v1beta1.GKEHubFeatureList).ListMeta}
	for _, item := range obj.(*v1beta1.GKEHubFeatureList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gKEHubFeatures.
func (c *FakeGKEHubFeatures) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gkehubfeaturesResource, c.ns, opts))

}

// Create takes the representation of a gKEHubFeature and creates it.  Returns the server's representation of the gKEHubFeature, and an error, if there is any.
func (c *FakeGKEHubFeatures) Create(ctx context.Context, gKEHubFeature *v1beta1.GKEHubFeature, opts v1.CreateOptions) (result *v1beta1.GKEHubFeature, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gkehubfeaturesResource, c.ns, gKEHubFeature), &v1beta1.GKEHubFeature{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GKEHubFeature), err
}

// Update takes the representation of a gKEHubFeature and updates it. Returns the server's representation of the gKEHubFeature, and an error, if there is any.
func (c *FakeGKEHubFeatures) Update(ctx context.Context, gKEHubFeature *v1beta1.GKEHubFeature, opts v1.UpdateOptions) (result *v1beta1.GKEHubFeature, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gkehubfeaturesResource, c.ns, gKEHubFeature), &v1beta1.GKEHubFeature{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GKEHubFeature), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGKEHubFeatures) UpdateStatus(ctx context.Context, gKEHubFeature *v1beta1.GKEHubFeature, opts v1.UpdateOptions) (*v1beta1.GKEHubFeature, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gkehubfeaturesResource, "status", c.ns, gKEHubFeature), &v1beta1.GKEHubFeature{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GKEHubFeature), err
}

// Delete takes name of the gKEHubFeature and deletes it. Returns an error if one occurs.
func (c *FakeGKEHubFeatures) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(gkehubfeaturesResource, c.ns, name, opts), &v1beta1.GKEHubFeature{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGKEHubFeatures) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gkehubfeaturesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.GKEHubFeatureList{})
	return err
}

// Patch applies the patch and returns the patched gKEHubFeature.
func (c *FakeGKEHubFeatures) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.GKEHubFeature, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gkehubfeaturesResource, c.ns, name, pt, data, subresources...), &v1beta1.GKEHubFeature{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GKEHubFeature), err
}
