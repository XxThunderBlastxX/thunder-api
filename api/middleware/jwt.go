package middleware

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"

	contribJwt "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	golangJwt "github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"

	"github.com/XxThunderBlastxX/thunder-api/domain"
	"github.com/XxThunderBlastxX/thunder-api/internal/common/enum"
	"github.com/XxThunderBlastxX/thunder-api/internal/gen/keycloakconfig"
	"github.com/XxThunderBlastxX/thunder-api/internal/model"
)

func NewJWTMiddleware(keycloakService domain.KeycloakService, keycloakConfig *keycloakconfig.KeycloakConfig) fiber.Handler {
	publicKey, err := parseKeycloakPublicKey(keycloakConfig.RealmRSA256PublicKey)
	if err != nil {
		panic(err)
	}

	return contribJwt.New(contribJwt.Config{
		SigningKey: contribJwt.SigningKey{
			JWTAlg: contribJwt.RS256,
			Key:    publicKey,
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			return successHandler(c, keycloakService)
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(model.WebResponse[*model.ErrorResponse]{
				Success: false,
				Error:   err.Error(),
			})
		},
	})
}

func successHandler(c *fiber.Ctx, keycloakService domain.KeycloakService) error {
	jwtToken := c.Locals("user").(*golangJwt.Token)
	claims := jwtToken.Claims.(golangJwt.MapClaims)

	ctx := c.UserContext()

	contextWithClaims := context.WithValue(ctx, enum.ContextKeyClaims, claims)
	c.SetUserContext(contextWithClaims)

	tokenInfo, err := keycloakService.IntrospectToken(ctx, jwtToken.Raw)
	if err != nil {
		panic(err)
	}

	if !*tokenInfo.Active {
		return c.Status(fiber.StatusUnauthorized).JSON(model.WebResponse[*model.ErrorResponse]{
			Success: false,
			Error:   "token is not active",
		})
	}

	return c.Next()
}

func parseKeycloakPublicKey(base64String string) (*rsa.PublicKey, error) {
	buff, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}

	parsedKey, err := x509.ParsePKIXPublicKey(buff)
	if err != nil {
		return nil, err
	}

	if publicKey, ok := parsedKey.(*rsa.PublicKey); ok {
		return publicKey, nil
	} else {
		return nil, errors.Errorf("unexpected key type %T", publicKey)
	}
}
