package token

import (
	"testing"
	"time"

	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	secretKey := utils.RandomString(30)
	maker, err := NewJWTMaker(secretKey)
	require.NoError(t, err)

	userInfo := UserInfo{
		Username: utils.RandomString(6),
		UserId:   utils.RandomInt(0, 99999),
	}
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(userInfo, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, userInfo, payload.UserInfo)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	secretKey := utils.RandomString(30)
	maker, err := NewJWTMaker(secretKey)
	require.NoError(t, err)

	userInfo := UserInfo{
		Username: utils.RandomString(6),
		UserId:   utils.RandomInt(0, 99999),
	}
	token, err := maker.CreateToken(userInfo, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	userInfo := UserInfo{
		Username: utils.RandomString(6),
		UserId:   utils.RandomInt(0, 99999),
	}
	payload, err := NewPayload(userInfo, time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	secretKey := utils.RandomString(30)
	maker, err := NewJWTMaker(secretKey)
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
