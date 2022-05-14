package provider

import (
	"context"

	"github.com/jailtonjunior94/tf_keycloak/keycloak"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceKeycloakRealm() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceKeycloakRealmCreate,
		Read:          nil,
		Update:        nil,
		Delete:        nil,

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
	client := meta.(*keycloak.KeycloakClient)

	realm := &keycloak.Realm{
		Id:          data.Get("realm").(string),
		Realm:       data.Get("realm").(string),
		DisplayName: data.Get("display_name").(string),
		Enabled:     data.Get("enabled").(bool),
	}

	err := client.NewRealm(ctx, realm)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
