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
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/typed/healthcare/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeHealthcareV1alpha1 struct {
	*testing.Fake
}

func (c *FakeHealthcareV1alpha1) HealthcareConsentStores(namespace string) v1alpha1.HealthcareConsentStoreInterface {
	return &FakeHealthcareConsentStores{c, namespace}
}

func (c *FakeHealthcareV1alpha1) HealthcareDICOMStores(namespace string) v1alpha1.HealthcareDICOMStoreInterface {
	return &FakeHealthcareDICOMStores{c, namespace}
}

func (c *FakeHealthcareV1alpha1) HealthcareDatasets(namespace string) v1alpha1.HealthcareDatasetInterface {
	return &FakeHealthcareDatasets{c, namespace}
}

func (c *FakeHealthcareV1alpha1) HealthcareFHIRStores(namespace string) v1alpha1.HealthcareFHIRStoreInterface {
	return &FakeHealthcareFHIRStores{c, namespace}
}

func (c *FakeHealthcareV1alpha1) HealthcareHL7V2Stores(namespace string) v1alpha1.HealthcareHL7V2StoreInterface {
	return &FakeHealthcareHL7V2Stores{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeHealthcareV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
