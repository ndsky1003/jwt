package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type option struct {
	Secret    *string
	Issuer    *string
	Subject   *string
	Audience  *jwt.ClaimStrings
	ExpiresAt *jwt.NumericDate
	NotBefore *jwt.NumericDate
	IssuedAt  *jwt.NumericDate
	ID        *string
}

func Option() *option {
	return new(option)
}

func (this *option) SetSecret(s string) *option {
	if this == nil {
		return this
	}
	this.Secret = &s
	return this
}

func (this *option) SetIssuer(s string) *option {
	if this == nil {
		return this
	}
	this.Issuer = &s
	return this
}

func (this *option) SetSubject(s string) *option {
	if this == nil {
		return this
	}
	this.Subject = &s
	return this
}

func (this *option) SetAudience(s ...string) *option {
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

func (this *option) SetExpiresAt(t time.Time) *option {
	if this == nil {
		return this
	}
	this.ExpiresAt = jwt.NewNumericDate(t)
	return this
}

func (this *option) SetNotBefore(t time.Time) *option {
	if this == nil {
		return this
	}
	this.NotBefore = jwt.NewNumericDate(t)
	return this
}

func (this *option) SetIssuedAt(t time.Time) *option {
	if this == nil {
		return this
	}
	this.IssuedAt = jwt.NewNumericDate(t)
	return this
}

func (this *option) SetID(v string) *option {
	if this == nil {
		return this
	}
	this.ID = &v
	return this
}

func (this *option) Merge(opts ...*option) *option {
	for _, opt := range opts {
		this.merge(opt)
	}
	return this
}

func (this *option) merge(opt *option) {
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
