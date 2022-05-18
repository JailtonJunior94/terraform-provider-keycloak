package keycloak

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Realm struct {
	ID          string `json:"id"`
	Realm       string `json:"realm"`
	DisplayName string `json:"displayName"`
	Enabled     bool   `json:"enabled"`
}

func (k *KeycloakSDK) FetchRealm(realm string) (*Realm, error) {
	uri := fmt.Sprintf("/admin/realms/%s", realm)
	response, err := request("GET", k.BaseURL, uri, "application/json", k.Session.AccessToken, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var r Realm
	err = json.Unmarshal(response, &r)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &r, nil
}

func (k *KeycloakSDK) CreateRealm(realm, displayName string, enable bool) (*Realm, error) {
	newRealm := &Realm{
		ID:          realm,
		Realm:       realm,
		DisplayName: displayName,
		Enabled:     enable,
	}

	json, err := json.Marshal(&newRealm)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	payload := bytes.NewBuffer(json)
	_, err = request("POST", k.BaseURL, "/admin/realms", "application/json", k.Session.AccessToken, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return k.FetchRealm(realm)
}

func (k *KeycloakSDK) UpdateRealm(id, realm, displayName string, enable bool) (*Realm, error) {
	updateRealm := &Realm{
		ID:          realm,
		Realm:       realm,
		DisplayName: displayName,
		Enabled:     enable,
	}

	json, err := json.Marshal(&updateRealm)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	uri := fmt.Sprintf("/admin/realms/%s", id)
	payload := bytes.NewBuffer(json)
	_, err = request("PUT", k.BaseURL, uri, "application/json", k.Session.AccessToken, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return k.FetchRealm(realm)
}

func (k *KeycloakSDK) DeleteRealm(realm string) error {
	uri := fmt.Sprintf("/admin/realms/%s", realm)
	_, err := request("DELETE", k.BaseURL, uri, "application/json", k.Session.AccessToken, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
