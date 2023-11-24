package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ndsky1003/jwt/options"
)

var mySigningKey = []byte("b1561c9fc5db1a321206bb821c3a8bc2")

func SetKey(key []byte) {
	mySigningKey = key
}

func New[T any](v T, opts ...*options.NewOptions) (string, error) {
	opt := options.New().Merge(opts...)
	var key []byte
	if opt.Secret != nil {
		key = []byte(*opt.Secret)
	} else {
		key = mySigningKey
	}
	vv := &CustomClaims[T]{
		Payload: v,
	}
	if opt.Issuer != nil {
		vv.Issuer = *opt.Issuer
	}

	if opt.Subject != nil {
		vv.Subject = *opt.Subject
	}

	if opt.Audience != nil {
		vv.Audience = *opt.Audience
	}

	if opt.ExpiresAt != nil {
		vv.ExpiresAt = opt.ExpiresAt
	}

	if opt.NotBefore != nil {
		vv.NotBefore = opt.NotBefore
	}

	if opt.IssuedAt != nil {
		vv.IssuedAt = opt.IssuedAt
	}

	if opt.ID != nil {
		vv.ID = *opt.ID
	}

	fmt.Printf("%+v\n", vv)
	return jwt.NewWithClaims(jwt.SigningMethodHS256, vv).SignedString(key)
}

func Parse[T any](tokenStr string, opts ...*options.NewOptions) (*CustomClaims[T], error) {
	opt := options.New().Merge(opts...)
	var key []byte
	if opt.Secret != nil {
		key = []byte(*opt.Secret)
	} else {
		key = mySigningKey
	}
	var payload T
	v := &CustomClaims[T]{
		Payload: payload,
	}
	_, err := jwt.ParseWithClaims(tokenStr, v, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	return v, nil
}

type CustomClaims[T any] struct {
	Payload T
	jwt.RegisteredClaims
}
