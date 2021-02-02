package controllers

import (
	"github.com/OhMinsSup/story-server/app"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SendEmailController 이메일 인증 코드 발급을 위한 코드
func SendEmailController(ctx *gin.Context) {
	var body dto.SendEmailBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, app.BadRequestErrorResponse(err.Error(), nil))
		return
	}

	result, _ := services.SendEmailService(body, ctx)
	ctx.JSON(result.Code, result)
}

// LocalRegisterController
func LocalRegisterController(ctx *gin.Context) {
	//var body dto.LocalRegisterBody
	//if err := ctx.BindJSON(&body); err != nil {
	//	ctx.AbortWithStatus(http.StatusBadRequest)
	//	return
	//}
	//
	//db := ctx.MustGet("db").(*gorm.DB)
	//
	//result, code, err := services.LocalRegisterService(body, db, ctx)
	//if err != nil {
	//	ctx.AbortWithError(code, err)
	//	return
	//}
	//
	//ctx.JSON(code, result)
}

// CodeController
func CodeController(ctx *gin.Context) {
	//// validation Code Params
	//var params = dto.CodeParams{
	//	Code: ctx.Param("code"),
	//}
	//
	//if params.Code == "" {
	//	ctx.AbortWithError(http.StatusBadRequest, errors.New("Code is Required"))
	//	return
	//}
	//
	//db := ctx.MustGet("db").(*gorm.DB)
	//
	//result, code, err := services.CodeService(params.Code, db, ctx)
	//if err != nil {
	//	ctx.AbortWithError(code, err)
	//	return
	//}
	//
	//ctx.JSON(code, result)
}

// LogoutController
func LogoutController(ctx *gin.Context) {
	//helpers.SetCookie(ctx, "", "")
	//ctx.Status(http.StatusNoContent)
}
