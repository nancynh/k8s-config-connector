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

// Code generated by main. DO NOT EDIT.

package versioned

import (
	"fmt"

	accesscontextmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/accesscontextmanager/v1beta1"
	artifactregistryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/artifactregistry/v1beta1"
	bigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/bigquery/v1beta1"
	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/bigtable/v1beta1"
	cloudbuildv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/cloudbuild/v1beta1"
	cloudschedulerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/cloudscheduler/v1beta1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/compute/v1beta1"
	containerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/container/v1beta1"
	containeranalysisv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/containeranalysis/v1beta1"
	dataflowv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/dataflow/v1beta1"
	dataprocv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/dataproc/v1beta1"
	dnsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/dns/v1beta1"
	firestorev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/firestore/v1beta1"
	gameservicesv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/gameservices/v1beta1"
	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/iam/v1beta1"
	iapv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/iap/v1beta1"
	identityplatformv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/identityplatform/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/k8s/v1alpha1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/kms/v1beta1"
	loggingv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/logging/v1beta1"
	memcachev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/memcache/v1beta1"
	monitoringv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/monitoring/v1beta1"
	osconfigv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/osconfig/v1beta1"
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/pubsub/v1beta1"
	redisv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/redis/v1beta1"
	resourcemanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/resourcemanager/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/secretmanager/v1beta1"
	servicenetworkingv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/servicenetworking/v1beta1"
	serviceusagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/serviceusage/v1beta1"
	sourcerepov1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/sourcerepo/v1beta1"
	spannerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/spanner/v1beta1"
	sqlv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/sql/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/storage/v1beta1"
	storagetransferv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/storagetransfer/v1beta1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AccesscontextmanagerV1beta1() accesscontextmanagerv1beta1.AccesscontextmanagerV1beta1Interface
	ArtifactregistryV1beta1() artifactregistryv1beta1.ArtifactregistryV1beta1Interface
	BigqueryV1beta1() bigqueryv1beta1.BigqueryV1beta1Interface
	BigtableV1beta1() bigtablev1beta1.BigtableV1beta1Interface
	CloudbuildV1beta1() cloudbuildv1beta1.CloudbuildV1beta1Interface
	CloudschedulerV1beta1() cloudschedulerv1beta1.CloudschedulerV1beta1Interface
	ComputeV1beta1() computev1beta1.ComputeV1beta1Interface
	ContainerV1beta1() containerv1beta1.ContainerV1beta1Interface
	ContaineranalysisV1beta1() containeranalysisv1beta1.ContaineranalysisV1beta1Interface
	DataflowV1beta1() dataflowv1beta1.DataflowV1beta1Interface
	DataprocV1beta1() dataprocv1beta1.DataprocV1beta1Interface
	DnsV1beta1() dnsv1beta1.DnsV1beta1Interface
	FirestoreV1beta1() firestorev1beta1.FirestoreV1beta1Interface
	GameservicesV1beta1() gameservicesv1beta1.GameservicesV1beta1Interface
	IamV1beta1() iamv1beta1.IamV1beta1Interface
	IapV1beta1() iapv1beta1.IapV1beta1Interface
	IdentityplatformV1beta1() identityplatformv1beta1.IdentityplatformV1beta1Interface
	K8sV1alpha1() k8sv1alpha1.K8sV1alpha1Interface
	KmsV1beta1() kmsv1beta1.KmsV1beta1Interface
	LoggingV1beta1() loggingv1beta1.LoggingV1beta1Interface
	MemcacheV1beta1() memcachev1beta1.MemcacheV1beta1Interface
	MonitoringV1beta1() monitoringv1beta1.MonitoringV1beta1Interface
	OsconfigV1beta1() osconfigv1beta1.OsconfigV1beta1Interface
	PubsubV1beta1() pubsubv1beta1.PubsubV1beta1Interface
	RedisV1beta1() redisv1beta1.RedisV1beta1Interface
	ResourcemanagerV1beta1() resourcemanagerv1beta1.ResourcemanagerV1beta1Interface
	SecretmanagerV1beta1() secretmanagerv1beta1.SecretmanagerV1beta1Interface
	ServicenetworkingV1beta1() servicenetworkingv1beta1.ServicenetworkingV1beta1Interface
	ServiceusageV1beta1() serviceusagev1beta1.ServiceusageV1beta1Interface
	SourcerepoV1beta1() sourcerepov1beta1.SourcerepoV1beta1Interface
	SpannerV1beta1() spannerv1beta1.SpannerV1beta1Interface
	SqlV1beta1() sqlv1beta1.SqlV1beta1Interface
	StorageV1beta1() storagev1beta1.StorageV1beta1Interface
	StoragetransferV1beta1() storagetransferv1beta1.StoragetransferV1beta1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	accesscontextmanagerV1beta1 *accesscontextmanagerv1beta1.AccesscontextmanagerV1beta1Client
	artifactregistryV1beta1     *artifactregistryv1beta1.ArtifactregistryV1beta1Client
	bigqueryV1beta1             *bigqueryv1beta1.BigqueryV1beta1Client
	bigtableV1beta1             *bigtablev1beta1.BigtableV1beta1Client
	cloudbuildV1beta1           *cloudbuildv1beta1.CloudbuildV1beta1Client
	cloudschedulerV1beta1       *cloudschedulerv1beta1.CloudschedulerV1beta1Client
	computeV1beta1              *computev1beta1.ComputeV1beta1Client
	containerV1beta1            *containerv1beta1.ContainerV1beta1Client
	containeranalysisV1beta1    *containeranalysisv1beta1.ContaineranalysisV1beta1Client
	dataflowV1beta1             *dataflowv1beta1.DataflowV1beta1Client
	dataprocV1beta1             *dataprocv1beta1.DataprocV1beta1Client
	dnsV1beta1                  *dnsv1beta1.DnsV1beta1Client
	firestoreV1beta1            *firestorev1beta1.FirestoreV1beta1Client
	gameservicesV1beta1         *gameservicesv1beta1.GameservicesV1beta1Client
	iamV1beta1                  *iamv1beta1.IamV1beta1Client
	iapV1beta1                  *iapv1beta1.IapV1beta1Client
	identityplatformV1beta1     *identityplatformv1beta1.IdentityplatformV1beta1Client
	k8sV1alpha1                 *k8sv1alpha1.K8sV1alpha1Client
	kmsV1beta1                  *kmsv1beta1.KmsV1beta1Client
	loggingV1beta1              *loggingv1beta1.LoggingV1beta1Client
	memcacheV1beta1             *memcachev1beta1.MemcacheV1beta1Client
	monitoringV1beta1           *monitoringv1beta1.MonitoringV1beta1Client
	osconfigV1beta1             *osconfigv1beta1.OsconfigV1beta1Client
	pubsubV1beta1               *pubsubv1beta1.PubsubV1beta1Client
	redisV1beta1                *redisv1beta1.RedisV1beta1Client
	resourcemanagerV1beta1      *resourcemanagerv1beta1.ResourcemanagerV1beta1Client
	secretmanagerV1beta1        *secretmanagerv1beta1.SecretmanagerV1beta1Client
	servicenetworkingV1beta1    *servicenetworkingv1beta1.ServicenetworkingV1beta1Client
	serviceusageV1beta1         *serviceusagev1beta1.ServiceusageV1beta1Client
	sourcerepoV1beta1           *sourcerepov1beta1.SourcerepoV1beta1Client
	spannerV1beta1              *spannerv1beta1.SpannerV1beta1Client
	sqlV1beta1                  *sqlv1beta1.SqlV1beta1Client
	storageV1beta1              *storagev1beta1.StorageV1beta1Client
	storagetransferV1beta1      *storagetransferv1beta1.StoragetransferV1beta1Client
}

// AccesscontextmanagerV1beta1 retrieves the AccesscontextmanagerV1beta1Client
func (c *Clientset) AccesscontextmanagerV1beta1() accesscontextmanagerv1beta1.AccesscontextmanagerV1beta1Interface {
	return c.accesscontextmanagerV1beta1
}

// ArtifactregistryV1beta1 retrieves the ArtifactregistryV1beta1Client
func (c *Clientset) ArtifactregistryV1beta1() artifactregistryv1beta1.ArtifactregistryV1beta1Interface {
	return c.artifactregistryV1beta1
}

// BigqueryV1beta1 retrieves the BigqueryV1beta1Client
func (c *Clientset) BigqueryV1beta1() bigqueryv1beta1.BigqueryV1beta1Interface {
	return c.bigqueryV1beta1
}

// BigtableV1beta1 retrieves the BigtableV1beta1Client
func (c *Clientset) BigtableV1beta1() bigtablev1beta1.BigtableV1beta1Interface {
	return c.bigtableV1beta1
}

// CloudbuildV1beta1 retrieves the CloudbuildV1beta1Client
func (c *Clientset) CloudbuildV1beta1() cloudbuildv1beta1.CloudbuildV1beta1Interface {
	return c.cloudbuildV1beta1
}

// CloudschedulerV1beta1 retrieves the CloudschedulerV1beta1Client
func (c *Clientset) CloudschedulerV1beta1() cloudschedulerv1beta1.CloudschedulerV1beta1Interface {
	return c.cloudschedulerV1beta1
}

// ComputeV1beta1 retrieves the ComputeV1beta1Client
func (c *Clientset) ComputeV1beta1() computev1beta1.ComputeV1beta1Interface {
	return c.computeV1beta1
}

// ContainerV1beta1 retrieves the ContainerV1beta1Client
func (c *Clientset) ContainerV1beta1() containerv1beta1.ContainerV1beta1Interface {
	return c.containerV1beta1
}

// ContaineranalysisV1beta1 retrieves the ContaineranalysisV1beta1Client
func (c *Clientset) ContaineranalysisV1beta1() containeranalysisv1beta1.ContaineranalysisV1beta1Interface {
	return c.containeranalysisV1beta1
}

// DataflowV1beta1 retrieves the DataflowV1beta1Client
func (c *Clientset) DataflowV1beta1() dataflowv1beta1.DataflowV1beta1Interface {
	return c.dataflowV1beta1
}

// DataprocV1beta1 retrieves the DataprocV1beta1Client
func (c *Clientset) DataprocV1beta1() dataprocv1beta1.DataprocV1beta1Interface {
	return c.dataprocV1beta1
}

// DnsV1beta1 retrieves the DnsV1beta1Client
func (c *Clientset) DnsV1beta1() dnsv1beta1.DnsV1beta1Interface {
	return c.dnsV1beta1
}

// FirestoreV1beta1 retrieves the FirestoreV1beta1Client
func (c *Clientset) FirestoreV1beta1() firestorev1beta1.FirestoreV1beta1Interface {
	return c.firestoreV1beta1
}

// GameservicesV1beta1 retrieves the GameservicesV1beta1Client
func (c *Clientset) GameservicesV1beta1() gameservicesv1beta1.GameservicesV1beta1Interface {
	return c.gameservicesV1beta1
}

// IamV1beta1 retrieves the IamV1beta1Client
func (c *Clientset) IamV1beta1() iamv1beta1.IamV1beta1Interface {
	return c.iamV1beta1
}

// IapV1beta1 retrieves the IapV1beta1Client
func (c *Clientset) IapV1beta1() iapv1beta1.IapV1beta1Interface {
	return c.iapV1beta1
}

// IdentityplatformV1beta1 retrieves the IdentityplatformV1beta1Client
func (c *Clientset) IdentityplatformV1beta1() identityplatformv1beta1.IdentityplatformV1beta1Interface {
	return c.identityplatformV1beta1
}

// K8sV1alpha1 retrieves the K8sV1alpha1Client
func (c *Clientset) K8sV1alpha1() k8sv1alpha1.K8sV1alpha1Interface {
	return c.k8sV1alpha1
}

// KmsV1beta1 retrieves the KmsV1beta1Client
func (c *Clientset) KmsV1beta1() kmsv1beta1.KmsV1beta1Interface {
	return c.kmsV1beta1
}

// LoggingV1beta1 retrieves the LoggingV1beta1Client
func (c *Clientset) LoggingV1beta1() loggingv1beta1.LoggingV1beta1Interface {
	return c.loggingV1beta1
}

// MemcacheV1beta1 retrieves the MemcacheV1beta1Client
func (c *Clientset) MemcacheV1beta1() memcachev1beta1.MemcacheV1beta1Interface {
	return c.memcacheV1beta1
}

// MonitoringV1beta1 retrieves the MonitoringV1beta1Client
func (c *Clientset) MonitoringV1beta1() monitoringv1beta1.MonitoringV1beta1Interface {
	return c.monitoringV1beta1
}

// OsconfigV1beta1 retrieves the OsconfigV1beta1Client
func (c *Clientset) OsconfigV1beta1() osconfigv1beta1.OsconfigV1beta1Interface {
	return c.osconfigV1beta1
}

// PubsubV1beta1 retrieves the PubsubV1beta1Client
func (c *Clientset) PubsubV1beta1() pubsubv1beta1.PubsubV1beta1Interface {
	return c.pubsubV1beta1
}

// RedisV1beta1 retrieves the RedisV1beta1Client
func (c *Clientset) RedisV1beta1() redisv1beta1.RedisV1beta1Interface {
	return c.redisV1beta1
}

// ResourcemanagerV1beta1 retrieves the ResourcemanagerV1beta1Client
func (c *Clientset) ResourcemanagerV1beta1() resourcemanagerv1beta1.ResourcemanagerV1beta1Interface {
	return c.resourcemanagerV1beta1
}

// SecretmanagerV1beta1 retrieves the SecretmanagerV1beta1Client
func (c *Clientset) SecretmanagerV1beta1() secretmanagerv1beta1.SecretmanagerV1beta1Interface {
	return c.secretmanagerV1beta1
}

// ServicenetworkingV1beta1 retrieves the ServicenetworkingV1beta1Client
func (c *Clientset) ServicenetworkingV1beta1() servicenetworkingv1beta1.ServicenetworkingV1beta1Interface {
	return c.servicenetworkingV1beta1
}

// ServiceusageV1beta1 retrieves the ServiceusageV1beta1Client
func (c *Clientset) ServiceusageV1beta1() serviceusagev1beta1.ServiceusageV1beta1Interface {
	return c.serviceusageV1beta1
}

// SourcerepoV1beta1 retrieves the SourcerepoV1beta1Client
func (c *Clientset) SourcerepoV1beta1() sourcerepov1beta1.SourcerepoV1beta1Interface {
	return c.sourcerepoV1beta1
}

// SpannerV1beta1 retrieves the SpannerV1beta1Client
func (c *Clientset) SpannerV1beta1() spannerv1beta1.SpannerV1beta1Interface {
	return c.spannerV1beta1
}

// SqlV1beta1 retrieves the SqlV1beta1Client
func (c *Clientset) SqlV1beta1() sqlv1beta1.SqlV1beta1Interface {
	return c.sqlV1beta1
}

// StorageV1beta1 retrieves the StorageV1beta1Client
func (c *Clientset) StorageV1beta1() storagev1beta1.StorageV1beta1Interface {
	return c.storageV1beta1
}

// StoragetransferV1beta1 retrieves the StoragetransferV1beta1Client
func (c *Clientset) StoragetransferV1beta1() storagetransferv1beta1.StoragetransferV1beta1Interface {
	return c.storagetransferV1beta1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.accesscontextmanagerV1beta1, err = accesscontextmanagerv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.artifactregistryV1beta1, err = artifactregistryv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.bigqueryV1beta1, err = bigqueryv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.bigtableV1beta1, err = bigtablev1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.cloudbuildV1beta1, err = cloudbuildv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.cloudschedulerV1beta1, err = cloudschedulerv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.computeV1beta1, err = computev1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.containerV1beta1, err = containerv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.containeranalysisV1beta1, err = containeranalysisv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.dataflowV1beta1, err = dataflowv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.dataprocV1beta1, err = dataprocv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.dnsV1beta1, err = dnsv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.firestoreV1beta1, err = firestorev1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.gameservicesV1beta1, err = gameservicesv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.iamV1beta1, err = iamv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.iapV1beta1, err = iapv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.identityplatformV1beta1, err = identityplatformv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.k8sV1alpha1, err = k8sv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.kmsV1beta1, err = kmsv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.loggingV1beta1, err = loggingv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.memcacheV1beta1, err = memcachev1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.monitoringV1beta1, err = monitoringv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.osconfigV1beta1, err = osconfigv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.pubsubV1beta1, err = pubsubv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.redisV1beta1, err = redisv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.resourcemanagerV1beta1, err = resourcemanagerv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.secretmanagerV1beta1, err = secretmanagerv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.servicenetworkingV1beta1, err = servicenetworkingv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.serviceusageV1beta1, err = serviceusagev1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.sourcerepoV1beta1, err = sourcerepov1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.spannerV1beta1, err = spannerv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.sqlV1beta1, err = sqlv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.storageV1beta1, err = storagev1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.storagetransferV1beta1, err = storagetransferv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.accesscontextmanagerV1beta1 = accesscontextmanagerv1beta1.NewForConfigOrDie(c)
	cs.artifactregistryV1beta1 = artifactregistryv1beta1.NewForConfigOrDie(c)
	cs.bigqueryV1beta1 = bigqueryv1beta1.NewForConfigOrDie(c)
	cs.bigtableV1beta1 = bigtablev1beta1.NewForConfigOrDie(c)
	cs.cloudbuildV1beta1 = cloudbuildv1beta1.NewForConfigOrDie(c)
	cs.cloudschedulerV1beta1 = cloudschedulerv1beta1.NewForConfigOrDie(c)
	cs.computeV1beta1 = computev1beta1.NewForConfigOrDie(c)
	cs.containerV1beta1 = containerv1beta1.NewForConfigOrDie(c)
	cs.containeranalysisV1beta1 = containeranalysisv1beta1.NewForConfigOrDie(c)
	cs.dataflowV1beta1 = dataflowv1beta1.NewForConfigOrDie(c)
	cs.dataprocV1beta1 = dataprocv1beta1.NewForConfigOrDie(c)
	cs.dnsV1beta1 = dnsv1beta1.NewForConfigOrDie(c)
	cs.firestoreV1beta1 = firestorev1beta1.NewForConfigOrDie(c)
	cs.gameservicesV1beta1 = gameservicesv1beta1.NewForConfigOrDie(c)
	cs.iamV1beta1 = iamv1beta1.NewForConfigOrDie(c)
	cs.iapV1beta1 = iapv1beta1.NewForConfigOrDie(c)
	cs.identityplatformV1beta1 = identityplatformv1beta1.NewForConfigOrDie(c)
	cs.k8sV1alpha1 = k8sv1alpha1.NewForConfigOrDie(c)
	cs.kmsV1beta1 = kmsv1beta1.NewForConfigOrDie(c)
	cs.loggingV1beta1 = loggingv1beta1.NewForConfigOrDie(c)
	cs.memcacheV1beta1 = memcachev1beta1.NewForConfigOrDie(c)
	cs.monitoringV1beta1 = monitoringv1beta1.NewForConfigOrDie(c)
	cs.osconfigV1beta1 = osconfigv1beta1.NewForConfigOrDie(c)
	cs.pubsubV1beta1 = pubsubv1beta1.NewForConfigOrDie(c)
	cs.redisV1beta1 = redisv1beta1.NewForConfigOrDie(c)
	cs.resourcemanagerV1beta1 = resourcemanagerv1beta1.NewForConfigOrDie(c)
	cs.secretmanagerV1beta1 = secretmanagerv1beta1.NewForConfigOrDie(c)
	cs.servicenetworkingV1beta1 = servicenetworkingv1beta1.NewForConfigOrDie(c)
	cs.serviceusageV1beta1 = serviceusagev1beta1.NewForConfigOrDie(c)
	cs.sourcerepoV1beta1 = sourcerepov1beta1.NewForConfigOrDie(c)
	cs.spannerV1beta1 = spannerv1beta1.NewForConfigOrDie(c)
	cs.sqlV1beta1 = sqlv1beta1.NewForConfigOrDie(c)
	cs.storageV1beta1 = storagev1beta1.NewForConfigOrDie(c)
	cs.storagetransferV1beta1 = storagetransferv1beta1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.accesscontextmanagerV1beta1 = accesscontextmanagerv1beta1.New(c)
	cs.artifactregistryV1beta1 = artifactregistryv1beta1.New(c)
	cs.bigqueryV1beta1 = bigqueryv1beta1.New(c)
	cs.bigtableV1beta1 = bigtablev1beta1.New(c)
	cs.cloudbuildV1beta1 = cloudbuildv1beta1.New(c)
	cs.cloudschedulerV1beta1 = cloudschedulerv1beta1.New(c)
	cs.computeV1beta1 = computev1beta1.New(c)
	cs.containerV1beta1 = containerv1beta1.New(c)
	cs.containeranalysisV1beta1 = containeranalysisv1beta1.New(c)
	cs.dataflowV1beta1 = dataflowv1beta1.New(c)
	cs.dataprocV1beta1 = dataprocv1beta1.New(c)
	cs.dnsV1beta1 = dnsv1beta1.New(c)
	cs.firestoreV1beta1 = firestorev1beta1.New(c)
	cs.gameservicesV1beta1 = gameservicesv1beta1.New(c)
	cs.iamV1beta1 = iamv1beta1.New(c)
	cs.iapV1beta1 = iapv1beta1.New(c)
	cs.identityplatformV1beta1 = identityplatformv1beta1.New(c)
	cs.k8sV1alpha1 = k8sv1alpha1.New(c)
	cs.kmsV1beta1 = kmsv1beta1.New(c)
	cs.loggingV1beta1 = loggingv1beta1.New(c)
	cs.memcacheV1beta1 = memcachev1beta1.New(c)
	cs.monitoringV1beta1 = monitoringv1beta1.New(c)
	cs.osconfigV1beta1 = osconfigv1beta1.New(c)
	cs.pubsubV1beta1 = pubsubv1beta1.New(c)
	cs.redisV1beta1 = redisv1beta1.New(c)
	cs.resourcemanagerV1beta1 = resourcemanagerv1beta1.New(c)
	cs.secretmanagerV1beta1 = secretmanagerv1beta1.New(c)
	cs.servicenetworkingV1beta1 = servicenetworkingv1beta1.New(c)
	cs.serviceusageV1beta1 = serviceusagev1beta1.New(c)
	cs.sourcerepoV1beta1 = sourcerepov1beta1.New(c)
	cs.spannerV1beta1 = spannerv1beta1.New(c)
	cs.sqlV1beta1 = sqlv1beta1.New(c)
	cs.storageV1beta1 = storagev1beta1.New(c)
	cs.storagetransferV1beta1 = storagetransferv1beta1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
