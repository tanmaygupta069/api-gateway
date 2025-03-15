package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanmaygupta069/api-gateway/config"
	Grpc "github.com/tanmaygupta069/api-gateway/internal/grpc"
	"github.com/tanmaygupta069/api-gateway/internal/middleware"
	"github.com/tanmaygupta069/api-gateway/internal/router/protected"
	"github.com/tanmaygupta069/api-gateway/internal/services/auth"
	"github.com/tanmaygupta069/api-gateway/internal/services/holding"
	"github.com/tanmaygupta069/api-gateway/internal/services/order"
	// auth_pb "github.com/tanmaygupta069/auth-service-go/generated"
	// order_pb "github.com/tanmaygupta069/order-service-go/generated"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Printf("error in getting config api gateway router")
	}

	router.Use(middleware.RateLimitMiddleware())

	authClient, err := Grpc.NewAuthClient(cfg.GrpcServerConfig.AuthConfig.Host+":"+cfg.GrpcServerConfig.AuthConfig.Port)
	if err != nil {
		fmt.Printf("error in grpc auth client connection : %v", err)
	}

	orderClient, err := Grpc.NewOrderClient(cfg.GrpcServerConfig.OrderConfig.Host+":"+cfg.GrpcServerConfig.OrderConfig.Port)
	if err != nil {
		fmt.Printf("error in grpc order client connection : %v", err)
	}

	holdingClient, err := Grpc.NewHoldingClient(cfg.GrpcServerConfig.OrderConfig.Host+":"+cfg.GrpcServerConfig.OrderConfig.Port)
	if err != nil {
		fmt.Printf("error in grpc holding client connection : %v", err)
	}


	authService := auth.NewService(authClient)
	orderService := order.NewOrderService(orderClient)
	holdingService := holding.NewHoldingService(holdingClient)

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello https",
		})
		return
	})
	router.POST("/login", authService.Login)
	router.POST("/signup", authService.Signup)
	router.GET("/stock",orderService.GetStockPrice)
	router.PUT("/order",orderService.CompleteOrder)

	auth := router.Group("/auth")
	{
		auth.Use(middleware.AuthMiddleware(authClient))
		protected.RegisterAuthRoutes(auth,orderService,holdingService)
	}

	return router
}
