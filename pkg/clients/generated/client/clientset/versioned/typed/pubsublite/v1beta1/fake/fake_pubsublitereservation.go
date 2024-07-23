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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/pubsublite/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePubSubLiteReservations implements PubSubLiteReservationInterface
type FakePubSubLiteReservations struct {
	Fake *FakePubsubliteV1beta1
	ns   string
}

var pubsublitereservationsResource = v1beta1.SchemeGroupVersion.WithResource("pubsublitereservations")

var pubsublitereservationsKind = v1beta1.SchemeGroupVersion.WithKind("PubSubLiteReservation")

// Get takes name of the pubSubLiteReservation, and returns the corresponding pubSubLiteReservation object, and an error if there is any.
func (c *FakePubSubLiteReservations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.PubSubLiteReservation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pubsublitereservationsResource, c.ns, name), &v1beta1.PubSubLiteReservation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PubSubLiteReservation), err
}

// List takes label and field selectors, and returns the list of PubSubLiteReservations that match those selectors.
func (c *FakePubSubLiteReservations) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PubSubLiteReservationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pubsublitereservationsResource, pubsublitereservationsKind, c.ns, opts), &v1beta1.PubSubLiteReservationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PubSubLiteReservationList{ListMeta: obj.(*v1beta1.PubSubLiteReservationList).ListMeta}
	for _, item := range obj.(*v1beta1.PubSubLiteReservationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pubSubLiteReservations.
func (c *FakePubSubLiteReservations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pubsublitereservationsResource, c.ns, opts))

}

// Create takes the representation of a pubSubLiteReservation and creates it.  Returns the server's representation of the pubSubLiteReservation, and an error, if there is any.
func (c *FakePubSubLiteReservations) Create(ctx context.Context, pubSubLiteReservation *v1beta1.PubSubLiteReservation, opts v1.CreateOptions) (result *v1beta1.PubSubLiteReservation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pubsublitereservationsResource, c.ns, pubSubLiteReservation), &v1beta1.PubSubLiteReservation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PubSubLiteReservation), err
}

// Update takes the representation of a pubSubLiteReservation and updates it. Returns the server's representation of the pubSubLiteReservation, and an error, if there is any.
func (c *FakePubSubLiteReservations) Update(ctx context.Context, pubSubLiteReservation *v1beta1.PubSubLiteReservation, opts v1.UpdateOptions) (result *v1beta1.PubSubLiteReservation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pubsublitereservationsResource, c.ns, pubSubLiteReservation), &v1beta1.PubSubLiteReservation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PubSubLiteReservation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePubSubLiteReservations) UpdateStatus(ctx context.Context, pubSubLiteReservation *v1beta1.PubSubLiteReservation, opts v1.UpdateOptions) (*v1beta1.PubSubLiteReservation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pubsublitereservationsResource, "status", c.ns, pubSubLiteReservation), &v1beta1.PubSubLiteReservation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PubSubLiteReservation), err
}

// Delete takes name of the pubSubLiteReservation and deletes it. Returns an error if one occurs.
func (c *FakePubSubLiteReservations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(pubsublitereservationsResource, c.ns, name, opts), &v1beta1.PubSubLiteReservation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePubSubLiteReservations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pubsublitereservationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.PubSubLiteReservationList{})
	return err
}

// Patch applies the patch and returns the patched pubSubLiteReservation.
func (c *FakePubSubLiteReservations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.PubSubLiteReservation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pubsublitereservationsResource, c.ns, name, pt, data, subresources...), &v1beta1.PubSubLiteReservation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PubSubLiteReservation), err
}
