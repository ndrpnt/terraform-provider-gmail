# gmail_label Resource

Manage a [Gmail label][1].

## Example Usage

```hcl
resource gmail_label example {
  name = "Example"
  label_list_visibility = "labelShowIfUnread"
  message_list_visibility = "show"
  background_color = "#a46a21"
  text_color = "#f691b2"
}
```

## Argument Reference

* `name` - (Required) The display name of the label.
* `label_list_visibility` - (Optional) The visibility of the label in the label
  list in the Gmail web interface.
* `message_list_visibility` - (Optional) The visibility of messages with this
  label in the message list in the Gmail web interface.
* `background_color` - (Optional) The background color of the label.
* `text_color` - (Optional) The text color of the label.

## Attribute Reference

* `id` - The immutable ID of the label.

[1]: https://developers.google.com/gmail/api/v1/reference/users/labels
