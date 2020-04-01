package gmail

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"google.golang.org/api/gmail/v1"
)

func resourceGmailFilter() *schema.Resource {
	criteria := []string{
		"criteria_exclude_chats",
		"criteria_from",
		"criteria_has_attachment",
		"criteria_negated_query",
		"criteria_query",
		"criteria_size_larger",
		"criteria_size_smaller",
		"criteria_subject",
		"criteria_to",
	}
	actions := []string{
		"action_add_label_ids",
		"action_forward",
		"action_remove_label_ids",
	}
	return &schema.Resource{
		Create: resourceGmailFilterCreate,
		Read:   resourceGmailFilterRead,
		Delete: resourceGmailFilterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"criteria_exclude_chats": {
				Type:         schema.TypeBool,
				Optional:     true,
				ForceNew:     true,
				AtLeastOneOf: criteria,
			},
			"criteria_from": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: noLeadingTrailingRepeatedSpaces(),
				AtLeastOneOf: criteria,
			},
			"criteria_has_attachment": {
				Type:         schema.TypeBool,
				Optional:     true,
				ForceNew:     true,
				AtLeastOneOf: criteria,
			},
			"criteria_negated_query": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: noLeadingTrailingRepeatedSpaces(),
				AtLeastOneOf: criteria,
			},
			"criteria_query": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: noLeadingTrailingRepeatedSpaces(),
				AtLeastOneOf: criteria,
			},
			"criteria_size_larger": {
				Type:          schema.TypeInt,
				Optional:      true,
				ForceNew:      true,
				ValidateFunc:  validation.IntAtLeast(0),
				ConflictsWith: []string{"criteria_size_smaller"},
				AtLeastOneOf:  criteria,
			},
			"criteria_size_smaller": {
				Type:          schema.TypeInt,
				Optional:      true,
				ForceNew:      true,
				ValidateFunc:  validation.IntAtLeast(0),
				ConflictsWith: []string{"criteria_size_larger"},
				AtLeastOneOf:  criteria,
			},
			"criteria_subject": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: noLeadingTrailingRepeatedSpaces(),
				AtLeastOneOf: criteria,
			},
			"criteria_to": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: noLeadingTrailingRepeatedSpaces(),
				AtLeastOneOf: criteria,
			},
			"action_add_label_ids": {
				Type:         schema.TypeSet,
				Elem:         &schema.Schema{Type: schema.TypeString},
				Optional:     true,
				ForceNew:     true,
				Set:          schema.HashString,
				AtLeastOneOf: actions,
			},
			"action_forward": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				AtLeastOneOf: actions,
			},
			"action_remove_label_ids": {
				Type:         schema.TypeSet,
				Elem:         &schema.Schema{Type: schema.TypeString},
				Optional:     true,
				ForceNew:     true,
				Set:          schema.HashString,
				AtLeastOneOf: actions,
			},
		},
	}
}

func resourceGmailFilterCreate(data *schema.ResourceData, meta interface{}) error {
	srv := meta.(*gmail.Service)

	filter := &gmail.Filter{
		Action: &gmail.FilterAction{
			AddLabelIds:    expandStringSet(data.Get("action_add_label_ids").(*schema.Set)),
			Forward:        data.Get("action_forward").(string),
			RemoveLabelIds: expandStringSet(data.Get("action_remove_label_ids").(*schema.Set)),
		},
		Criteria: &gmail.FilterCriteria{
			ExcludeChats:  data.Get("criteria_exclude_chats").(bool),
			From:          data.Get("criteria_from").(string),
			HasAttachment: data.Get("criteria_has_attachment").(bool),
			NegatedQuery:  data.Get("criteria_negated_query").(string),
			Query:         data.Get("criteria_query").(string),
			Subject:       data.Get("criteria_subject").(string),
			To:            data.Get("criteria_to").(string),
		},
	}

	if size, ok := data.GetOk("criteria_size_larger"); ok {
		filter.Criteria.Size = int64(size.(int))
		filter.Criteria.SizeComparison = "larger"
	} else if size, ok := data.GetOk("criteria_size_smaller"); ok {
		filter.Criteria.Size = int64(size.(int))
		filter.Criteria.SizeComparison = "smaller"
	}

	ret, err := srv.Users.Settings.Filters.Create(user, filter).Do()
	if err != nil {
		return fmt.Errorf("could not create filter: %v", err)
	}

	data.SetId(ret.Id)

	return resourceGmailFilterRead(data, meta)
}

func resourceGmailFilterRead(data *schema.ResourceData, meta interface{}) error {
	srv := meta.(*gmail.Service)

	filter, err := srv.Users.Settings.Filters.Get(user, data.Id()).Do()
	if err != nil {
		if is404Error(err) {
			log.Printf("[WARN] Removing filter %s from state because it's already gone", data.Id())
			data.SetId("")
			return nil
		}
		return fmt.Errorf("could not fetch filter %s: %v", data.Id(), err)
	}

	data.Set("action_add_label_ids", flattenStringList(filter.Action.AddLabelIds))
	data.Set("action_forward", filter.Action.Forward)
	data.Set("action_remove_label_ids", flattenStringList(filter.Action.RemoveLabelIds))
	data.Set("criteria_exclude_chats", filter.Criteria.ExcludeChats)
	data.Set("criteria_from", filter.Criteria.From)
	data.Set("criteria_has_attachment", filter.Criteria.HasAttachment)
	data.Set("criteria_negated_query", filter.Criteria.NegatedQuery)
	data.Set("criteria_query", filter.Criteria.Query)
	data.Set("criteria_subject", filter.Criteria.Subject)
	data.Set("criteria_to", filter.Criteria.To)

	if filter.Criteria.SizeComparison == "larger" {
		data.Set("criteria_size_larger", filter.Criteria.Size)
	} else if filter.Criteria.SizeComparison == "smaller" {
		data.Set("criteria_size_smaller", filter.Criteria.Size)
	}

	return nil
}

func resourceGmailFilterDelete(data *schema.ResourceData, meta interface{}) error {
	srv := meta.(*gmail.Service)

	if err := srv.Users.Settings.Filters.Delete(user, data.Id()).Do(); err != nil {
		return fmt.Errorf("could not delete filter %s: %v", data.Id(), err)
	}

	return nil
}
