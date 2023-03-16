package token

import (
	"errors"
	"time"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	} else {
		return nil
	}
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	return &Payload{
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, nil
}
