package gmail

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"google.golang.org/api/gmail/v1"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func gmailService() (*gmail.Service, error) {
	provider := Provider().(*schema.Provider)

	if err := provider.Configure(terraform.NewResourceConfigRaw(nil)); err != nil {
		return nil, fmt.Errorf("failed to configure provider: %v", err)
	}
	return provider.Meta().(*gmail.Service), nil
}
