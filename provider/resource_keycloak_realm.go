package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	keycloak "github.com/jailtonjunior94/keycloak-sdk-go"
)

func resourceKeycloakRealm() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceKeycloakRealmCreate,
		ReadContext:   resourceKeycloakRealmRead,
		UpdateContext: resourceKeycloakRealmUpdate,
		DeleteContext: resourceKeycloakRealmDelete,

		Schema: map[string]*schema.Schema{
			"realm": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceKeycloakRealmCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	realm := data.Get("realm").(string)
	displayName := data.Get("display_name").(string)
	enable := data.Get("enabled").(bool)

	sdk := meta.(*keycloak.KeycloakSDK)
	_, _ = sdk.CreateRealm(realm, displayName, enable)

	return nil
}

func resourceKeycloakRealmRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceKeycloakRealmUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceKeycloakRealmDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
