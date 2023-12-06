package api

import "github.com/gin-gonic/gin"

func successOrAbort(ctx *gin.Context, code int, err error) (success bool) {
	if err != nil {
		ctx.AbortWithError(code, err)
		ctx.String(code, err.Error())
	}
	return err == nil
}
