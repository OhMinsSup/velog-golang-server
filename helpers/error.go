package helpers

import "errors"

var ErrorGenerateToken = errors.New("Generate Token Error")
var ErrorSigningMethod = errors.New("Unexpected signing method")
var ErrorInvalidToken = errors.New("invalid token")

var ErrorNotFoundEmailAuth = errors.New("Not Found Email Auth")
var ErrorTokenAlreadyUse = errors.New("Token Already Use")
var ErrorTokenExpiredCode = errors.New("Expireed Code")

var ErrorUserProfileDefine = errors.New("User Profile Define")
var ErrorAlreadyExists = errors.New("Already Exists")
var ErrorUserIsMissing = errors.New("User is missing")
var ErrorProviderValided = errors.New("Provider valid")
var ErrorNotFound = errors.New("Not Found")
var ErrorForbidden = errors.New("Forbidden")
