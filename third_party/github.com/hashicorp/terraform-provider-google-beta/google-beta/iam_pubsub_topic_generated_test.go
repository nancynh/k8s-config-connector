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

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccPubsubTopicIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubTopicIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_pubsub_topic_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/topics/%s roles/viewer", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-topic%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccPubsubTopicIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_pubsub_topic_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/topics/%s roles/viewer", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-topic%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPubsubTopicIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccPubsubTopicIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_pubsub_topic_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/topics/%s roles/viewer user:admin@hashicorptest.com", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-topic%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPubsubTopicIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubTopicIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_pubsub_topic_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/topics/%s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-topic%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccPubsubTopicIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_pubsub_topic_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/topics/%s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-topic%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPubsubTopicIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"

  labels = {
    foo = "bar"
  }

  message_retention_duration = "86600s"
}

resource "google_pubsub_topic_iam_member" "foo" {
  project = google_pubsub_topic.example.project
  topic = google_pubsub_topic.example.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccPubsubTopicIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"

  labels = {
    foo = "bar"
  }

  message_retention_duration = "86600s"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_pubsub_topic_iam_policy" "foo" {
  project = google_pubsub_topic.example.project
  topic = google_pubsub_topic.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccPubsubTopicIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"

  labels = {
    foo = "bar"
  }

  message_retention_duration = "86600s"
}

data "google_iam_policy" "foo" {
}

resource "google_pubsub_topic_iam_policy" "foo" {
  project = google_pubsub_topic.example.project
  topic = google_pubsub_topic.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccPubsubTopicIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"

  labels = {
    foo = "bar"
  }

  message_retention_duration = "86600s"
}

resource "google_pubsub_topic_iam_binding" "foo" {
  project = google_pubsub_topic.example.project
  topic = google_pubsub_topic.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccPubsubTopicIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"

  labels = {
    foo = "bar"
  }

  message_retention_duration = "86600s"
}

resource "google_pubsub_topic_iam_binding" "foo" {
  project = google_pubsub_topic.example.project
  topic = google_pubsub_topic.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
