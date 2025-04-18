package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretSize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretSize {
		return nil, fmt.Errorf("invalid key size: %v", secretKey)
	}
	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username, role string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, role, duration)
	if err != nil {
		return "", payload, err
	}

	jwToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
