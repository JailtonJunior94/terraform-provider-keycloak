package provider

import (
	"github.com/jailtonjunior94/tf_keycloak/keycloak"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	Protocols = []string{"openid-connect", "saml"}
)

func resourceKeycloakClientScope() *schema.Resource {
	return &schema.Resource{
		Read:   resourceKeycloakClientScopeRead,
		Create: resourceKeycloakClientScopeCreate,
		Update: resourceKeycloakClientScopeUpdate,
		Delete: resourceKeycloakClientScopeDelete,

		Schema: map[string]*schema.Schema{
			"realm_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "openid-connect",
			},
		},
	}
}

func resourceKeycloakClientScopeCreate(data *schema.ResourceData, meta interface{}) error {
	sdk := meta.(*keycloak.KeycloakSDK)

	realmID := data.Get("realm_id").(string)
	name := data.Get("name").(string)
	description := data.Get("description").(string)
	protocol := data.Get("protocol").(string)

	new, err := sdk.CreateClientScope(realmID, name, description, protocol)
	if err != nil {
		return err
	}

	data.SetId(new.ID)
	return nil
}

func resourceKeycloakClientScopeRead(data *schema.ResourceData, meta interface{}) error {
	sdk := meta.(*keycloak.KeycloakSDK)

	realmID := data.Get("realm_id").(string)

	clienteScope, err := sdk.FetchClientScope(realmID, data.Id())
	if err != nil {
		data.SetId("")
	}

	data.Set("name", clienteScope.Name)
	return nil
}

func resourceKeycloakClientScopeUpdate(data *schema.ResourceData, meta interface{}) error {
	sdk := meta.(*keycloak.KeycloakSDK)

	realmID := data.Get("realm_id").(string)
	name := data.Get("name").(string)
	description := data.Get("description").(string)
	protocol := data.Get("protocol").(string)

	update, err := sdk.UpdateClientScope(realmID, data.Id(), name, description, protocol)
	if err != nil {
		data.SetId("")
	}

	data.SetId(update.ID)
	return nil
}

func resourceKeycloakClientScopeDelete(data *schema.ResourceData, meta interface{}) error {
	sdk := meta.(*keycloak.KeycloakSDK)

	realmID := data.Get("realm_id").(string)

	err := sdk.DeleteClientScope(realmID, data.Id())
	if err != nil {
		return err
	}

	data.SetId("")
	return nil
}
