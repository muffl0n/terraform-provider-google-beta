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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccHealthcareConsentStoreIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareConsentStoreIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_healthcare_consent_store_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("%s/consentStores/%s roles/viewer", fmt.Sprintf("projects/%s/locations/%s/datasets/tf-test-my-dataset%s", GetTestProjectFromEnv(), GetTestRegionFromEnv(), context["random_suffix"]), fmt.Sprintf("tf-test-my-consent-store%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccHealthcareConsentStoreIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_healthcare_consent_store_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("%s/consentStores/%s roles/viewer", fmt.Sprintf("projects/%s/locations/%s/datasets/tf-test-my-dataset%s", GetTestProjectFromEnv(), GetTestRegionFromEnv(), context["random_suffix"]), fmt.Sprintf("tf-test-my-consent-store%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccHealthcareConsentStoreIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccHealthcareConsentStoreIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_healthcare_consent_store_iam_member.foo",
				ImportStateId:     fmt.Sprintf("%s/consentStores/%s roles/viewer user:admin@hashicorptest.com", fmt.Sprintf("projects/%s/locations/%s/datasets/tf-test-my-dataset%s", GetTestProjectFromEnv(), GetTestRegionFromEnv(), context["random_suffix"]), fmt.Sprintf("tf-test-my-consent-store%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccHealthcareConsentStoreIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareConsentStoreIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_healthcare_consent_store_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("%s/consentStores/%s", fmt.Sprintf("projects/%s/locations/%s/datasets/tf-test-my-dataset%s", GetTestProjectFromEnv(), GetTestRegionFromEnv(), context["random_suffix"]), fmt.Sprintf("tf-test-my-consent-store%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccHealthcareConsentStoreIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_healthcare_consent_store_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("%s/consentStores/%s", fmt.Sprintf("projects/%s/locations/%s/datasets/tf-test-my-dataset%s", GetTestProjectFromEnv(), GetTestRegionFromEnv(), context["random_suffix"]), fmt.Sprintf("tf-test-my-consent-store%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccHealthcareConsentStoreIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_healthcare_dataset" "dataset" {
  location = "us-central1"
  name     = "tf-test-my-dataset%{random_suffix}"
}

resource "google_healthcare_consent_store" "my-consent" {
  dataset = google_healthcare_dataset.dataset.id
  name    = "tf-test-my-consent-store%{random_suffix}"
}

resource "google_healthcare_consent_store_iam_member" "foo" {
  dataset = google_healthcare_consent_store.my-consent.dataset
  consent_store_id = google_healthcare_consent_store.my-consent.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccHealthcareConsentStoreIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_healthcare_dataset" "dataset" {
  location = "us-central1"
  name     = "tf-test-my-dataset%{random_suffix}"
}

resource "google_healthcare_consent_store" "my-consent" {
  dataset = google_healthcare_dataset.dataset.id
  name    = "tf-test-my-consent-store%{random_suffix}"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_healthcare_consent_store_iam_policy" "foo" {
  dataset = google_healthcare_consent_store.my-consent.dataset
  consent_store_id = google_healthcare_consent_store.my-consent.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccHealthcareConsentStoreIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_healthcare_dataset" "dataset" {
  location = "us-central1"
  name     = "tf-test-my-dataset%{random_suffix}"
}

resource "google_healthcare_consent_store" "my-consent" {
  dataset = google_healthcare_dataset.dataset.id
  name    = "tf-test-my-consent-store%{random_suffix}"
}

data "google_iam_policy" "foo" {
}

resource "google_healthcare_consent_store_iam_policy" "foo" {
  dataset = google_healthcare_consent_store.my-consent.dataset
  consent_store_id = google_healthcare_consent_store.my-consent.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccHealthcareConsentStoreIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_healthcare_dataset" "dataset" {
  location = "us-central1"
  name     = "tf-test-my-dataset%{random_suffix}"
}

resource "google_healthcare_consent_store" "my-consent" {
  dataset = google_healthcare_dataset.dataset.id
  name    = "tf-test-my-consent-store%{random_suffix}"
}

resource "google_healthcare_consent_store_iam_binding" "foo" {
  dataset = google_healthcare_consent_store.my-consent.dataset
  consent_store_id = google_healthcare_consent_store.my-consent.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccHealthcareConsentStoreIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_healthcare_dataset" "dataset" {
  location = "us-central1"
  name     = "tf-test-my-dataset%{random_suffix}"
}

resource "google_healthcare_consent_store" "my-consent" {
  dataset = google_healthcare_dataset.dataset.id
  name    = "tf-test-my-consent-store%{random_suffix}"
}

resource "google_healthcare_consent_store_iam_binding" "foo" {
  dataset = google_healthcare_consent_store.my-consent.dataset
  consent_store_id = google_healthcare_consent_store.my-consent.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
