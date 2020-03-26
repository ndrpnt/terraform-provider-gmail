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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"message_list_visibility": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"label_list_visibility": {
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

	labels, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		return err
	}

	name := data.Get("name").(string)
	for _, label := range labels.Labels {
		if label.Name == name {
			data.SetId(label.Id)
			data.Set("message_list_visibility", label.MessageListVisibility)
			data.Set("label_list_visibility", label.LabelListVisibility)
			data.Set("type", label.Type)

			return nil
		}
	}

	return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
}
