package keycloak

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jailtonjunior94/tf_keycloak/shared"

	uuid "github.com/satori/go.uuid"
)

type ClientScope struct {
	ID          string                 `json:"id,omitempty"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Protocol    string                 `json:"protocol"`
	Attributes  *ClientScopeAttributes `json:"attributes"`
}

type ClientScopeAttributes struct {
	ConsentScreenText      string `json:"consent.screen.text"`
	DisplayOnConsentScreen string `json:"display.on.consent.screen"`
}

func (k *KeycloakSDK) ClientScope(realm string) ([]*ClientScope, error) {
	uri := fmt.Sprintf("/admin/realms/%s/client-scopes", realm)
	response, err := request("GET", k.BaseURL, uri, "application/json", k.Session.AccessToken, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var c []*ClientScope
	err = json.Unmarshal(response, &c)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return c, nil
}

func (k *KeycloakSDK) FetchClientScope(realm, id string) (*ClientScope, error) {
	uri := fmt.Sprintf("/admin/realms/%s/client-scopes/%s", realm, id)
	response, err := request("GET", k.BaseURL, uri, "application/json", k.Session.AccessToken, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var c ClientScope
	err = json.Unmarshal(response, &c)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &c, nil
}

func (k *KeycloakSDK) CreateClientScope(realm, name, description, protocol string) (*ClientScope, error) {
	scopes, _ := k.ClientScope(realm)
	filterByName := shared.Filter(scopes, func(v *ClientScope) bool {
		return v.Name == name
	})

	if len(filterByName) > 0 {
		return filterByName[0], nil
	}

	new := &ClientScope{
		ID:          uuid.NewV4().String(),
		Name:        name,
		Description: description,
		Protocol:    protocol,
		Attributes: &ClientScopeAttributes{
			ConsentScreenText:      "${offlineAccessScopeConsentText}",
			DisplayOnConsentScreen: "true",
		},
	}

	json, err := json.Marshal(&new)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	uri := fmt.Sprintf("/admin/realms/%s/client-scopes", realm)
	payload := bytes.NewBuffer(json)
	_, err = request("POST", k.BaseURL, uri, "application/json", k.Session.AccessToken, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return k.FetchClientScope(realm, new.ID)
}

func (k *KeycloakSDK) UpdateClientScope(realm, id, name, description, protocol string) (*ClientScope, error) {
	update := &ClientScope{
		Name:        name,
		Description: description,
		Protocol:    protocol,
		Attributes: &ClientScopeAttributes{
			ConsentScreenText:      "${offlineAccessScopeConsentText}",
			DisplayOnConsentScreen: "true",
		},
	}

	json, err := json.Marshal(&update)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	uri := fmt.Sprintf("/admin/realms/%s/client-scopes/%s", realm, id)
	payload := bytes.NewBuffer(json)
	_, err = request("PUT", k.BaseURL, uri, "application/json", k.Session.AccessToken, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return k.FetchClientScope(realm, id)
}

func (k *KeycloakSDK) DeleteClientScope(realm, id string) error {
	uri := fmt.Sprintf("/admin/realms/%s/client-scopes/%s", realm, id)
	_, err := request("DELETE", k.BaseURL, uri, "application/json", k.Session.AccessToken, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
