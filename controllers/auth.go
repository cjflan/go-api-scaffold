package controllers

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"

	middleware "github.com/cjflan/go-api-scaffold/middleware"
)

func Login(ctx *gin.Context) {
	claims, ok := ctx.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	if !ok {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			map[string]string{"message": "Failed to get validated JWT claims."},
		)
		return
	}

	customClaims, ok := claims.CustomClaims.(*middleware.CustomClaimsExample)
	if !ok {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			map[string]string{"message": "Failed to cast custom JWT claims to specific type."},
		)
		return
	}

	if len(customClaims.Username) == 0 {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			map[string]string{"message": "Username in JWT claims was empty."},
		)
		return
	}

	ctx.JSON(http.StatusOK, claims)
}
