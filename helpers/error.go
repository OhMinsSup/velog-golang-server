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

var ErrorProviderValided = errors.New("Provider valid")
