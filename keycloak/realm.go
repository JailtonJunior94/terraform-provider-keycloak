package keycloak

import (
	"context"
)

type Realm struct {
	Id          string `json:"id,omitempty"`
	Realm       string `json:"realm"`
	Enabled     bool   `json:"enabled"`
	DisplayName string `json:"displayName"`
}

func (keycloakClient *KeycloakClient) NewRealm(ctx context.Context, realm *Realm) error {
	// _, _, err := keycloakClient.post(ctx, "/realms", realm)

	return nil
}
