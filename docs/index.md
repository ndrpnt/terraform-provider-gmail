# Gmail Provider

The Gmail provider is used to interact with the [Gmail API][1]. The provider
needs to be configured with the proper credentials before it can be used.

## Example Usage

```hcl
# Configure the Gmail Provider
provider gmail {
  credentials_file = "credentials.json"
  token_file = "token.json"
}

# Create a label
resource gmail_label example {
  name = "Example"
}
```

## Authentication

Follow the [quickstart][2] to obtain a `credentials.json` and `token.json`
file. Change the scopes to
`gmail.GmailSettingsBasicScope, gmail.GmailLabelsScope`.

## Argument Reference

* `credentials_file` - (Required) The path to a `credentials.json` file. Can
  alternatively be provided via the `GMAIL_CREDENTIALS_FILE` environment
  variable.
* `token_file` - (Required) The path to a `token.json` file. Can alternatively
  be provided via the `GMAIL_TOKEN_FILE` environment variable.

[1]: https://developers.google.com/gmail/api
[2]: https://developers.google.com/gmail/api/quickstart/go
