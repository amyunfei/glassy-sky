package options

import (
	"time"

	"github.com/amyunfei/glassy-sky/cmd/config"
	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/token"
)

type AppOptions struct {
	TimeZone               *time.Location
	tokenMaker             token.Maker
	TokenExpirationMinutes time.Duration
}

func (o AppOptions) GetTimeZone() *time.Location {
	return o.TimeZone
}
func (o AppOptions) GetTokenMaker() token.Maker {
	return o.tokenMaker
}
func (o AppOptions) GetTokenExpirationMinutes() time.Duration {
	return o.TokenExpirationMinutes
}

func NewAppOptions(config *config.Config) (*AppOptions, error) {
	tokenMaker, err := token.NewJWTMaker(config.JWT_SECRET)
	if err != nil {
		return nil, err
	}
	timeZone, err := time.LoadLocation(config.TimeZone)
	if err != nil {
		return nil, err
	}
	return &AppOptions{
		TimeZone:               timeZone,
		tokenMaker:             tokenMaker,
		TokenExpirationMinutes: time.Duration(config.TokenExpirationMinutes),
	}, nil
}
