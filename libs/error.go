package libs

import "errors"

var ErrorGenerateToken = errors.New("GENERATE_TOKEN_ERROR")
var ErrorSigningMethod = errors.New("UNEXPECTED_SIGNING_METHOD")
var ErrorInvalidToken = errors.New("INVALID_TOKEN")
var ErrorNotFoundEmailAuth = errors.New("NOT_FOUND_EMAIL_AUTH")
var ErrorTokenAlreadyUse = errors.New("TOKEN_ALREADY_USE")
var ErrorTokenExpiredCode = errors.New("EXPIRED_CODE")
var ErrorUserProfileDefine = errors.New("USER_PROFILE_DEFINE")
var ErrorAlreadyExists = errors.New("ALREADY_EXISTS")
var ErrorUserIsMissing = errors.New("USER_IS_MISSING")
var ErrorProviderValid = errors.New("PROVIDER_VALID")
var ErrorNotFound = errors.New("NOT_FOUND")
var ErrorForbidden = errors.New("FORBIDDEN")
var ErrorPermission = errors.New("NO_PERMISSION")
var ErrorParamRequired = errors.New("REQUIRED_DATA")
var ErrorLimited = errors.New("MAX_LIMIT_IS_100")
var ErrorMaxCommentLevel = errors.New("MAXIMUM_COMMENT_LEVEL_IS_2")
var ErrorInvalidCursor = errors.New("INVALID_CURSOR")
var ErrorUpdateFailed = errors.New("DATA_UPDATE_FAIL")
var ErrorEnvParamsNotFound = errors.New("ENV_NOT_FOUND")
var ErrorInvalidData = errors.New("INVALID_DATA")
