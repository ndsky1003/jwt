package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte("b1561c9fc5db1a321206bb821c3a8bc2")

func SetKey(key []byte) {
	mySigningKey = key
}

func New[T any](v T, opts ...*Option) (string, error) {
	opt := Options().Merge(opts...)
	var key []byte
	if opt.Secret != nil {
		key = []byte(*opt.Secret)
	} else {
		key = mySigningKey
	}
	vv := &CustomClaims[T]{
		P: v,
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

	return jwt.NewWithClaims(jwt.SigningMethodHS256, vv).SignedString(key)
}

// 就算有错,令牌过期这些逻辑错误,都会返回这个token所承载的内容
func Parse[T any](token string, opts ...*Option) (*CustomClaims[T], error) {
	opt := Options().Merge(opts...)
	var key []byte
	if opt.Secret != nil {
		key = []byte(*opt.Secret)
	} else {
		key = mySigningKey
	}
	var payload T
	v := &CustomClaims[T]{
		P: payload,
	}
	_, err := jwt.ParseWithClaims(token, v, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	return v, err //有些时候错误也是要其内容,比如过期,也可以自动续签
}

type CustomClaims[T any] struct {
	P T
	jwt.RegisteredClaims
}
