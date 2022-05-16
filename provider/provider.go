package provider

import (
	"context"
	"fmt"

	"github.com/jailtonjunior94/tf_keycloak/keycloak"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
		},
	}

	provider.ConfigureContextFunc = func(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
		url := data.Get("url").(string)
		basePath := data.Get("base_path").(string)
		username := data.Get("username").(string)
		password := data.Get("password").(string)

		var diags diag.Diagnostics

		sdk, err := keycloak.NewKeycloakSDK(ctx, fmt.Sprintf("%s%s", url, basePath), username, password)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "error initializing keycloak provider",
				Detail:   err.Error(),
			})
		}

		return sdk, diags
	}

	return provider
}
