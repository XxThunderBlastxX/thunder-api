package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gookit/goutil"
)

type Auth interface {
	IsTokenActive() (bool, error)
	IsUserInRealmRole(role string) bool
	TokenHasScope(scope string) bool
}

type accessToken struct {
	claims       jwt.MapClaims
	realmRoles   []string
	accountRoles []string
	scopes       []string
}

func NewAuth(claims jwt.MapClaims) Auth {
	return &accessToken{
		claims:       claims,
		realmRoles:   parseRealmRoles(claims),
		accountRoles: parseAccountRoles(claims),
		scopes:       parseScopes(claims),
	}
}

// IsTokenActive checks if the token is active.
// Returns true if the token is active, false otherwise.
func (a *accessToken) IsTokenActive() (bool, error) {
	timer, err := a.claims.GetExpirationTime()
	if err != nil {
		return false, err
	}
	return !timer.Time.Before(time.Now()), nil
}

func (a *accessToken) IsUserInRealmRole(role string) bool {
	return goutil.Contains(a.realmRoles, role)
}

func (a *accessToken) TokenHasScope(scope string) bool {
	return goutil.Contains(a.scopes, scope)
}

func parseRealmRoles(claims jwt.MapClaims) []string {
	var realmRoles []string = make([]string, 0)

	if claim, ok := claims["realm_access"]; ok {
		if roles, ok := claim.(map[string]interface{})["roles"]; ok {
			for _, role := range roles.([]interface{}) {
				realmRoles = append(realmRoles, role.(string))
			}
		}
	}
	return realmRoles
}

func parseAccountRoles(claims jwt.MapClaims) []string {
	var accountRoles []string = make([]string, 0)

	if claim, ok := claims["resource_access"]; ok {
		if acc, ok := claim.(map[string]interface{})["account"]; ok {
			if roles, ok := acc.(map[string]interface{})["roles"]; ok {
				for _, role := range roles.([]interface{}) {
					accountRoles = append(accountRoles, role.(string))
				}
			}
		}
	}
	return accountRoles
}

func parseScopes(claims jwt.MapClaims) []string {
	scopeStr, err := parseString(claims, "scope")
	if err != nil {
		return make([]string, 0)
	}
	scopes := strings.Split(scopeStr, " ")
	return scopes
}

func parseString(claims jwt.MapClaims, key string) (string, error) {
	var (
		ok  bool
		raw interface{}
		iss string
	)
	raw, ok = claims[key]
	if !ok {
		return "", nil
	}

	iss, ok = raw.(string)
	if !ok {
		return "", fmt.Errorf("key %s is invalid", key)
	}
	return iss, nil
}
