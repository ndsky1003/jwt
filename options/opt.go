package options

import "github.com/golang-jwt/jwt/v5"

type NewOptions struct {
	Secret    *string
	Issuer    *string
	Subject   *string
	Audience  *jwt.ClaimStrings
	ExpiresAt *jwt.NumericDate
	NotBefore *jwt.NumericDate
	IssuedAt  *jwt.NumericDate
	ID        *string
}

func New() *NewOptions {
	return new(NewOptions)
}

func (this *NewOptions) SetSecret(s string) *NewOptions {
	if this == nil {
		return this
	}
	this.Secret = &s
	return this
}

func (this *NewOptions) SetIssuer(s string) *NewOptions {
	if this == nil {
		return this
	}
	this.Issuer = &s
	return this
}

func (this *NewOptions) SetSubject(s string) *NewOptions {
	if this == nil {
		return this
	}
	this.Subject = &s
	return this
}

func (this *NewOptions) SetAudience(s *jwt.ClaimStrings) *NewOptions {
	if this == nil {
		return this
	}
	this.Audience = s
	return this
}

func (this *NewOptions) SetExpiresAt(s *jwt.NumericDate) *NewOptions {
	if this == nil {
		return this
	}
	this.ExpiresAt = s
	return this
}

func (this *NewOptions) SetNotBefore(s *jwt.NumericDate) *NewOptions {
	if this == nil {
		return this
	}
	this.NotBefore = s
	return this
}

func (this *NewOptions) SetIssuedAt(s *jwt.NumericDate) *NewOptions {
	if this == nil {
		return this
	}
	this.IssuedAt = s
	return this
}

func (this *NewOptions) SetID(v string) *NewOptions {
	if this == nil {
		return this
	}
	this.ID = &v
	return this
}

func (this *NewOptions) Merge(opts ...*NewOptions) *NewOptions {
	for _, opt := range opts {
		this.merge(opt)
	}
	return this
}

func (this *NewOptions) merge(opt *NewOptions) {
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

// ------------------ParseOption------------------
type ParseOption struct {
	Secret *string
}

func Parse() *ParseOption {
	return new(ParseOption)
}

func (this *ParseOption) SetSecret(s string) *ParseOption {
	if this == nil {
		return this
	}
	this.Secret = &s
	return this
}

func (this *ParseOption) Merge(opts ...*ParseOption) *ParseOption {
	for _, opt := range opts {
		this.merge(opt)
	}
	return this
}

func (this *ParseOption) merge(opt *ParseOption) {
	if opt.Secret != nil {
		this.Secret = opt.Secret
	}
}
