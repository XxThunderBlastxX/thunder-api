package repository

import (
	"context"

	"github.com/Nerzal/gocloak/v13"

	"github.com/XxThunderBlastxX/thunder-api/domain"
)

type keycloakRepository struct {
	AuthURL      string
	Realm        string
	ClientID     string
	ClientSecret string
}

func NewKeycloakRepository(authURL string, realm string, clientID string, clientSecret string) domain.KeycloakRepository {
	return &keycloakRepository{
		AuthURL:      authURL,
		Realm:        realm,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

func (kc *keycloakRepository) IntrospectToken(ctx context.Context, token string) (*gocloak.IntroSpectTokenResult, error) {
	client := gocloak.NewClient(kc.AuthURL)

	tokenResult, err := client.RetrospectToken(ctx, token, kc.ClientID, kc.ClientSecret, kc.Realm)
	if err != nil {
		return nil, err
	}

	return tokenResult, nil
}
