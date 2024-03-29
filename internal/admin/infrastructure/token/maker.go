package token

import "time"

type Maker interface {
	CreateToken(userInfo UserInfo, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
