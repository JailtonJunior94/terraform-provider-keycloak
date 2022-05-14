package provider

import (
	"context"

	"github.com/jailtonjunior94/tf_keycloak/keycloak"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceKeycloakRealm() *schema.Resource {
	return &schema.Resource{
		Create:        resourceKeycloakRealmCreate,
		ReadContext:   resourceKeycloakRealmRead,
		UpdateContext: resourceKeycloakRealmUpdate,
		Delete:        resourceKeycloakRealmDelete,

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

func resourceKeycloakRealmCreate(data *schema.ResourceData, meta interface{}) error {
	sdk := meta.(*keycloak.KeycloakSDK)

	realmName := data.Get("realm").(string)
	displayName := data.Get("display_name").(string)
	enable := data.Get("enabled").(bool)

	newRealm, err := sdk.CreateRealm(realmName, displayName, enable)
	if err != nil {
		return err
	}

	data.SetId(newRealm.ID)
	return nil
}

func resourceKeycloakRealmRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceKeycloakRealmUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceKeycloakRealmDelete(data *schema.ResourceData, meta interface{}) error {
	sdk := meta.(*keycloak.KeycloakSDK)

	realmName := data.Get("realm").(string)

	err := sdk.DeleteRealm(realmName)
	if err != nil {
		return err
	}

	data.SetId("")
	return nil
}
