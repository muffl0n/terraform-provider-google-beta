// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSpannerInstanceIamBinding(t *testing.T) {
	t.Parallel()

	account := fmt.Sprintf("tf-test-%d", RandInt(t))
	role := "roles/spanner.databaseAdmin"
	project := GetTestProjectFromEnv()
	instance := fmt.Sprintf("tf-test-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSpannerInstanceIamBinding_basic(account, instance, role),
			},
			{
				ResourceName: "google_spanner_instance_iam_binding.foo",
				ImportStateId: fmt.Sprintf("%s %s", spannerInstanceId{
					Project:  project,
					Instance: instance,
				}.terraformId(), role),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccSpannerInstanceIamBinding_update(account, instance, role),
			},
			{
				ResourceName: "google_spanner_instance_iam_binding.foo",
				ImportStateId: fmt.Sprintf("%s %s", spannerInstanceId{
					Project:  project,
					Instance: instance,
				}.terraformId(), role),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSpannerInstanceIamMember(t *testing.T) {
	t.Parallel()

	project := GetTestProjectFromEnv()
	account := fmt.Sprintf("tf-test-%d", RandInt(t))
	role := "roles/spanner.databaseAdmin"
	instance := fmt.Sprintf("tf-test-%s", RandString(t, 10))
	conditionTitle := "Access only database one"

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccSpannerInstanceIamMember_basic(account, instance, role),
			},
			{
				ResourceName: "google_spanner_instance_iam_member.foo",
				ImportStateId: fmt.Sprintf("%s %s serviceAccount:%s@%s.iam.gserviceaccount.com %s", spannerInstanceId{
					Instance: instance,
					Project:  project,
				}.terraformId(), role, account, project, conditionTitle),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSpannerInstanceIamPolicy(t *testing.T) {
	t.Parallel()

	project := GetTestProjectFromEnv()
	account := fmt.Sprintf("tf-test-%d", RandInt(t))
	role := "roles/spanner.databaseAdmin"
	instance := fmt.Sprintf("tf-test-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSpannerInstanceIamPolicy_basic(account, instance, role),
			},
			// Test a few import formats
			{
				ResourceName: "google_spanner_instance_iam_policy.foo",
				ImportStateId: spannerInstanceId{
					Instance: instance,
					Project:  project,
				}.terraformId(),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSpannerInstanceIamBinding_basic(account, instance, roleId string) string {
	return fmt.Sprintf(`
resource "google_service_account" "test_account" {
  account_id   = "%s"
  display_name = "Spanner Instance Iam Testing Account"
}

resource "google_spanner_instance" "instance" {
  name         = "%s"
  config       = "regional-us-central1"
  display_name = "%s"
  num_nodes    = 1
}

resource "google_spanner_instance_iam_binding" "foo" {
  project  = google_spanner_instance.instance.project
  instance = google_spanner_instance.instance.name
  role     = "%s"
  members  = ["serviceAccount:${google_service_account.test_account.email}"]
}
`, account, instance, instance, roleId)
}

func testAccSpannerInstanceIamBinding_update(account, instance, roleId string) string {
	return fmt.Sprintf(`
resource "google_service_account" "test_account" {
  account_id   = "%s"
  display_name = "Spanner Instance Iam Testing Account"
}

resource "google_service_account" "test_account_2" {
  account_id   = "%s-2"
  display_name = "Spanner Instance Iam Testing Account"
}

resource "google_spanner_instance" "instance" {
  name         = "%s"
  config       = "regional-us-central1"
  display_name = "%s"
  num_nodes    = 1
}

resource "google_spanner_instance_iam_binding" "foo" {
  project  = google_spanner_instance.instance.project
  instance = google_spanner_instance.instance.name
  role     = "%s"
  members = [
    "serviceAccount:${google_service_account.test_account.email}",
    "serviceAccount:${google_service_account.test_account_2.email}",
  ]
}
`, account, account, instance, instance, roleId)
}

func testAccSpannerInstanceIamMember_basic(account, instance, roleId string) string {
	return fmt.Sprintf(`
resource "google_service_account" "test_account" {
  account_id   = "%s"
  display_name = "Spanner Instance Iam Testing Account"
}

resource "google_spanner_instance" "instance" {
  name         = "%s"
  config       = "regional-us-central1"
  display_name = "%s"
  num_nodes    = 1
}

resource "google_spanner_instance_iam_member" "foo" {
  project  = google_spanner_instance.instance.project
  instance = google_spanner_instance.instance.name
  role     = "%s"
  member   = "serviceAccount:${google_service_account.test_account.email}"
  condition {
    title      = "Access only database one"
    expression = "resource.type == \"spanner.googleapis.com/DatabaseRole\" && resource.name.endsWith(\"/databaseRoles/parent\")"
  }
}
`, account, instance, instance, roleId)
}

func testAccSpannerInstanceIamPolicy_basic(account, instance, roleId string) string {
	return fmt.Sprintf(`
resource "google_service_account" "test_account" {
  account_id   = "%s"
  display_name = "Spanner Instance Iam Testing Account"
}

resource "google_spanner_instance" "instance" {
  name         = "%s"
  config       = "regional-us-central1"
  display_name = "%s"
  num_nodes    = 1
}

data "google_iam_policy" "foo" {
  binding {
    role = "%s"

    members = ["serviceAccount:${google_service_account.test_account.email}"]
  }
}

resource "google_spanner_instance_iam_policy" "foo" {
  project     = google_spanner_instance.instance.project
  instance    = google_spanner_instance.instance.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, account, instance, instance, roleId)
}
