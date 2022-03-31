package auth

import (
	"context"
	"github.com/ArifulProtik/arifulprotik-api/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JwtCustomClaim struct {
	ID uuid.UUID
	jwt.StandardClaims
}

func GenAccessToken(ctx context.Context, userID uuid.UUID) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 10).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	token, err := t.SignedString([]byte(config.C.Jwt_Key))
	if err != nil {
		return "", err
	}

	return token, nil
}
