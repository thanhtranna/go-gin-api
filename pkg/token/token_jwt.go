package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (t *token) JwtSign(userId int64, userName string, expireDuration time.Duration) (tokenString string, err error) {
	// The token content.
	// iss: (Issuer) issuer
	// iat: (Issued At) issuance time, expressed in Unix timestamp
	// exp: (Expiration Time) expiration time, expressed in Unix timestamp
	// aud: (Audience) The party receiving the JWT
	// sub: (Subject) the subject of the JWT
	// nbf: (Not Before) Don’t be earlier than this time
	// jti: (JWT ID) the unique ID used to identify JWT
	claims := claims{
		userId,
		userName,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
		},
	}
	tokenString, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.secret))
	return
}

func (t *token) JwtParse(tokenString string) (*claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
