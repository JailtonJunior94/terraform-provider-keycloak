package keycloak

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type AuthRequest struct {
	ClientID  string
	Username  string
	Password  string
	GrantType string
}

type AuthResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

type KeycloakSDK struct {
	BaseURL  string
	Username string
	Password string
	Context  context.Context
	Session  *AuthResponse
}

var httpClient = &http.Client{Timeout: 30 * time.Second}

func NewKeycloakSDK(ctx context.Context, baseURL, username, password string) (*KeycloakSDK, error) {
	session, err := auth(ctx, baseURL, username, password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	keycloakSDK := &KeycloakSDK{
		BaseURL:  baseURL,
		Username: username,
		Password: password,
		Context:  ctx,
		Session:  session,
	}
	return keycloakSDK, nil
}

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

func auth(ctx context.Context, baseURL, username, password string) (*AuthResponse, error) {
	authReq := &AuthRequest{
		Username:  username,
		Password:  password,
		ClientID:  "admin-cli",
		GrantType: "password",
	}

	data := url.Values{}
	data.Set("client_id", authReq.ClientID)
	data.Set("username", authReq.Username)
	data.Set("password", authReq.Password)
	data.Set("grant_type", authReq.GrantType)

	uri := fmt.Sprintf("%s/realms/master/protocol/openid-connect/token", baseURL)
	req, err := http.NewRequestWithContext(ctx, "POST", uri, strings.NewReader(data.Encode()))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := httpClient.Do(req)
	statusCode := resp.StatusCode
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if statusCode < 200 || statusCode > 299 {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("[ERROR] [%s]\n", err))
	}

	var authResponse AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		log.Printf("[ERROR] [%s]\n", err)
		return nil, err
	}

	return &authResponse, nil
}

func request(method, baseURI, uri, contentType, token string, payload io.Reader) ([]byte, error) {
	url := fmt.Sprintf("%s%s", baseURI, uri)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("%s%s", "Bearer ", token))
	}

	resp, err := httpClient.Do(req)
	statusCode := resp.StatusCode
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	if statusCode < 200 || statusCode > 299 {
		return nil, errors.New(fmt.Sprintf("[ERROR] [StatusCode] [%d]", statusCode))
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	return bytes, err
}
