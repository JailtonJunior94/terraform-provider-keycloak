package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jailtonjunior94/keycloak-sdk-go/keycloak"
)

func resourceKeycloakClient() *schema.Resource {
	return &schema.Resource{
		Read:   resourceKeycloakClientRead,
		Create: resourceKeycloakClientCreate,
		Update: resourceKeycloakClientUpdate,
		Delete: resourceKeycloakClientDelete,

		Schema: map[string]*schema.Schema{
			"realm_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_scope": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"base_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"public_client": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"service_accounts_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"client_secret": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceKeycloakClientCreate(data *schema.ResourceData, meta interface{}) error {
	sdk := meta.(*keycloak.KeycloakSDK)

	realmID := data.Get("realm_id").(string)
	clientID := data.Get("client_id").(string)
	name := data.Get("name").(string)
	description := data.Get("description").(string)
	protocol := data.Get("protocol").(string)
	baseURL := data.Get("base_url").(string)
	clientScope := data.Get("client_scope").(string)
	publicClient := data.Get("public_client").(bool)
	serviceAccountsEnabled := data.Get("service_accounts_enabled").(bool)

	new, err := sdk.CreateClient(realmID, clientID, name, description, protocol, baseURL, clientScope, publicClient, serviceAccountsEnabled)
	if err != nil {
		return err
	}

	if !publicClient {
		data.Set("client_secret", new.ClientSecret.Value)
	}

	data.SetId(new.ID)
	return nil
}

func resourceKeycloakClientRead(data *schema.ResourceData, meta interface{}) error {
	sdk := meta.(*keycloak.KeycloakSDK)

	realmID := data.Get("realm_id").(string)

	client, err := sdk.FetchClient(realmID, data.Id())
	if err != nil {
		data.SetId("")
	}

	data.Set("client_id", client.ClientID)
	if !client.PublicClient {
		data.Set("client_secret", client.ClientSecret.Value)
	}

	return nil
}

func resourceKeycloakClientUpdate(data *schema.ResourceData, meta interface{}) error {
	sdk := meta.(*keycloak.KeycloakSDK)

	realmID := data.Get("realm_id").(string)
	clientID := data.Get("client_id").(string)
	name := data.Get("name").(string)
	description := data.Get("description").(string)
	protocol := data.Get("protocol").(string)
	baseURL := data.Get("base_url").(string)
	clientScope := data.Get("client_scope").(string)
	publicClient := data.Get("public_client").(bool)
	serviceAccountsEnabled := data.Get("service_accounts_enabled").(bool)

	update, err := sdk.UpdateClient(realmID, data.Id(), clientID, name, description, protocol, baseURL, clientScope, publicClient, serviceAccountsEnabled)
	if err != nil {
		data.SetId("")
	}

	if !publicClient {
		data.Set("client_secret", update.ClientSecret.Value)
	}

	data.SetId(update.ID)
	return nil
}

func resourceKeycloakClientDelete(data *schema.ResourceData, meta interface{}) error {
	sdk := meta.(*keycloak.KeycloakSDK)

	realmID := data.Get("realm_id").(string)

	err := sdk.DeleteClient(realmID, data.Id())
	if err != nil {
		return err
	}

	data.SetId("")
	return nil
}
