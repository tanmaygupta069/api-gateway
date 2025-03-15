package protected

import (
	"github.com/gin-gonic/gin"
	// "github.com/tanmaygupta069/api-gateway/internal/services/holding"
	"github.com/tanmaygupta069/api-gateway/internal/services/holding"
	"github.com/tanmaygupta069/api-gateway/internal/services/order"
)

func RegisterAuthRoutes(rg *gin.RouterGroup,orderService order.OrderService,holdingClient holding.HoldingService) {
	auth := rg.Group("/")
	{
		
		auth.POST("/order",orderService.CreateOrder)
		auth.DELETE("/order",orderService.CancelOrder)
		auth.GET("/orders",orderService.GetOrderHistory)
		auth.GET("/holdings",holdingClient.GetCurrentHoldings)
	}

}
