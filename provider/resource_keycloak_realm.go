package provider

import (
	"github.com/jailtonjunior94/tf_keycloak/keycloak"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceKeycloakRealm() *schema.Resource {
	return &schema.Resource{
		Read:   resourceKeycloakRealmRead,
		Create: resourceKeycloakRealmCreate,
		Update: resourceKeycloakRealmUpdate,
		Delete: resourceKeycloakRealmDelete,

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

	data.SetId(newRealm.Realm)
	return nil
}

func resourceKeycloakRealmRead(data *schema.ResourceData, meta interface{}) error {
	sdk := meta.(*keycloak.KeycloakSDK)

	realm, err := sdk.FetchRealm(data.Id())
	if err != nil {
		data.SetId("")
	}

	data.Set("realm", realm.Realm)
	return nil
}

func resourceKeycloakRealmUpdate(data *schema.ResourceData, meta interface{}) error {
	sdk := meta.(*keycloak.KeycloakSDK)

	realmName := data.Get("realm").(string)
	displayName := data.Get("display_name").(string)
	enable := data.Get("enabled").(bool)

	updateRealm, err := sdk.UpdateRealm(data.Id(), realmName, displayName, enable)
	if err != nil {
		data.SetId("")
	}

	data.SetId(updateRealm.Realm)
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
