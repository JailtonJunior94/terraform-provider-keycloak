package keycloak

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-version"
)

type KeycloakClient struct {
	baseUrl           string
	realm             string
	clientCredentials *ClientCredentials
	httpClient        *http.Client
	initialLogin      bool
	userAgent         string
	version           *version.Version
	additionalHeaders map[string]string
	debug             bool
}

type ClientCredentials struct {
	ClientId     string
	ClientSecret string
	Username     string
	Password     string
	GrantType    string
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

const (
	apiUrl   = "/admin"
	tokenUrl = "%s/realms/%s/protocol/openid-connect/token"
)

func NewKeycloak(ctx context.Context, url, basePath, realm, username, password string) (*KeycloakClient, error) {

	return nil, nil
}
