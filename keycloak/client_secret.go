package keycloak

import (
	"encoding/json"
	"fmt"
	"log"
)

type ClientSecret struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (k *KeycloakSDK) FetchClientSecret(realm, id string) (*ClientSecret, error) {
	uri := fmt.Sprintf("/admin/realms/%s/clients/%s/client-secret", realm, id)
	response, err := request("GET", k.BaseURL, uri, "application/json", k.Session.AccessToken, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var c ClientSecret
	err = json.Unmarshal(response, &c)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &c, nil
}
