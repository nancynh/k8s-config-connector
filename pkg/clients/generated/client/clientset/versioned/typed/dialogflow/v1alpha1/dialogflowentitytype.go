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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dialogflow/v1alpha1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DialogflowEntityTypesGetter has a method to return a DialogflowEntityTypeInterface.
// A group's client should implement this interface.
type DialogflowEntityTypesGetter interface {
	DialogflowEntityTypes(namespace string) DialogflowEntityTypeInterface
}

// DialogflowEntityTypeInterface has methods to work with DialogflowEntityType resources.
type DialogflowEntityTypeInterface interface {
	Create(ctx context.Context, dialogflowEntityType *v1alpha1.DialogflowEntityType, opts v1.CreateOptions) (*v1alpha1.DialogflowEntityType, error)
	Update(ctx context.Context, dialogflowEntityType *v1alpha1.DialogflowEntityType, opts v1.UpdateOptions) (*v1alpha1.DialogflowEntityType, error)
	UpdateStatus(ctx context.Context, dialogflowEntityType *v1alpha1.DialogflowEntityType, opts v1.UpdateOptions) (*v1alpha1.DialogflowEntityType, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DialogflowEntityType, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DialogflowEntityTypeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DialogflowEntityType, err error)
	DialogflowEntityTypeExpansion
}

// dialogflowEntityTypes implements DialogflowEntityTypeInterface
type dialogflowEntityTypes struct {
	client rest.Interface
	ns     string
}

// newDialogflowEntityTypes returns a DialogflowEntityTypes
func newDialogflowEntityTypes(c *DialogflowV1alpha1Client, namespace string) *dialogflowEntityTypes {
	return &dialogflowEntityTypes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dialogflowEntityType, and returns the corresponding dialogflowEntityType object, and an error if there is any.
func (c *dialogflowEntityTypes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DialogflowEntityType, err error) {
	result = &v1alpha1.DialogflowEntityType{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dialogflowentitytypes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DialogflowEntityTypes that match those selectors.
func (c *dialogflowEntityTypes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DialogflowEntityTypeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DialogflowEntityTypeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dialogflowentitytypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dialogflowEntityTypes.
func (c *dialogflowEntityTypes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dialogflowentitytypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dialogflowEntityType and creates it.  Returns the server's representation of the dialogflowEntityType, and an error, if there is any.
func (c *dialogflowEntityTypes) Create(ctx context.Context, dialogflowEntityType *v1alpha1.DialogflowEntityType, opts v1.CreateOptions) (result *v1alpha1.DialogflowEntityType, err error) {
	result = &v1alpha1.DialogflowEntityType{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dialogflowentitytypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dialogflowEntityType).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dialogflowEntityType and updates it. Returns the server's representation of the dialogflowEntityType, and an error, if there is any.
func (c *dialogflowEntityTypes) Update(ctx context.Context, dialogflowEntityType *v1alpha1.DialogflowEntityType, opts v1.UpdateOptions) (result *v1alpha1.DialogflowEntityType, err error) {
	result = &v1alpha1.DialogflowEntityType{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dialogflowentitytypes").
		Name(dialogflowEntityType.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dialogflowEntityType).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *dialogflowEntityTypes) UpdateStatus(ctx context.Context, dialogflowEntityType *v1alpha1.DialogflowEntityType, opts v1.UpdateOptions) (result *v1alpha1.DialogflowEntityType, err error) {
	result = &v1alpha1.DialogflowEntityType{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dialogflowentitytypes").
		Name(dialogflowEntityType.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dialogflowEntityType).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dialogflowEntityType and deletes it. Returns an error if one occurs.
func (c *dialogflowEntityTypes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dialogflowentitytypes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dialogflowEntityTypes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dialogflowentitytypes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dialogflowEntityType.
func (c *dialogflowEntityTypes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DialogflowEntityType, err error) {
	result = &v1alpha1.DialogflowEntityType{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dialogflowentitytypes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
