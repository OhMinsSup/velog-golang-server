package libs

import "errors"

var ErrorGenerateToken = errors.New("GENERATE_TOKEN_ERROR")
var ErrorSigningMethod = errors.New("UNEXPECTED_SIGNING_METHOD")
var ErrorInvalidToken = errors.New("INVALID_TOKEN")
