package gmail

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

func init() {
	resource.AddTestSweepers("gmail_label", &resource.Sweeper{
		Name: "gmail_label",
		F:    testSweepLabels,
	})
}

func testSweepLabels(_ string) error {
	srv := testAccProvider.Meta().(*gmail.Service)

	labels, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		return fmt.Errorf("error retrieving labels: %v", err)
	}

	for _, label := range labels.Labels {
		if err := srv.Users.Labels.Delete(user, label.Id).Do(); err != nil {
			return fmt.Errorf("error deleting label %s: %v", label.Name, err)
		}
	}

	return nil
}

func TestAccGmailLabelResource_basic(t *testing.T) {
	labelName1 := "terraform-label-" + acctest.RandString(10)
	labelName2 := "terraform-label-" + acctest.RandString(10)
	var label1, label2 gmail.Label

	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckGmailLabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGmailLabelResourceConfigBasic(labelName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmailLabelExists("gmail_label.this", &label1),
					resource.TestCheckResourceAttr("gmail_label.this", "name", labelName1),
					resource.TestCheckResourceAttr("gmail_label.this", "label_list_visibility", "labelShow"),
					resource.TestCheckResourceAttr("gmail_label.this", "message_list_visibility", "show"),
					resource.TestCheckResourceAttr("gmail_label.this", "background_color", "#999999"),
					resource.TestCheckResourceAttr("gmail_label.this", "text_color", "#f3f3f3"),
				),
			},
			{
				Config: testAccGmailLabelResourceConfigUpdated(labelName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmailLabelExists("gmail_label.this", &label2),
					resource.TestCheckResourceAttr("gmail_label.this", "name", labelName2),
					resource.TestCheckResourceAttr("gmail_label.this", "label_list_visibility", "labelShowIfUnread"),
					resource.TestCheckResourceAttr("gmail_label.this", "message_list_visibility", "hide"),
					resource.TestCheckResourceAttr("gmail_label.this", "background_color", "#a46a21"),
					resource.TestCheckResourceAttr("gmail_label.this", "text_color", "#f691b2"),
					testAccCheckGmailLabelForceNew(&label1, &label2, false),
				),
			},
		},
	})
}

func TestAccGmailLabelResource_import(t *testing.T) {
	labelName := "terraform-label-" + acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckGmailLabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGmailLabelResourceConfigBasic(labelName),
			},
			{
				ResourceName:      "gmail_label.this",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckGmailLabelExists(resourceName string, label *gmail.Label) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found in state", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID set for resource %s", resourceName)
		}

		srv := testAccProvider.Meta().(*gmail.Service)

		fetchedLabel, err := srv.Users.Labels.Get(user, rs.Primary.ID).Do()
		if err != nil {
			return fmt.Errorf("error fetching resource %s: %v", resourceName, err)
		}

		*label = *fetchedLabel
		return nil
	}
}

func testAccCheckGmailLabelDestroy(s *terraform.State) error {
	srv := testAccProvider.Meta().(*gmail.Service)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "gmail_label" {
			continue
		}

		if label, err := srv.Users.Labels.Get(user, rs.Primary.ID).Do(); err == nil {
			return fmt.Errorf("label %s still exists", label.Name)
		} else if apiError, ok := err.(*googleapi.Error); !ok {
			return fmt.Errorf("invalid error returned: %v", err)
		} else if apiError.Code != http.StatusNotFound {
			return fmt.Errorf("could not fetch label: %v", err)
		}
	}

	return nil
}

func testAccCheckGmailLabelForceNew(old, new *gmail.Label, wantNew bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if wantNew {
			if old.Id == new.Id {
				return fmt.Errorf("expecting label ID to differ, got %+v and %+v", old, new)
			}
		} else {
			if old.Id != new.Id {
				return fmt.Errorf("expecting label ID to be equal, got %+v and %+v", old, new)
			}
		}
		return nil
	}
}

func testAccGmailLabelResourceConfigBasic(labelName string) string {
	return fmt.Sprintf(`
resource gmail_label this {
  name = "%s"
	label_list_visibility = "labelShow"
	message_list_visibility = "show"
	background_color = "#999999"
	text_color = "#f3f3f3"
}
`, labelName)
}

func testAccGmailLabelResourceConfigUpdated(labelName string) string {
	return fmt.Sprintf(`
resource gmail_label this {
  name = "%s"
	label_list_visibility = "labelShowIfUnread"
	message_list_visibility = "hide"
	background_color = "#a46a21"
	text_color = "#f691b2"
}
`, labelName)
}
