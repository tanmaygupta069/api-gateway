package middleware

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/tanmaygupta069/api-gateway/config"
)

var cfg,_=config.GetConfig()

func KillSwitch(key string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if cfg.KillSwitchMap[key]{
			ctx.AbortWithError(http.StatusForbidden,fmt.Errorf("kill switch hit for %s",key))
			return
		}
		ctx.Next()
	}
}