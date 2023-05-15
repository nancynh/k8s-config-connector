// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccComputeInterconnectAttachment_interconnectAttachmentBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeInterconnectAttachmentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInterconnectAttachment_interconnectAttachmentBasicExample(context),
			},
			{
				ResourceName:            "google_compute_interconnect_attachment.on_prem",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"router", "candidate_subnets", "region"},
			},
		},
	})
}

func testAccComputeInterconnectAttachment_interconnectAttachmentBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_interconnect_attachment" "on_prem" {
  name                     = "tf-test-on-prem-attachment%{random_suffix}"
  edge_availability_domain = "AVAILABILITY_DOMAIN_1"
  type                     = "PARTNER"
  router                   = google_compute_router.foobar.id
  mtu                      = 1500
}

resource "google_compute_router" "foobar" {
  name    = "tf-test-router-1%{random_suffix}"
  network = google_compute_network.foobar.name
  bgp {
    asn = 16550
  }
}

resource "google_compute_network" "foobar" {
  name                    = "tf-test-network-1%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccComputeInterconnectAttachment_computeInterconnectAttachmentIpsecEncryptionExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeInterconnectAttachmentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInterconnectAttachment_computeInterconnectAttachmentIpsecEncryptionExample(context),
			},
			{
				ResourceName:            "google_compute_interconnect_attachment.ipsec-encrypted-interconnect-attachment",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"router", "candidate_subnets", "region"},
			},
		},
	})
}

func testAccComputeInterconnectAttachment_computeInterconnectAttachmentIpsecEncryptionExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_interconnect_attachment" "ipsec-encrypted-interconnect-attachment" {
  name                     = "tf-test-test-interconnect-attachment%{random_suffix}"
  edge_availability_domain = "AVAILABILITY_DOMAIN_1"
  type                     = "PARTNER"
  router                   = google_compute_router.router.id
  encryption               = "IPSEC"
  ipsec_internal_addresses = [
    google_compute_address.address.self_link,
  ]
}

resource "google_compute_address" "address" {
  name          = "tf-test-test-address%{random_suffix}"
  address_type  = "INTERNAL"
  purpose       = "IPSEC_INTERCONNECT"
  address       = "192.168.1.0"
  prefix_length = 29
  network       = google_compute_network.network.self_link
}

resource "google_compute_router" "router" {
  name                          = "tf-test-test-router%{random_suffix}"
  network                       = google_compute_network.network.name
  encrypted_interconnect_router = true
  bgp {
    asn = 16550
  }
}

resource "google_compute_network" "network" {
  name                    = "tf-test-test-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func testAccCheckComputeInterconnectAttachmentDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_interconnect_attachment" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeInterconnectAttachment still exists at %s", url)
			}
		}

		return nil
	}
}
