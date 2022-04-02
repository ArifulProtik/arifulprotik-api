package auth

import (
	"context"
	"github.com/ArifulProtik/arifulprotik-api/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JwtCustomClaim struct {
	ID      uuid.UUID
	RF_UUID uuid.UUID
	jwt.StandardClaims
}

func GenAccessToken(ctx context.Context, userID uuid.UUID) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	token, err := t.SignedString([]byte(config.C.Jwt_Key))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenRefreshToken(ctx context.Context, userID uuid.UUID, rf_id uuid.UUID) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		ID:      userID,
		RF_UUID: rf_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	token, err := t.SignedString([]byte(config.C.Jwt_Key))
	if err != nil {
		return "", err
	}

	return token, nil
}
