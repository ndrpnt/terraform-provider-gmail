package gmail

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/api/gmail/v1"
)

func dataSourceGmailLabel() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGmailLabelRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"id", "name"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"id", "name"},
			},
			"label_list_visibility": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"message_list_visibility": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"background_color": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"text_color": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGmailLabelRead(data *schema.ResourceData, meta interface{}) error {
	srv := meta.(*gmail.Service)

	if id, ok := data.GetOk("id"); !ok {
		labels, err := srv.Users.Labels.List(user).Do()
		if err != nil {
			return fmt.Errorf("could not fetch labels: %v", err)
		}

		name := data.Get("name").(string)
		for _, label := range labels.Labels {
			if label.Name == name {
				data.SetId(label.Id)
				break
			}
		}
	} else {
		data.SetId(id.(string))
	}

	label, err := srv.Users.Labels.Get(user, data.Id()).Do()
	if err != nil {
		if is404Error(err) {
			return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
		}
		return fmt.Errorf("could not fetch label %s: %v", data.Id(), err)
	}

	data.Set("name", label.Name)
	data.Set("label_list_visibility", label.LabelListVisibility)
	data.Set("message_list_visibility", label.MessageListVisibility)
	if label.Color != nil {
		data.Set("background_color", label.Color.BackgroundColor)
		data.Set("text_color", label.Color.TextColor)
	} else {
		data.Set("background_color", defaultLabelBackgroundColor)
		data.Set("text_color", defaultLabelTextColor)
	}
	data.Set("type", label.Type)

	return nil
}
