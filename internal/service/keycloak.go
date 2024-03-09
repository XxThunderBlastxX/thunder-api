package service

import (
	"context"

	"github.com/Nerzal/gocloak/v13"

	"github.com/XxThunderBlastxX/thunder-api/domain"
)

type keycloakService struct {
	keycloakRepo domain.KeycloakRepository
}

func NewKeycloakService(keycloakRepo domain.KeycloakRepository) domain.KeycloakService {
	return &keycloakService{
		keycloakRepo: keycloakRepo,
	}
}

func (k *keycloakService) IntrospectToken(ctx context.Context, token string) (*gocloak.IntroSpectTokenResult, error) {
	return k.keycloakRepo.IntrospectToken(ctx, token)
}
