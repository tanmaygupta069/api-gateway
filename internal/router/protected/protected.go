package protected

import (
	"github.com/gin-gonic/gin"
	// "github.com/tanmaygupta069/api-gateway/internal/services/holding"
	"github.com/tanmaygupta069/api-gateway/config"
	"github.com/tanmaygupta069/api-gateway/internal/middleware"
	"github.com/tanmaygupta069/api-gateway/internal/services/holding"
	"github.com/tanmaygupta069/api-gateway/internal/services/order"
)

func RegisterAuthRoutes(rg *gin.RouterGroup,orderService order.OrderService,holdingClient holding.HoldingService) {
	auth := rg.Group("/")
	{
		
		auth.POST("/order",middleware.KillSwitch(config.PLACE_ORDER_KILL_SWITCH),orderService.CreateOrder)
		auth.PUT("/order",middleware.KillSwitch(config.CANCEL_ORDER_KILL_SWITCH),orderService.CancelOrder)
		auth.GET("/orders",middleware.KillSwitch(config.GET_ORDER_HISTORY_KILL_SWITCH),orderService.GetOrderHistory)
		auth.GET("/holdings",middleware.KillSwitch(config.GET_HOLDING_KILL_SWITCH),holdingClient.GetCurrentHoldings)
	}

}
