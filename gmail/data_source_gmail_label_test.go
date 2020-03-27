package gmail

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGmailLabelDataSource_system(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccGmailLabelDataSourceConfigSystem,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.gmail_label.chat", "id", "CHAT"),
					resource.TestCheckResourceAttr("data.gmail_label.chat", "label_list_visibility", "labelHide"),
					resource.TestCheckResourceAttr("data.gmail_label.chat", "message_list_visibility", "hide"),
					resource.TestCheckResourceAttr("data.gmail_label.chat", "name", "CHAT"),
					resource.TestCheckNoResourceAttr("data.gmail_label.chat", "background_color"),
					resource.TestCheckNoResourceAttr("data.gmail_label.chat", "text_color"),
					resource.TestCheckResourceAttr("data.gmail_label.chat", "type", "system"),

					resource.TestCheckResourceAttr("data.gmail_label.sent", "id", "SENT"),
					resource.TestCheckResourceAttr("data.gmail_label.sent", "label_list_visibility", ""),
					resource.TestCheckResourceAttr("data.gmail_label.sent", "message_list_visibility", ""),
					resource.TestCheckResourceAttr("data.gmail_label.sent", "name", "SENT"),
					resource.TestCheckNoResourceAttr("data.gmail_label.sent", "background_color"),
					resource.TestCheckNoResourceAttr("data.gmail_label.sent", "text_color"),
					resource.TestCheckResourceAttr("data.gmail_label.sent", "type", "system"),

					resource.TestCheckResourceAttr("data.gmail_label.inbox", "id", "INBOX"),
					resource.TestCheckResourceAttr("data.gmail_label.inbox", "label_list_visibility", ""),
					resource.TestCheckResourceAttr("data.gmail_label.inbox", "message_list_visibility", ""),
					resource.TestCheckResourceAttr("data.gmail_label.inbox", "name", "INBOX"),
					resource.TestCheckNoResourceAttr("data.gmail_label.inbox", "background_color"),
					resource.TestCheckNoResourceAttr("data.gmail_label.inbox", "text_color"),
					resource.TestCheckResourceAttr("data.gmail_label.inbox", "type", "system"),

					resource.TestCheckResourceAttr("data.gmail_label.important", "id", "IMPORTANT"),
					resource.TestCheckResourceAttr("data.gmail_label.important", "label_list_visibility", "labelHide"),
					resource.TestCheckResourceAttr("data.gmail_label.important", "message_list_visibility", "hide"),
					resource.TestCheckResourceAttr("data.gmail_label.important", "name", "IMPORTANT"),
					resource.TestCheckNoResourceAttr("data.gmail_label.important", "background_color"),
					resource.TestCheckNoResourceAttr("data.gmail_label.important", "text_color"),
					resource.TestCheckResourceAttr("data.gmail_label.important", "type", "system"),

					resource.TestCheckResourceAttr("data.gmail_label.trash", "id", "TRASH"),
					resource.TestCheckResourceAttr("data.gmail_label.trash", "label_list_visibility", "labelHide"),
					resource.TestCheckResourceAttr("data.gmail_label.trash", "message_list_visibility", "hide"),
					resource.TestCheckResourceAttr("data.gmail_label.trash", "name", "TRASH"),
					resource.TestCheckNoResourceAttr("data.gmail_label.trash", "background_color"),
					resource.TestCheckNoResourceAttr("data.gmail_label.trash", "text_color"),
					resource.TestCheckResourceAttr("data.gmail_label.trash", "type", "system"),

					resource.TestCheckResourceAttr("data.gmail_label.draft", "id", "DRAFT"),
					resource.TestCheckResourceAttr("data.gmail_label.draft", "label_list_visibility", ""),
					resource.TestCheckResourceAttr("data.gmail_label.draft", "message_list_visibility", ""),
					resource.TestCheckResourceAttr("data.gmail_label.draft", "name", "DRAFT"),
					resource.TestCheckNoResourceAttr("data.gmail_label.draft", "background_color"),
					resource.TestCheckNoResourceAttr("data.gmail_label.draft", "text_color"),
					resource.TestCheckResourceAttr("data.gmail_label.draft", "type", "system"),

					resource.TestCheckResourceAttr("data.gmail_label.spam", "id", "SPAM"),
					resource.TestCheckResourceAttr("data.gmail_label.spam", "label_list_visibility", "labelHide"),
					resource.TestCheckResourceAttr("data.gmail_label.spam", "message_list_visibility", "hide"),
					resource.TestCheckResourceAttr("data.gmail_label.spam", "name", "SPAM"),
					resource.TestCheckNoResourceAttr("data.gmail_label.spam", "background_color"),
					resource.TestCheckNoResourceAttr("data.gmail_label.spam", "text_color"),
					resource.TestCheckResourceAttr("data.gmail_label.spam", "type", "system"),

					resource.TestCheckResourceAttr("data.gmail_label.category_forums", "id", "CATEGORY_FORUMS"),
					resource.TestCheckResourceAttr("data.gmail_label.category_forums", "label_list_visibility", "labelHide"),
					resource.TestCheckResourceAttr("data.gmail_label.category_forums", "message_list_visibility", "hide"),
					resource.TestCheckResourceAttr("data.gmail_label.category_forums", "name", "CATEGORY_FORUMS"),
					resource.TestCheckNoResourceAttr("data.gmail_label.category_forums", "background_color"),
					resource.TestCheckNoResourceAttr("data.gmail_label.category_forums", "text_color"),
					resource.TestCheckResourceAttr("data.gmail_label.category_forums", "type", "system"),

					resource.TestCheckResourceAttr("data.gmail_label.category_updates", "id", "CATEGORY_UPDATES"),
					resource.TestCheckResourceAttr("data.gmail_label.category_updates", "label_list_visibility", "labelHide"),
					resource.TestCheckResourceAttr("data.gmail_label.category_updates", "message_list_visibility", "hide"),
					resource.TestCheckResourceAttr("data.gmail_label.category_updates", "name", "CATEGORY_UPDATES"),
					resource.TestCheckNoResourceAttr("data.gmail_label.category_updates", "background_color"),
					resource.TestCheckNoResourceAttr("data.gmail_label.category_updates", "text_color"),
					resource.TestCheckResourceAttr("data.gmail_label.category_updates", "type", "system"),

					resource.TestCheckResourceAttr("data.gmail_label.category_personal", "id", "CATEGORY_PERSONAL"),
					resource.TestCheckResourceAttr("data.gmail_label.category_personal", "label_list_visibility", "labelHide"),
					resource.TestCheckResourceAttr("data.gmail_label.category_personal", "message_list_visibility", "hide"),
					resource.TestCheckResourceAttr("data.gmail_label.category_personal", "name", "CATEGORY_PERSONAL"),
					resource.TestCheckNoResourceAttr("data.gmail_label.category_personal", "background_color"),
					resource.TestCheckNoResourceAttr("data.gmail_label.category_personal", "text_color"),
					resource.TestCheckResourceAttr("data.gmail_label.category_personal", "type", "system"),

					resource.TestCheckResourceAttr("data.gmail_label.category_promotions", "id", "CATEGORY_PROMOTIONS"),
					resource.TestCheckResourceAttr("data.gmail_label.category_promotions", "label_list_visibility", "labelHide"),
					resource.TestCheckResourceAttr("data.gmail_label.category_promotions", "message_list_visibility", "hide"),
					resource.TestCheckResourceAttr("data.gmail_label.category_promotions", "name", "CATEGORY_PROMOTIONS"),
					resource.TestCheckNoResourceAttr("data.gmail_label.category_promotions", "background_color"),
					resource.TestCheckNoResourceAttr("data.gmail_label.category_promotions", "text_color"),
					resource.TestCheckResourceAttr("data.gmail_label.category_promotions", "type", "system"),

					resource.TestCheckResourceAttr("data.gmail_label.category_social", "id", "CATEGORY_SOCIAL"),
					resource.TestCheckResourceAttr("data.gmail_label.category_social", "label_list_visibility", "labelHide"),
					resource.TestCheckResourceAttr("data.gmail_label.category_social", "message_list_visibility", "hide"),
					resource.TestCheckResourceAttr("data.gmail_label.category_social", "name", "CATEGORY_SOCIAL"),
					resource.TestCheckNoResourceAttr("data.gmail_label.category_social", "background_color"),
					resource.TestCheckNoResourceAttr("data.gmail_label.category_social", "text_color"),
					resource.TestCheckResourceAttr("data.gmail_label.category_social", "type", "system"),

					resource.TestCheckResourceAttr("data.gmail_label.starred", "id", "STARRED"),
					resource.TestCheckResourceAttr("data.gmail_label.starred", "label_list_visibility", ""),
					resource.TestCheckResourceAttr("data.gmail_label.starred", "message_list_visibility", ""),
					resource.TestCheckResourceAttr("data.gmail_label.starred", "name", "STARRED"),
					resource.TestCheckNoResourceAttr("data.gmail_label.starred", "background_color"),
					resource.TestCheckNoResourceAttr("data.gmail_label.starred", "text_color"),
					resource.TestCheckResourceAttr("data.gmail_label.starred", "type", "system"),

					resource.TestCheckResourceAttr("data.gmail_label.unread", "id", "UNREAD"),
					resource.TestCheckResourceAttr("data.gmail_label.unread", "label_list_visibility", ""),
					resource.TestCheckResourceAttr("data.gmail_label.unread", "message_list_visibility", ""),
					resource.TestCheckResourceAttr("data.gmail_label.unread", "name", "UNREAD"),
					resource.TestCheckNoResourceAttr("data.gmail_label.unread", "background_color"),
					resource.TestCheckNoResourceAttr("data.gmail_label.unread", "text_color"),
					resource.TestCheckResourceAttr("data.gmail_label.unread", "type", "system"),
				),
			},
		},
	})
}

func TestAccGmailLabelDataSource_userById(t *testing.T) {
	labelName := "terraform-label-" + acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccGmailLabelDataSourceConfigUserById(labelName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("gmail_label.this", "id", "data.gmail_label.this", "id"),
					resource.TestCheckResourceAttrPair("gmail_label.this", "label_list_visibility", "data.gmail_label.this", "label_list_visibility"),
					resource.TestCheckResourceAttrPair("gmail_label.this", "message_list_visibility", "data.gmail_label.this", "message_list_visibility"),
					resource.TestCheckResourceAttrPair("gmail_label.this", "name", "data.gmail_label.this", "name"),
					resource.TestCheckResourceAttrPair("gmail_label.this", "background_color", "data.gmail_label.this", "background_color"),
					resource.TestCheckResourceAttrPair("gmail_label.this", "text_color", "data.gmail_label.this", "text_color"),
					resource.TestCheckResourceAttr("data.gmail_label.this", "type", "user"),
				),
			},
		},
	})
}

func TestAccGmailLabelDataSource_userByName(t *testing.T) {
	labelName := "terraform-label-" + acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccGmailLabelDataSourceConfigUserByName(labelName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("gmail_label.this", "id", "data.gmail_label.this", "id"),
					resource.TestCheckResourceAttrPair("gmail_label.this", "label_list_visibility", "data.gmail_label.this", "label_list_visibility"),
					resource.TestCheckResourceAttrPair("gmail_label.this", "message_list_visibility", "data.gmail_label.this", "message_list_visibility"),
					resource.TestCheckResourceAttrPair("gmail_label.this", "name", "data.gmail_label.this", "name"),
					resource.TestCheckResourceAttrPair("gmail_label.this", "background_color", "data.gmail_label.this", "background_color"),
					resource.TestCheckResourceAttrPair("gmail_label.this", "text_color", "data.gmail_label.this", "text_color"),
					resource.TestCheckResourceAttr("data.gmail_label.this", "type", "user"),
				),
			},
		},
	})
}

const testAccGmailLabelDataSourceConfigSystem = `
data gmail_label chat {
  name = "CHAT"
}

data gmail_label sent {
  name = "SENT"
}

data gmail_label inbox {
  name = "INBOX"
}

data gmail_label important {
  name = "IMPORTANT"
}

data gmail_label trash {
  name = "TRASH"
}

data gmail_label draft {
  name = "DRAFT"
}

data gmail_label spam {
  name = "SPAM"
}

data gmail_label category_forums {
  name = "CATEGORY_FORUMS"
}

data gmail_label category_updates {
  name = "CATEGORY_UPDATES"
}

data gmail_label category_personal {
  name = "CATEGORY_PERSONAL"
}

data gmail_label category_promotions {
  name = "CATEGORY_PROMOTIONS"
}

data gmail_label category_social {
  name = "CATEGORY_SOCIAL"
}

data gmail_label starred {
  name = "STARRED"
}

data gmail_label unread {
  name = "UNREAD"
}
`

func testAccGmailLabelDataSourceConfigUserById(labelName string) string {
	return fmt.Sprintf(`
resource gmail_label this {
  name = "%s"
	label_list_visibility = "labelShowIfUnread"
	message_list_visibility = "hide"
	background_color = "#a46a21"
	text_color = "#f691b2"
}

data gmail_label this {
  id = gmail_label.this.id
}
`, labelName)
}

func testAccGmailLabelDataSourceConfigUserByName(labelName string) string {
	return fmt.Sprintf(`
resource gmail_label this {
  name = "%s"
	label_list_visibility = "labelShowIfUnread"
	message_list_visibility = "hide"
	background_color = "#a46a21"
	text_color = "#f691b2"
}

data gmail_label this {
  name = gmail_label.this.name
}
`, labelName)
}
