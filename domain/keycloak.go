package domain

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

type KeycloakRepository interface {
	IntrospectToken(ctx context.Context, token string) (*gocloak.IntroSpectTokenResult, error)
}

type KeycloakService interface {
	IntrospectToken(ctx context.Context, token string) (*gocloak.IntroSpectTokenResult, error)
}
