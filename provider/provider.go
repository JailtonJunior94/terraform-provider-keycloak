package provider

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jailtonjunior94/keycloak-sdk-go/keycloak"
)

func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Required:    true,
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("KEYCLOAK_URL", nil),
			},
			"base_path": {
				Optional: true,
				Type:     schema.TypeString,
				Default:  "/auth",
			},
			"username": {
				Required:    true,
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("KEYCLOAK_USERNAME", nil),
			},
			"password": {
				Required:    true,
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("KEYCLOAK_PASSWORD", nil),
			},
			"realm": {
				Optional:    true,
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("KEYCLOAK_REALM", "master"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"keycloak_realm":        resourceKeycloakRealm(),
			"keycloak_client_scope": resourceKeycloakClientScope(),
			"keycloak_client":       resourceKeycloakClient(),
		},
	}

	provider.ConfigureFunc = func(data *schema.ResourceData) (interface{}, error) {
		url := data.Get("url").(string)
		basePath := data.Get("base_path").(string)
		username := data.Get("username").(string)
		password := data.Get("password").(string)

		sdk, err := keycloak.NewKeycloakSDK(fmt.Sprintf("%s%s", url, basePath), username, password)
		if err != nil {
			return nil, errors.New("error initializing keycloak provider")
		}

		return sdk, nil
	}

	return provider
}
