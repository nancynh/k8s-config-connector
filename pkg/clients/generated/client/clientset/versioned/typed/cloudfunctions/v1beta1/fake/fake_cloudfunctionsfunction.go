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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudfunctions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudFunctionsFunctions implements CloudFunctionsFunctionInterface
type FakeCloudFunctionsFunctions struct {
	Fake *FakeCloudfunctionsV1beta1
	ns   string
}

var cloudfunctionsfunctionsResource = v1beta1.SchemeGroupVersion.WithResource("cloudfunctionsfunctions")

var cloudfunctionsfunctionsKind = v1beta1.SchemeGroupVersion.WithKind("CloudFunctionsFunction")

// Get takes name of the cloudFunctionsFunction, and returns the corresponding cloudFunctionsFunction object, and an error if there is any.
func (c *FakeCloudFunctionsFunctions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.CloudFunctionsFunction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudfunctionsfunctionsResource, c.ns, name), &v1beta1.CloudFunctionsFunction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudFunctionsFunction), err
}

// List takes label and field selectors, and returns the list of CloudFunctionsFunctions that match those selectors.
func (c *FakeCloudFunctionsFunctions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CloudFunctionsFunctionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudfunctionsfunctionsResource, cloudfunctionsfunctionsKind, c.ns, opts), &v1beta1.CloudFunctionsFunctionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CloudFunctionsFunctionList{ListMeta: obj.(*v1beta1.CloudFunctionsFunctionList).ListMeta}
	for _, item := range obj.(*v1beta1.CloudFunctionsFunctionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudFunctionsFunctions.
func (c *FakeCloudFunctionsFunctions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudfunctionsfunctionsResource, c.ns, opts))

}

// Create takes the representation of a cloudFunctionsFunction and creates it.  Returns the server's representation of the cloudFunctionsFunction, and an error, if there is any.
func (c *FakeCloudFunctionsFunctions) Create(ctx context.Context, cloudFunctionsFunction *v1beta1.CloudFunctionsFunction, opts v1.CreateOptions) (result *v1beta1.CloudFunctionsFunction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudfunctionsfunctionsResource, c.ns, cloudFunctionsFunction), &v1beta1.CloudFunctionsFunction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudFunctionsFunction), err
}

// Update takes the representation of a cloudFunctionsFunction and updates it. Returns the server's representation of the cloudFunctionsFunction, and an error, if there is any.
func (c *FakeCloudFunctionsFunctions) Update(ctx context.Context, cloudFunctionsFunction *v1beta1.CloudFunctionsFunction, opts v1.UpdateOptions) (result *v1beta1.CloudFunctionsFunction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudfunctionsfunctionsResource, c.ns, cloudFunctionsFunction), &v1beta1.CloudFunctionsFunction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudFunctionsFunction), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudFunctionsFunctions) UpdateStatus(ctx context.Context, cloudFunctionsFunction *v1beta1.CloudFunctionsFunction, opts v1.UpdateOptions) (*v1beta1.CloudFunctionsFunction, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudfunctionsfunctionsResource, "status", c.ns, cloudFunctionsFunction), &v1beta1.CloudFunctionsFunction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudFunctionsFunction), err
}

// Delete takes name of the cloudFunctionsFunction and deletes it. Returns an error if one occurs.
func (c *FakeCloudFunctionsFunctions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cloudfunctionsfunctionsResource, c.ns, name, opts), &v1beta1.CloudFunctionsFunction{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudFunctionsFunctions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudfunctionsfunctionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.CloudFunctionsFunctionList{})
	return err
}

// Patch applies the patch and returns the patched cloudFunctionsFunction.
func (c *FakeCloudFunctionsFunctions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CloudFunctionsFunction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudfunctionsfunctionsResource, c.ns, name, pt, data, subresources...), &v1beta1.CloudFunctionsFunction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudFunctionsFunction), err
}
