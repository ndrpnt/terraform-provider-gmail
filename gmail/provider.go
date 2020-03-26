package gmail

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

const user = "me"

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"gmail_label": dataSourceGmailLabel(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"gmail_label": resourceGmailLabel(),
		},
		Schema: map[string]*schema.Schema{
			"credentials_file": {
				Required:    true,
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("GMAIL_CREDENTIALS_FILE", nil),
			},
			"token_file": {
				Required:    true,
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("GMAIL_TOKEN_FILE", nil),
			},
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	credentialsFile := data.Get("credentials_file").(string)
	config, err := credentialsFromFile(credentialsFile)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials_file: %v", err)
	}

	tokenFile := data.Get("token_file").(string)
	token, err := tokenFromFile(tokenFile)
	if err != nil {
		return nil, fmt.Errorf("invalid token_file: %v", err)
	}

	return gmail.NewService(
		context.Background(),
		option.WithTokenSource(config.TokenSource(context.Background(), token)),
		option.WithScopes(gmail.GmailSettingsBasicScope, gmail.GmailLabelsScope),
	)
}

func credentialsFromFile(credentialsFile string) (*oauth2.Config, error) {
	credentials, err := ioutil.ReadFile(credentialsFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read %s: %v", credentialsFile, err)
	}

	config, err := google.ConfigFromJSON(credentials, gmail.GmailSettingsBasicScope, gmail.GmailLabelsScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse %s: %v", credentialsFile, err)
	}

	return config, nil
}

func tokenFromFile(tokenFile string) (*oauth2.Token, error) {
	file, err := os.Open(tokenFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read %s: %v", tokenFile, err)
	}
	defer file.Close()

	token := &oauth2.Token{}
	err = json.NewDecoder(file).Decode(token)
	if err != nil {
		return nil, fmt.Errorf("unable to parse %s: %v", tokenFile, err)
	}

	return token, err
}
