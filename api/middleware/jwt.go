package middleware

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"

	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	contribJwt "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"

	"github.com/XxThunderBlastxX/thunder-api/internal/model"
)

func NewJWTMiddleware() fiber.Handler {
	k, err := getAuthPublicKey()
	if err != nil {
		panic(err)
	}
	publicKey, err := parseAuthPublicKey(k)
	if err != nil {
		panic(err)
	}

	return contribJwt.New(contribJwt.Config{
		SigningKey: contribJwt.SigningKey{
			JWTAlg: contribJwt.RS256,
			Key:    publicKey,
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(model.WebResponse[*model.ErrorResponse]{
				Success: false,
				Error:   err.Error(),
			})
		},
	})
}

func parseAuthPublicKey(base64String string) (*rsa.PublicKey, error) {
	buff, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}

	parsedKey, err := x509.ParseCertificate(buff)
	if err != nil {
		return nil, err
	}

	if publicKey, ok := parsedKey.PublicKey.(*rsa.PublicKey); ok {
		return publicKey, nil
	} else {
		return nil, errors.Errorf("unexpected key type %T", publicKey)
	}
}

func getAuthPublicKey() (string, error) {
	var key string
	var data interface{}

	client := resty.New()

	res, err := client.R().Get("https://thunder.jp.auth0.com/.well-known/jwks.json")
	if err != nil {
		return "", err
	}

	if res.StatusCode() != 200 {
		return "", errors.New("could not fetch public key")
	}

	err = json.Unmarshal(res.Body(), &data)
	if err != nil {
		return "", err
	}

	key = data.(map[string]interface{})["keys"].([]interface{})[0].(map[string]interface{})["x5c"].([]interface{})[0].(string)

	return key, nil
}
