package gmail

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"google.golang.org/api/gmail/v1"
)

const (
	defaultLabelBackgroundColor = "#999999"
	defaultLabelTextColor       = "#f3f3f3"
)

func resourceGmailLabel() *schema.Resource {
	validColors := []string{
		"#000000",
		"#434343",
		"#666666",
		"#999999",
		"#cccccc",
		"#efefef",
		"#f3f3f3",
		"#ffffff",
		"#fb4c2f",
		"#ffad47",
		"#fad165",
		"#16a766",
		"#43d692",
		"#4a86e8",
		"#a479e2",
		"#f691b3",
		"#f6c5be",
		"#ffe6c7",
		"#fef1d1",
		"#b9e4d0",
		"#c6f3de",
		"#c9daf8",
		"#e4d7f5",
		"#fcdee8",
		"#efa093",
		"#ffd6a2",
		"#fce8b3",
		"#89d3b2",
		"#a0eac9",
		"#a4c2f4",
		"#d0bcf1",
		"#fbc8d9",
		"#e66550",
		"#ffbc6b",
		"#fcda83",
		"#44b984",
		"#68dfa9",
		"#6d9eeb",
		"#b694e8",
		"#f7a7c0",
		"#cc3a21",
		"#eaa041",
		"#f2c960",
		"#149e60",
		"#3dc789",
		"#3c78d8",
		"#8e63ce",
		"#e07798",
		"#ac2b16",
		"#cf8933",
		"#d5ae49",
		"#0b804b",
		"#2a9c68",
		"#285bac",
		"#653e9b",
		"#b65775",
		"#822111",
		"#a46a21",
		"#aa8831",
		"#076239",
		"#1a764d",
		"#1c4587",
		"#41236d",
		"#83334c",
		"#464646",
		"#e7e7e7",
		"#0d3472",
		"#b6cff5",
		"#0d3b44",
		"#98d7e4",
		"#3d188e",
		"#e3d7ff",
		"#711a36",
		"#fbd3e0",
		"#8a1c0a",
		"#f2b2a8",
		"#7a2e0b",
		"#ffc8af",
		"#7a4706",
		"#ffdeb5",
		"#594c05",
		"#fbe983",
		"#684e07",
		"#fdedc1",
		"#0b4f30",
		"#b3efd3",
		"#04502e",
		"#a2dcc1",
		"#c2c2c2",
		"#4986e7",
		"#2da2bb",
		"#b99aff",
		"#994a64",
		"#f691b2",
		"#ff7537",
		"#ffad46",
		"#662e37",
		"#ebdbde",
		"#cca6ac",
		"#094228",
		"#42d692",
		"#16a765",
	}

	return &schema.Resource{
		Create: resourceGmailLabelCreate,
		Read:   resourceGmailLabelRead,
		Update: resourceGmailLabelUpdate,
		Delete: resourceGmailLabelDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: noLeadingTrailingRepeatedSpaces(),
			},
			"label_list_visibility": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "labelShow",
				ValidateFunc: validation.StringInSlice([]string{"labelHide", "labelShow", "labelShowIfUnread"}, false),
			},
			"message_list_visibility": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "show",
				ValidateFunc: validation.StringInSlice([]string{"hide", "show"}, false),
			},
			"background_color": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      defaultLabelBackgroundColor,
				ValidateFunc: validation.StringInSlice(validColors, false),
			},
			"text_color": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      defaultLabelTextColor,
				ValidateFunc: validation.StringInSlice(validColors, false),
			},
		},
	}
}

func resourceGmailLabelCreate(data *schema.ResourceData, meta interface{}) error {
	srv := meta.(*gmail.Service)
	label := &gmail.Label{
		Color: &gmail.LabelColor{
			BackgroundColor: data.Get("background_color").(string),
			TextColor:       data.Get("text_color").(string),
		},
		LabelListVisibility:   data.Get("label_list_visibility").(string),
		MessageListVisibility: data.Get("message_list_visibility").(string),
		Name:                  data.Get("name").(string),
	}

	ret, err := srv.Users.Labels.Create(user, label).Do()
	if err != nil {
		return fmt.Errorf("could not create label %s: %v", label.Name, err)
	}

	data.SetId(ret.Id)

	return resourceGmailLabelRead(data, meta)
}

func resourceGmailLabelRead(data *schema.ResourceData, meta interface{}) error {
	srv := meta.(*gmail.Service)

	label, err := srv.Users.Labels.Get(user, data.Id()).Do()
	if err != nil {
		if is404Error(err) {
			log.Printf("[WARN] Removing label %s from state because it's already gone", data.Id())
			data.SetId("")
			return nil
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

	return nil
}

func resourceGmailLabelUpdate(data *schema.ResourceData, meta interface{}) error {
	srv := meta.(*gmail.Service)
	label := &gmail.Label{
		Color: &gmail.LabelColor{
			BackgroundColor: data.Get("background_color").(string),
			TextColor:       data.Get("text_color").(string),
		},
		LabelListVisibility:   data.Get("label_list_visibility").(string),
		MessageListVisibility: data.Get("message_list_visibility").(string),
		Name:                  data.Get("name").(string),
	}

	_, err := srv.Users.Labels.Update(user, data.Id(), label).Do()
	if err != nil {
		return fmt.Errorf("could not update label %s: %v", label.Name, err)
	}

	return resourceGmailLabelRead(data, meta)
}

func resourceGmailLabelDelete(data *schema.ResourceData, meta interface{}) error {
	srv := meta.(*gmail.Service)

	if err := srv.Users.Labels.Delete(user, data.Id()).Do(); err != nil {
		return fmt.Errorf("could not delete label %s: %v", data.Id(), err)
	}

	return nil
}
