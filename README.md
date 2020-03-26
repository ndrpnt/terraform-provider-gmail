# Terraform Provider for Gmail

## Getting started

[Install](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins)
the plugin.

Follow the [quickstart](https://developers.google.com/gmail/api/quickstart/go)
to obtain a `credentials.json` and `token.json` file. Change the scope to
`gmail.GmailSettingsBasicScope, gmail.GmailLabelsScope`.

Other quickstarts are fine too (except the _Browser_ one), as long as the token
is retrieved in Json format.
