package repository

import (
	"context"

	"github.com/Nerzal/gocloak/v13"

	"github.com/XxThunderBlastxX/thunder-api/domain"
	"github.com/XxThunderBlastxX/thunder-api/internal/gen/keycloakconfig"
)

type keycloakRepository struct {
	AuthURL      string
	Realm        string
	ClientID     string
	ClientSecret string
}

func NewKeycloakRepository(keycloakConfig *keycloakconfig.KeycloakConfig) domain.KeycloakRepository {
	return &keycloakRepository{
		AuthURL:      keycloakConfig.AuthUrl,
		Realm:        keycloakConfig.Realm,
		ClientID:     keycloakConfig.ClientId,
		ClientSecret: keycloakConfig.ClientSecret,
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
