provider gmail {
  credentials_file = "credentials.json"
  token_file       = "token.json"
}

data gmail_label inbox {
  name = "INBOX"
}

data gmail_label trash {
  name = "TRASH"
}

data gmail_label important {
  name = "IMPORTANT"
}

resource gmail_label work {
  name                    = "Work"
  label_list_visibility   = "labelShowIfUnread"
  message_list_visibility = "show"
  background_color        = "#a46a21"
  text_color              = "#f691b2"
}

resource gmail_filter ignore {
  criteria_from           = "spammer@gmail.com"
  action_add_label_ids    = [data.gmail_label.trash.id]
}

resource gmail_filter work {
  criteria_from           = "foobar"
  criteria_has_attachment = true
  criteria_query          = "rfc822msgid: is:unread"
  criteria_to             = "bazqux"
  action_add_label_ids    = [gmail_label.work.id, data.gmail_label.important.id]
  action_remove_label_ids = [data.gmail_label.inbox.id]
}

output "inbox_label_type" {
  value = data.gmail_label.inbox.type
}
