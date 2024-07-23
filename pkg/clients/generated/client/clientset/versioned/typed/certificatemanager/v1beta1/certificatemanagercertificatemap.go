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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/certificatemanager/v1beta1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CertificateManagerCertificateMapsGetter has a method to return a CertificateManagerCertificateMapInterface.
// A group's client should implement this interface.
type CertificateManagerCertificateMapsGetter interface {
	CertificateManagerCertificateMaps(namespace string) CertificateManagerCertificateMapInterface
}

// CertificateManagerCertificateMapInterface has methods to work with CertificateManagerCertificateMap resources.
type CertificateManagerCertificateMapInterface interface {
	Create(ctx context.Context, certificateManagerCertificateMap *v1beta1.CertificateManagerCertificateMap, opts v1.CreateOptions) (*v1beta1.CertificateManagerCertificateMap, error)
	Update(ctx context.Context, certificateManagerCertificateMap *v1beta1.CertificateManagerCertificateMap, opts v1.UpdateOptions) (*v1beta1.CertificateManagerCertificateMap, error)
	UpdateStatus(ctx context.Context, certificateManagerCertificateMap *v1beta1.CertificateManagerCertificateMap, opts v1.UpdateOptions) (*v1beta1.CertificateManagerCertificateMap, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.CertificateManagerCertificateMap, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.CertificateManagerCertificateMapList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CertificateManagerCertificateMap, err error)
	CertificateManagerCertificateMapExpansion
}

// certificateManagerCertificateMaps implements CertificateManagerCertificateMapInterface
type certificateManagerCertificateMaps struct {
	client rest.Interface
	ns     string
}

// newCertificateManagerCertificateMaps returns a CertificateManagerCertificateMaps
func newCertificateManagerCertificateMaps(c *CertificatemanagerV1beta1Client, namespace string) *certificateManagerCertificateMaps {
	return &certificateManagerCertificateMaps{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the certificateManagerCertificateMap, and returns the corresponding certificateManagerCertificateMap object, and an error if there is any.
func (c *certificateManagerCertificateMaps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.CertificateManagerCertificateMap, err error) {
	result = &v1beta1.CertificateManagerCertificateMap{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("certificatemanagercertificatemaps").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CertificateManagerCertificateMaps that match those selectors.
func (c *certificateManagerCertificateMaps) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CertificateManagerCertificateMapList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.CertificateManagerCertificateMapList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("certificatemanagercertificatemaps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested certificateManagerCertificateMaps.
func (c *certificateManagerCertificateMaps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("certificatemanagercertificatemaps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a certificateManagerCertificateMap and creates it.  Returns the server's representation of the certificateManagerCertificateMap, and an error, if there is any.
func (c *certificateManagerCertificateMaps) Create(ctx context.Context, certificateManagerCertificateMap *v1beta1.CertificateManagerCertificateMap, opts v1.CreateOptions) (result *v1beta1.CertificateManagerCertificateMap, err error) {
	result = &v1beta1.CertificateManagerCertificateMap{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("certificatemanagercertificatemaps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateManagerCertificateMap).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a certificateManagerCertificateMap and updates it. Returns the server's representation of the certificateManagerCertificateMap, and an error, if there is any.
func (c *certificateManagerCertificateMaps) Update(ctx context.Context, certificateManagerCertificateMap *v1beta1.CertificateManagerCertificateMap, opts v1.UpdateOptions) (result *v1beta1.CertificateManagerCertificateMap, err error) {
	result = &v1beta1.CertificateManagerCertificateMap{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("certificatemanagercertificatemaps").
		Name(certificateManagerCertificateMap.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateManagerCertificateMap).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *certificateManagerCertificateMaps) UpdateStatus(ctx context.Context, certificateManagerCertificateMap *v1beta1.CertificateManagerCertificateMap, opts v1.UpdateOptions) (result *v1beta1.CertificateManagerCertificateMap, err error) {
	result = &v1beta1.CertificateManagerCertificateMap{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("certificatemanagercertificatemaps").
		Name(certificateManagerCertificateMap.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateManagerCertificateMap).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the certificateManagerCertificateMap and deletes it. Returns an error if one occurs.
func (c *certificateManagerCertificateMaps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("certificatemanagercertificatemaps").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *certificateManagerCertificateMaps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("certificatemanagercertificatemaps").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched certificateManagerCertificateMap.
func (c *certificateManagerCertificateMaps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CertificateManagerCertificateMap, err error) {
	result = &v1beta1.CertificateManagerCertificateMap{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("certificatemanagercertificatemaps").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
