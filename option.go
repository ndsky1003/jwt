package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Option struct {
	Secret    *string
	Issuer    *string
	Subject   *string
	Audience  *jwt.ClaimStrings
	ExpiresAt *jwt.NumericDate
	NotBefore *jwt.NumericDate
	IssuedAt  *jwt.NumericDate
	ID        *string
}

func Options() *Option {
	return &Option{}
}

func (this *Option) SetSecret(s string) *Option {
	if this == nil {
		return this
	}
	this.Secret = &s
	return this
}

func (this *Option) SetIssuer(s string) *Option {
	if this == nil {
		return this
	}
	this.Issuer = &s
	return this
}

func (this *Option) SetSubject(s string) *Option {
	if this == nil {
		return this
	}
	this.Subject = &s
	return this
}

func (this *Option) SetAudience(s ...string) *Option {
	if this == nil {
		return this
	}
	var a = make(jwt.ClaimStrings, 0, len(s))
	for _, v := range s {
		a = append(a, v)
	}
	this.Audience = &a
	return this
}

func (this *Option) SetExpiresAt(t time.Time) *Option {
	if this == nil {
		return this
	}
	this.ExpiresAt = jwt.NewNumericDate(t)
	return this
}

func (this *Option) SetNotBefore(t time.Time) *Option {
	if this == nil {
		return this
	}
	this.NotBefore = jwt.NewNumericDate(t)
	return this
}

func (this *Option) SetIssuedAt(t time.Time) *Option {
	if this == nil {
		return this
	}
	this.IssuedAt = jwt.NewNumericDate(t)
	return this
}

func (this *Option) SetID(v string) *Option {
	if this == nil {
		return this
	}
	this.ID = &v
	return this
}

func (this *Option) Merge(opts ...*Option) *Option {
	for _, opt := range opts {
		this.merge(opt)
	}
	return this
}

func (this *Option) merge(opt *Option) {
	if opt.Secret != nil {
		this.Secret = opt.Secret
	}

	if opt.Issuer != nil {
		this.Issuer = opt.Issuer
	}

	if opt.Subject != nil {
		this.Subject = opt.Subject
	}

	if opt.Audience != nil {
		this.Audience = opt.Audience
	}

	if opt.ExpiresAt != nil {
		this.ExpiresAt = opt.ExpiresAt
	}

	if opt.NotBefore != nil {
		this.NotBefore = opt.NotBefore
	}

	if opt.IssuedAt != nil {
		this.IssuedAt = opt.IssuedAt
	}

	if opt.ID != nil {
		this.ID = opt.ID
	}
}
