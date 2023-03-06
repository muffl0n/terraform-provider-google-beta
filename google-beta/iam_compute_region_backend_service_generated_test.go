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

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeRegionBackendServiceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendServiceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_region_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/backendServices/%s roles/compute.admin", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-region-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccComputeRegionBackendServiceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_compute_region_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/backendServices/%s roles/compute.admin", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-region-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendServiceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccComputeRegionBackendServiceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_region_backend_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/backendServices/%s roles/compute.admin user:admin@hashicorptest.com", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-region-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendServiceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendServiceIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_region_backend_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/backendServices/%s", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-region-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionBackendServiceIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_compute_region_backend_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/backendServices/%s", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-region-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendServiceIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendServiceIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_region_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/backendServices/%s roles/compute.admin %s", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-region-service%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendServiceIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendServiceIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_region_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/backendServices/%s roles/compute.admin", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-region-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_region_backend_service_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/backendServices/%s roles/compute.admin %s", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-region-service%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_region_backend_service_iam_binding.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/backendServices/%s roles/compute.admin %s", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-region-service%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendServiceIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendServiceIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_region_backend_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/backendServices/%s roles/compute.admin user:admin@hashicorptest.com %s", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-region-service%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendServiceIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendServiceIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_region_backend_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/backendServices/%s roles/compute.admin user:admin@hashicorptest.com", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-region-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_region_backend_service_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/backendServices/%s roles/compute.admin user:admin@hashicorptest.com %s", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-region-service%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_region_backend_service_iam_member.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/backendServices/%s roles/compute.admin user:admin@hashicorptest.com %s", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-region-service%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendServiceIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.admin",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	// Test should have 2 bindings: one with a description and one without. Any < chars are converted to a unicode character by the API.
	expectedPolicyData := Nprintf(`{"bindings":[{"condition":{"description":"%{condition_desc}","expression":"%{condition_expr}","title":"%{condition_title}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"},{"condition":{"expression":"%{condition_expr}","title":"%{condition_title}-no-description"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"}]}`, context)
	expectedPolicyData = strings.Replace(expectedPolicyData, "<", "\\u003c", -1)

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendServiceIamPolicy_withConditionGenerated(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO(SarahFrench) - uncomment once https://github.com/GoogleCloudPlatform/magic-modules/pull/6466 merged
					// resource.TestCheckResourceAttr("data.google_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttr("google_compute_region_backend_service_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", checkGoogleIamPolicy),
				),
			},
			{
				ResourceName:      "google_compute_region_backend_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/backendServices/%s", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-region-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionBackendServiceIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_backend_service" "default" {
  name                            = "tf-test-region-service%{random_suffix}"
  region                          = "us-central1"
  health_checks                   = [google_compute_health_check.default.id]
  connection_draining_timeout_sec = 10
  session_affinity                = "CLIENT_IP"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-rbs-health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_region_backend_service_iam_member" "foo" {
  project = google_compute_region_backend_service.default.project
  region = google_compute_region_backend_service.default.region
  name = google_compute_region_backend_service.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccComputeRegionBackendServiceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_backend_service" "default" {
  name                            = "tf-test-region-service%{random_suffix}"
  region                          = "us-central1"
  health_checks                   = [google_compute_health_check.default.id]
  connection_draining_timeout_sec = 10
  session_affinity                = "CLIENT_IP"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-rbs-health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_compute_region_backend_service_iam_policy" "foo" {
  project = google_compute_region_backend_service.default.project
  region = google_compute_region_backend_service.default.region
  name = google_compute_region_backend_service.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeRegionBackendServiceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_backend_service" "default" {
  name                            = "tf-test-region-service%{random_suffix}"
  region                          = "us-central1"
  health_checks                   = [google_compute_health_check.default.id]
  connection_draining_timeout_sec = 10
  session_affinity                = "CLIENT_IP"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-rbs-health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

data "google_iam_policy" "foo" {
}

resource "google_compute_region_backend_service_iam_policy" "foo" {
  project = google_compute_region_backend_service.default.project
  region = google_compute_region_backend_service.default.region
  name = google_compute_region_backend_service.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeRegionBackendServiceIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_backend_service" "default" {
  name                            = "tf-test-region-service%{random_suffix}"
  region                          = "us-central1"
  health_checks                   = [google_compute_health_check.default.id]
  connection_draining_timeout_sec = 10
  session_affinity                = "CLIENT_IP"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-rbs-health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_region_backend_service_iam_binding" "foo" {
  project = google_compute_region_backend_service.default.project
  region = google_compute_region_backend_service.default.region
  name = google_compute_region_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccComputeRegionBackendServiceIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_backend_service" "default" {
  name                            = "tf-test-region-service%{random_suffix}"
  region                          = "us-central1"
  health_checks                   = [google_compute_health_check.default.id]
  connection_draining_timeout_sec = 10
  session_affinity                = "CLIENT_IP"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-rbs-health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_region_backend_service_iam_binding" "foo" {
  project = google_compute_region_backend_service.default.project
  region = google_compute_region_backend_service.default.region
  name = google_compute_region_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccComputeRegionBackendServiceIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_backend_service" "default" {
  name                            = "tf-test-region-service%{random_suffix}"
  region                          = "us-central1"
  health_checks                   = [google_compute_health_check.default.id]
  connection_draining_timeout_sec = 10
  session_affinity                = "CLIENT_IP"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-rbs-health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_region_backend_service_iam_binding" "foo" {
  project = google_compute_region_backend_service.default.project
  region = google_compute_region_backend_service.default.region
  name = google_compute_region_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccComputeRegionBackendServiceIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_backend_service" "default" {
  name                            = "tf-test-region-service%{random_suffix}"
  region                          = "us-central1"
  health_checks                   = [google_compute_health_check.default.id]
  connection_draining_timeout_sec = 10
  session_affinity                = "CLIENT_IP"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-rbs-health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_region_backend_service_iam_binding" "foo" {
  project = google_compute_region_backend_service.default.project
  region = google_compute_region_backend_service.default.region
  name = google_compute_region_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_compute_region_backend_service_iam_binding" "foo2" {
  project = google_compute_region_backend_service.default.project
  region = google_compute_region_backend_service.default.region
  name = google_compute_region_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_compute_region_backend_service_iam_binding" "foo3" {
  project = google_compute_region_backend_service.default.project
  region = google_compute_region_backend_service.default.region
  name = google_compute_region_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccComputeRegionBackendServiceIamMember_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_backend_service" "default" {
  name                            = "tf-test-region-service%{random_suffix}"
  region                          = "us-central1"
  health_checks                   = [google_compute_health_check.default.id]
  connection_draining_timeout_sec = 10
  session_affinity                = "CLIENT_IP"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-rbs-health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_region_backend_service_iam_member" "foo" {
  project = google_compute_region_backend_service.default.project
  region = google_compute_region_backend_service.default.region
  name = google_compute_region_backend_service.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccComputeRegionBackendServiceIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_backend_service" "default" {
  name                            = "tf-test-region-service%{random_suffix}"
  region                          = "us-central1"
  health_checks                   = [google_compute_health_check.default.id]
  connection_draining_timeout_sec = 10
  session_affinity                = "CLIENT_IP"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-rbs-health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_region_backend_service_iam_member" "foo" {
  project = google_compute_region_backend_service.default.project
  region = google_compute_region_backend_service.default.region
  name = google_compute_region_backend_service.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_compute_region_backend_service_iam_member" "foo2" {
  project = google_compute_region_backend_service.default.project
  region = google_compute_region_backend_service.default.region
  name = google_compute_region_backend_service.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_compute_region_backend_service_iam_member" "foo3" {
  project = google_compute_region_backend_service.default.project
  region = google_compute_region_backend_service.default.region
  name = google_compute_region_backend_service.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccComputeRegionBackendServiceIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_backend_service" "default" {
  name                            = "tf-test-region-service%{random_suffix}"
  region                          = "us-central1"
  health_checks                   = [google_compute_health_check.default.id]
  connection_draining_timeout_sec = 10
  session_affinity                = "CLIENT_IP"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-rbs-health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      # Check that lack of description doesn't cause any issues
      # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
      title       = "%{condition_title_no_desc}"
      expression  = "%{condition_expr_no_desc}"
    }
  }
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "%{condition_desc}"
      expression  = "%{condition_expr}"
    }
  }
}

resource "google_compute_region_backend_service_iam_policy" "foo" {
  project = google_compute_region_backend_service.default.project
  region = google_compute_region_backend_service.default.region
  name = google_compute_region_backend_service.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
