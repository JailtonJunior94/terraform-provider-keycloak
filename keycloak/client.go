package keycloak

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jailtonjunior94/tf_keycloak/shared"

	uuid "github.com/satori/go.uuid"
)

type Client struct {
	ID                           string        `json:"id,omitempty"`
	ClientID                     string        `json:"clientId"`
	Name                         string        `json:"name"`
	Description                  string        `json:"description"`
	Protocol                     string        `json:"protocol"`
	PublicClient                 bool          `json:"publicClient"`
	ServiceAccountsEnabled       bool          `json:"serviceAccountsEnabled"`
	Enabled                      bool          `json:"enabled"`
	AuthorizationServicesEnabled bool          `json:"authorizationServicesEnabled"`
	DefaultClientScopes          []string      `json:"defaultClientScopes"`
	BaseURL                      string        `json:"baseUrl"`
	RedirectUris                 []string      `json:"redirectUris"`
	WebOrigins                   []string      `json:"webOrigins"`
	ClientSecret                 *ClientSecret `json:"clientSecret,omitempty"`
}

func (k *KeycloakSDK) Clients(realm string) ([]*Client, error) {
	uri := fmt.Sprintf("/admin/realms/%s/clients", realm)
	response, err := request("GET", k.BaseURL, uri, "application/json", k.Session.AccessToken, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var c []*Client
	err = json.Unmarshal(response, &c)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return c, nil
}

func (k *KeycloakSDK) FetchClient(realm, id string) (*Client, error) {
	uri := fmt.Sprintf("/admin/realms/%s/clients/%s", realm, id)
	response, err := request("GET", k.BaseURL, uri, "application/json", k.Session.AccessToken, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var c Client
	err = json.Unmarshal(response, &c)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	secret, err := k.FetchClientSecret(realm, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	c.ClientSecret = secret
	return &c, nil
}

func (k *KeycloakSDK) CreateClient(realm, clientID, name, description, protocol, baseURL, clientScope string, publicClient, serviceAccountsEnabled bool) (*Client, error) {
	clients, _ := k.Clients(realm)
	filterByName := shared.Filter(clients, func(v *Client) bool {
		return v.ClientID == clientID
	})

	if len(filterByName) > 0 {
		return filterByName[0], nil
	}

	new := &Client{
		ID:                           uuid.NewV4().String(),
		ClientID:                     clientID,
		Name:                         name,
		Description:                  description,
		Protocol:                     protocol,
		PublicClient:                 publicClient,
		ServiceAccountsEnabled:       serviceAccountsEnabled,
		Enabled:                      true,
		AuthorizationServicesEnabled: serviceAccountsEnabled,
		BaseURL:                      baseURL,
		DefaultClientScopes:          []string{"email", "profile", "roles", clientScope},
		RedirectUris:                 []string{baseURL},
		WebOrigins:                   []string{baseURL},
	}

	json, err := json.Marshal(&new)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	uri := fmt.Sprintf("/admin/realms/%s/clients", realm)
	payload := bytes.NewBuffer(json)
	_, err = request("POST", k.BaseURL, uri, "application/json", k.Session.AccessToken, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return k.FetchClient(realm, new.ID)
}

func (k *KeycloakSDK) UpdateClient(realm, id, clientID, name, description, protocol, baseURL, clientScope string, publicClient, serviceAccountsEnabled bool) (*Client, error) {
	update := &Client{
		ClientID:                     clientID,
		Name:                         name,
		Description:                  description,
		Protocol:                     protocol,
		PublicClient:                 publicClient,
		ServiceAccountsEnabled:       serviceAccountsEnabled,
		Enabled:                      true,
		AuthorizationServicesEnabled: serviceAccountsEnabled,
		BaseURL:                      baseURL,
		DefaultClientScopes:          []string{"email", "profile", "roles", clientScope},
		RedirectUris:                 []string{baseURL},
		WebOrigins:                   []string{baseURL},
	}

	json, err := json.Marshal(&update)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	uri := fmt.Sprintf("/admin/realms/%s/clients/%s", realm, id)
	payload := bytes.NewBuffer(json)
	_, err = request("PUT", k.BaseURL, uri, "application/json", k.Session.AccessToken, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return k.FetchClient(realm, id)
}

func (k *KeycloakSDK) DeleteClient(realm, id string) error {
	uri := fmt.Sprintf("/admin/realms/%s/clients/%s", realm, id)
	_, err := request("DELETE", k.BaseURL, uri, "application/json", k.Session.AccessToken, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
