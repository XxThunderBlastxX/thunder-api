package auth

import (
	"context"

	"github.com/Nerzal/gocloak/v13"

	"github.com/XxThunderBlastxX/thunder-api/internal/gen/keycloakconfig"
)

type Auth interface {
	IntrospectToken(ctx context.Context, token string) (*gocloak.IntroSpectTokenResult, error)
}

type auth struct {
	AuthURL      string
	Realm        string
	ClientID     string
	ClientSecret string
}

func NewAuth(keycloakConfig *keycloakconfig.KeycloakConfig) Auth {
	return &auth{
		AuthURL:      keycloakConfig.AuthUrl,
		Realm:        keycloakConfig.Realm,
		ClientID:     keycloakConfig.ClientId,
		ClientSecret: keycloakConfig.ClientSecret,
	}
}

func (a *auth) IntrospectToken(ctx context.Context, token string) (*gocloak.IntroSpectTokenResult, error) {
	client := gocloak.NewClient(a.AuthURL)

	tokenResult, err := client.RetrospectToken(ctx, token, a.ClientID, a.ClientSecret, a.Realm)
	if err != nil {
		return nil, err
	}

	return tokenResult, nil
}
