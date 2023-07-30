package token

import (
	"errors"
	"time"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type UserInfo struct {
	Username string `json:"username"`
	UserId   int64  `json:"userId"`
}
type Payload struct {
	UserInfo
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

func NewPayload(userInfo UserInfo, duration time.Duration) (*Payload, error) {
	return &Payload{
		UserInfo:  userInfo,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, nil
}
