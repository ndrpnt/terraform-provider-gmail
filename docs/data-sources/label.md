# gmail_label Data Source

Get the ID (and other info) of a [Gmail label][1].

## Example Usage

```hcl
data gmail_label inbox {
  name = "INBOX"
}

data gmail_label trash {
  name = "TRASH"
}

resource gmail_filter spam {
  criteria_from = "spammer@gmail.com"
  action_add_label_ids = [data.gmail_label.trash.id]
  action_remove_label_ids = [data.gmail_label.inbox.id]
}

output "inbox_label_type" {
  value = data.gmail_label.inbox.type
}
```

## Argument Reference

* `id` - (Optional) The immutable ID of the label.
* `name` - (Optional) The display name of the label.

## Attribute Reference

* `label_list_visibility` - The visibility of the label in the label list in the
  Gmail web interface.
* `message_list_visibility` - The visibility of messages with this label in the
  message list in the Gmail web interface.
* `background_color` - The background color of the label.
* `text_color` - The text color of the label.
* `type` - The owner type for the label.

[1]: https://developers.google.com/gmail/api/v1/reference/users/labels
