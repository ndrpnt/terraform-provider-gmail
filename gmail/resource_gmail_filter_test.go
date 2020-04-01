package gmail

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"google.golang.org/api/gmail/v1"
)

func init() {
	resource.AddTestSweepers("gmail_filter", &resource.Sweeper{
		Name: "gmail_filter",
		F:    testSweepFilters,
	})
}

func testSweepFilters(_ string) error {
	srv, err := gmailService()
	if err != nil {
		return fmt.Errorf("could not retrieve gmail service: %v", err)
	}

	filters, err := srv.Users.Settings.Filters.List(user).Do()
	if err != nil {
		return fmt.Errorf("could not fetch filters: %v", err)
	}

	for _, filter := range filters.Filter {
		if err := srv.Users.Settings.Filters.Delete(user, filter.Id).Do(); err != nil {
			return fmt.Errorf("could not delete filter %s: %v", filter.Id, err)
		}
	}

	return nil
}

func TestAccGmailFilterResource_basic(t *testing.T) {
	labelName := testLabelPrefix + acctest.RandString(10)
	var filter1, filter2 gmail.Filter

	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckGmailFilterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGmailFilterResourceConfigBasic(labelName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmailFilterExists("gmail_filter.this", &filter1),
					resource.TestCheckResourceAttr("gmail_filter.this", "criteria_exclude_chats", "false"),
					resource.TestCheckResourceAttr("gmail_filter.this", "criteria_from", ""),
					resource.TestCheckResourceAttr("gmail_filter.this", "criteria_has_attachment", "false"),
					resource.TestCheckResourceAttr("gmail_filter.this", "criteria_negated_query", ""),
					resource.TestCheckResourceAttr("gmail_filter.this", "criteria_query", ""),
					resource.TestCheckResourceAttr("gmail_filter.this", "criteria_size_larger", "25"),
					resource.TestCheckNoResourceAttr("gmail_filter.this", "criteria_size_smaller"),
					resource.TestCheckResourceAttr("gmail_filter.this", "criteria_subject", ""),
					resource.TestCheckResourceAttr("gmail_filter.this", "criteria_to", ""),
					resource.TestCheckResourceAttr("gmail_filter.this", "action_add_label_ids.#", "1"),
					resource.TestCheckResourceAttr("gmail_filter.this", "action_forward", ""),
					resource.TestCheckResourceAttr("gmail_filter.this", "action_remove_label_ids.#", "0"),
				),
			},
			{
				Config: testAccGmailFilterResourceConfigUpdated(labelName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmailFilterExists("gmail_filter.this", &filter1),
					resource.TestCheckResourceAttr("gmail_filter.this", "criteria_exclude_chats", "false"),
					resource.TestCheckResourceAttr("gmail_filter.this", "criteria_from", "foobar"),
					resource.TestCheckResourceAttr("gmail_filter.this", "criteria_has_attachment", "true"),
					resource.TestCheckResourceAttr("gmail_filter.this", "criteria_negated_query", ""),
					resource.TestCheckResourceAttr("gmail_filter.this", "criteria_query", "from:foo to:bar"),
					resource.TestCheckNoResourceAttr("gmail_filter.this", "criteria_size_larger"),
					resource.TestCheckNoResourceAttr("gmail_filter.this", "criteria_size_smaller"),
					resource.TestCheckResourceAttr("gmail_filter.this", "criteria_subject", ""),
					resource.TestCheckResourceAttr("gmail_filter.this", "criteria_to", "bazqux"),
					resource.TestCheckResourceAttr("gmail_filter.this", "action_add_label_ids.#", "2"),
					resource.TestCheckResourceAttr("gmail_filter.this", "action_forward", ""),
					resource.TestCheckResourceAttr("gmail_filter.this", "action_remove_label_ids.#", "1"),
					testAccCheckGmailFilterForceNew(&filter1, &filter2, true),
				),
			},
		},
	})
}

func TestAccGmailFilterResource_import(t *testing.T) {
	labelName := testLabelPrefix + acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckGmailFilterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGmailFilterResourceConfigBasic(labelName),
			},
			{
				ResourceName:      "gmail_filter.this",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckGmailFilterExists(resourceName string, filter *gmail.Filter) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found in state", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID set for resource %s", resourceName)
		}

		srv := testAccProvider.Meta().(*gmail.Service)

		fetchedFilter, err := srv.Users.Settings.Filters.Get(user, rs.Primary.ID).Do()
		if err != nil {
			return fmt.Errorf("error fetching resource %s: %v", resourceName, err)
		}

		*filter = *fetchedFilter
		return nil
	}
}

func testAccCheckGmailFilterDestroy(s *terraform.State) error {
	srv := testAccProvider.Meta().(*gmail.Service)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "gmail_filter" {
			continue
		}

		if filter, err := srv.Users.Settings.Filters.Get(user, rs.Primary.ID).Do(); err == nil {
			return fmt.Errorf("filter %s still exists", filter.Id)
		} else if !is404Error(err) {
			return fmt.Errorf("could not fetch filter: %v", err)
		}
	}

	return nil
}

func testAccCheckGmailFilterForceNew(old, new *gmail.Filter, wantNew bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if wantNew {
			if old.Id == new.Id {
				return fmt.Errorf("expecting filter ID to differ, got %+v and %+v", old, new)
			}
		} else {
			if old.Id != new.Id {
				return fmt.Errorf("expecting filter ID to be equal, got %+v and %+v", old, new)
			}
		}
		return nil
	}
}

func testAccGmailFilterResourceConfigBasic(labelName string) string {
	return fmt.Sprintf(`
resource gmail_label this {
  name = "%s"
}

resource gmail_filter this {
  criteria_size_larger = 25
	action_add_label_ids = [gmail_label.this.id]
}
`, labelName)
}

func testAccGmailFilterResourceConfigUpdated(labelName string) string {
	return fmt.Sprintf(`
resource gmail_label this {
  name = "%s"
}

data gmail_label important {
  name = "IMPORTANT"
}

data gmail_label inbox {
  name = "INBOX"
}

resource gmail_filter this {
	criteria_from = "foobar"
	criteria_has_attachment = true
	criteria_query = "from:foo to:bar"
	criteria_to = "bazqux"
	action_add_label_ids = [gmail_label.this.id, data.gmail_label.important.id]
	action_remove_label_ids = [data.gmail_label.inbox.id]
}
`, labelName)
}
