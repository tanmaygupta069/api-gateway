package middleware

import (
	"errors"
	// "fmt"
	"net/http"

	pb "github.com/tanmaygupta069/auth-service-go/generated"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(client pb.AuthServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("Authorization")
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		if token==""{
			ctx.AbortWithError(http.StatusUnauthorized,errors.New("token expired login"))
		}
		res,err:=client.ValidateToken(ctx,&pb.ValidateTokenRequest{Token:token})
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		if res.Valid == false{
			ctx.AbortWithError(int(res.Response.Code), errors.New(res.Response.Message))
			return
		}
		ctx.Next()
	}
}