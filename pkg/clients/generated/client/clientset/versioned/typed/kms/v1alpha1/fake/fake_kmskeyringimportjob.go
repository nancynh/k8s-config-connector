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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/kms/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKMSKeyRingImportJobs implements KMSKeyRingImportJobInterface
type FakeKMSKeyRingImportJobs struct {
	Fake *FakeKmsV1alpha1
	ns   string
}

var kmskeyringimportjobsResource = v1alpha1.SchemeGroupVersion.WithResource("kmskeyringimportjobs")

var kmskeyringimportjobsKind = v1alpha1.SchemeGroupVersion.WithKind("KMSKeyRingImportJob")

// Get takes name of the kMSKeyRingImportJob, and returns the corresponding kMSKeyRingImportJob object, and an error if there is any.
func (c *FakeKMSKeyRingImportJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KMSKeyRingImportJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kmskeyringimportjobsResource, c.ns, name), &v1alpha1.KMSKeyRingImportJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KMSKeyRingImportJob), err
}

// List takes label and field selectors, and returns the list of KMSKeyRingImportJobs that match those selectors.
func (c *FakeKMSKeyRingImportJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KMSKeyRingImportJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kmskeyringimportjobsResource, kmskeyringimportjobsKind, c.ns, opts), &v1alpha1.KMSKeyRingImportJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KMSKeyRingImportJobList{ListMeta: obj.(*v1alpha1.KMSKeyRingImportJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.KMSKeyRingImportJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kMSKeyRingImportJobs.
func (c *FakeKMSKeyRingImportJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kmskeyringimportjobsResource, c.ns, opts))

}

// Create takes the representation of a kMSKeyRingImportJob and creates it.  Returns the server's representation of the kMSKeyRingImportJob, and an error, if there is any.
func (c *FakeKMSKeyRingImportJobs) Create(ctx context.Context, kMSKeyRingImportJob *v1alpha1.KMSKeyRingImportJob, opts v1.CreateOptions) (result *v1alpha1.KMSKeyRingImportJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kmskeyringimportjobsResource, c.ns, kMSKeyRingImportJob), &v1alpha1.KMSKeyRingImportJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KMSKeyRingImportJob), err
}

// Update takes the representation of a kMSKeyRingImportJob and updates it. Returns the server's representation of the kMSKeyRingImportJob, and an error, if there is any.
func (c *FakeKMSKeyRingImportJobs) Update(ctx context.Context, kMSKeyRingImportJob *v1alpha1.KMSKeyRingImportJob, opts v1.UpdateOptions) (result *v1alpha1.KMSKeyRingImportJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kmskeyringimportjobsResource, c.ns, kMSKeyRingImportJob), &v1alpha1.KMSKeyRingImportJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KMSKeyRingImportJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKMSKeyRingImportJobs) UpdateStatus(ctx context.Context, kMSKeyRingImportJob *v1alpha1.KMSKeyRingImportJob, opts v1.UpdateOptions) (*v1alpha1.KMSKeyRingImportJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kmskeyringimportjobsResource, "status", c.ns, kMSKeyRingImportJob), &v1alpha1.KMSKeyRingImportJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KMSKeyRingImportJob), err
}

// Delete takes name of the kMSKeyRingImportJob and deletes it. Returns an error if one occurs.
func (c *FakeKMSKeyRingImportJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(kmskeyringimportjobsResource, c.ns, name, opts), &v1alpha1.KMSKeyRingImportJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKMSKeyRingImportJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kmskeyringimportjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KMSKeyRingImportJobList{})
	return err
}

// Patch applies the patch and returns the patched kMSKeyRingImportJob.
func (c *FakeKMSKeyRingImportJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KMSKeyRingImportJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kmskeyringimportjobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KMSKeyRingImportJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KMSKeyRingImportJob), err
}
