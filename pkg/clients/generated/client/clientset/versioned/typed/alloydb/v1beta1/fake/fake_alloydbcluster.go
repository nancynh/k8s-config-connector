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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/alloydb/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAlloyDBClusters implements AlloyDBClusterInterface
type FakeAlloyDBClusters struct {
	Fake *FakeAlloydbV1beta1
	ns   string
}

var alloydbclustersResource = v1beta1.SchemeGroupVersion.WithResource("alloydbclusters")

var alloydbclustersKind = v1beta1.SchemeGroupVersion.WithKind("AlloyDBCluster")

// Get takes name of the alloyDBCluster, and returns the corresponding alloyDBCluster object, and an error if there is any.
func (c *FakeAlloyDBClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.AlloyDBCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(alloydbclustersResource, c.ns, name), &v1beta1.AlloyDBCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AlloyDBCluster), err
}

// List takes label and field selectors, and returns the list of AlloyDBClusters that match those selectors.
func (c *FakeAlloyDBClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.AlloyDBClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(alloydbclustersResource, alloydbclustersKind, c.ns, opts), &v1beta1.AlloyDBClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.AlloyDBClusterList{ListMeta: obj.(*v1beta1.AlloyDBClusterList).ListMeta}
	for _, item := range obj.(*v1beta1.AlloyDBClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alloyDBClusters.
func (c *FakeAlloyDBClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(alloydbclustersResource, c.ns, opts))

}

// Create takes the representation of a alloyDBCluster and creates it.  Returns the server's representation of the alloyDBCluster, and an error, if there is any.
func (c *FakeAlloyDBClusters) Create(ctx context.Context, alloyDBCluster *v1beta1.AlloyDBCluster, opts v1.CreateOptions) (result *v1beta1.AlloyDBCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(alloydbclustersResource, c.ns, alloyDBCluster), &v1beta1.AlloyDBCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AlloyDBCluster), err
}

// Update takes the representation of a alloyDBCluster and updates it. Returns the server's representation of the alloyDBCluster, and an error, if there is any.
func (c *FakeAlloyDBClusters) Update(ctx context.Context, alloyDBCluster *v1beta1.AlloyDBCluster, opts v1.UpdateOptions) (result *v1beta1.AlloyDBCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(alloydbclustersResource, c.ns, alloyDBCluster), &v1beta1.AlloyDBCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AlloyDBCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAlloyDBClusters) UpdateStatus(ctx context.Context, alloyDBCluster *v1beta1.AlloyDBCluster, opts v1.UpdateOptions) (*v1beta1.AlloyDBCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(alloydbclustersResource, "status", c.ns, alloyDBCluster), &v1beta1.AlloyDBCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AlloyDBCluster), err
}

// Delete takes name of the alloyDBCluster and deletes it. Returns an error if one occurs.
func (c *FakeAlloyDBClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(alloydbclustersResource, c.ns, name, opts), &v1beta1.AlloyDBCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlloyDBClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(alloydbclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.AlloyDBClusterList{})
	return err
}

// Patch applies the patch and returns the patched alloyDBCluster.
func (c *FakeAlloyDBClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.AlloyDBCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(alloydbclustersResource, c.ns, name, pt, data, subresources...), &v1beta1.AlloyDBCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AlloyDBCluster), err
}
