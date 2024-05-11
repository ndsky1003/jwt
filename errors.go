package jwt

import "github.com/golang-jwt/jwt/v5"

var (
	ErrInvalidKey                = jwt.ErrInvalidKey
	ErrInvalidKeyType            = jwt.ErrInvalidKeyType
	ErrHashUnavailable           = jwt.ErrHashUnavailable
	ErrTokenMalformed            = jwt.ErrTokenMalformed
	ErrTokenUnverifiable         = jwt.ErrTokenUnverifiable
	ErrTokenSignatureInvalid     = jwt.ErrTokenSignatureInvalid
	ErrTokenRequiredClaimMissing = jwt.ErrTokenRequiredClaimMissing
	ErrTokenInvalidAudience      = jwt.ErrTokenInvalidAudience
	ErrTokenExpired              = jwt.ErrTokenExpired
	ErrTokenUsedBeforeIssued     = jwt.ErrTokenUsedBeforeIssued
	ErrTokenInvalidIssuer        = jwt.ErrTokenInvalidIssuer
	ErrTokenInvalidSubject       = jwt.ErrTokenInvalidSubject
	ErrTokenNotValidYet          = jwt.ErrTokenNotValidYet
	ErrTokenInvalidId            = jwt.ErrTokenInvalidId
	ErrTokenInvalidClaims        = jwt.ErrTokenInvalidClaims
	ErrInvalidType               = jwt.ErrInvalidType
)
