// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

package compute_test

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

func TestAccComputeRegionTargetTcpProxy_regionTargetTcpProxyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionTargetTcpProxyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionTargetTcpProxy_regionTargetTcpProxyBasicExample(context),
			},
			{
				ResourceName:            "google_compute_region_target_tcp_proxy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "region"},
			},
		},
	})
}

func testAccComputeRegionTargetTcpProxy_regionTargetTcpProxyBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_target_tcp_proxy" "default" {
  name            = "tf-test-test-proxy%{random_suffix}"
  region          = "europe-west4"
  backend_service = google_compute_region_backend_service.default.id
}

resource "google_compute_region_backend_service" "default" {
  name        = "tf-test-backend-service%{random_suffix}"
  protocol    = "TCP"
  timeout_sec = 10
  region      = "europe-west4"

  health_checks         = [google_compute_region_health_check.default.id]
  load_balancing_scheme = "INTERNAL_MANAGED"
}

resource "google_compute_region_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  region             = "europe-west4"
  timeout_sec        = 1
  check_interval_sec = 1
  
  tcp_health_check {
    port = "80"
  }
}
`, context)
}

func testAccCheckComputeRegionTargetTcpProxyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_region_target_tcp_proxy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetTcpProxies/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ComputeRegionTargetTcpProxy still exists at %s", url)
			}
		}

		return nil
	}
}
