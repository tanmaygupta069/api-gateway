package order

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	OrderPb "github.com/tanmaygupta069/order-service-go/generated/order"
	"google.golang.org/grpc/metadata"
)

type OrderService interface {
	CreateOrder(ctx *gin.Context)
	CancelOrder(ctx *gin.Context)
	GetOrderHistory(ctx *gin.Context)
	GetStockPrice(ctx *gin.Context)
	CompleteOrder(ctx *gin.Context)
}

type OrderServiceImp struct {
	OrderClient OrderPb.OrderServiceClient
	OrderPb.UnimplementedOrderServiceServer
}

func NewOrderService(client OrderPb.OrderServiceClient) OrderService {
	return &OrderServiceImp{OrderClient: client}
}

func (s *OrderServiceImp) CreateOrder(ctx *gin.Context) {
	fmt.Printf("token")
	token, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Response": gin.H{
				"Code":    http.StatusUnauthorized,
				"Message": http.StatusText(http.StatusUnauthorized),
			},
		})
		return
	}

	// parts := strings.Split(token," ")
	// fmt.Printf(token)

	md := metadata.New(map[string]string{
		"Authorization": token, // Add token in metadata
	})

	contx := metadata.NewOutgoingContext(context.Background(), md)
	fmt.Print("in create order")
	var req OrderPb.OrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.OrderClient.PlaceOrder(contx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Response": gin.H{
				"Code":    http.StatusInternalServerError,
				"Message": err.Error(),
			},
		})
		return
	}
	ctx.JSON(int(res.Response.Code), gin.H{
		"Response": res.Response,
		"Order":    res.Order,
	})
	return
}

func (s *OrderServiceImp) CancelOrder(ctx *gin.Context) {
	token, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Response": gin.H{
				"Code":    http.StatusUnauthorized,
				"Message": http.StatusText(http.StatusUnauthorized),
			},
		})
		return
	}

	md := metadata.New(map[string]string{
		"Authorization": token, // Add token in metadata
	})

	// Create a new context with the metadata
	contx := metadata.NewOutgoingContext(context.Background(), md)
	var req OrderPb.CancelOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.OrderClient.CancelOrder(contx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Response": gin.H{
				"Code":    http.StatusInternalServerError,
				"Message": err.Error(),
			},
		})
		return
	}
	ctx.JSON(int(res.Response.Code), gin.H{
		"Response": res.Response,
		"Order":    res.Order,
	})
	return
}

func (s *OrderServiceImp) GetOrderHistory(ctx *gin.Context) {

	fmt.Print("in order history")
	token, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Response": gin.H{
				"Code":    http.StatusUnauthorized,
				"Message": http.StatusText(http.StatusUnauthorized),
			},
		})
		return
	}
	md := metadata.New(map[string]string{
		"Authorization": token,
	})
	contx := metadata.NewOutgoingContext(context.Background(), md)

	var req OrderPb.OrderHistoryRequest

	res, err := s.OrderClient.GetOrderHistory(contx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Response": gin.H{
				"Code":    http.StatusInternalServerError,
				"Message": err.Error(),
			},
		})
		return
	}
	ctx.JSON(int(res.Response.Code), gin.H{
		"Response": res.Response,
		"Orders":   res.Orders,
	})
	return
}

func (s *OrderServiceImp) GetStockPrice(ctx *gin.Context) {
	symbol := ctx.Query("symbol")
	if len(symbol) < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Response": gin.H{
				"Code":    http.StatusBadGateway,
				"Message": "symbol must have atleast 1 character",
			},
		})
		return
	}

	res, err := s.OrderClient.GetCurrentPrice(ctx, &OrderPb.GetCurrentPriceRequest{
		Symbol: symbol,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Response": gin.H{
				"Code":    http.StatusInternalServerError,
				"Message": err.Error(),
			},
		})
		return
	}
	ctx.JSON(int(res.Response.Code), gin.H{
		"Response": res.Response,
		"Price":    res.Price,
	})
	return
}

func (s *OrderServiceImp) CompleteOrder(ctx *gin.Context) {
	var req OrderPb.CompleteOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.OrderClient.CompleteOrder(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Response": gin.H{
				"Code":    http.StatusInternalServerError,
				"Message": err.Error(),
			},
		})
		return
	}
	ctx.JSON(int(res.Response.Code), gin.H{
		"Response": res.Response,
		"Order":    res.Order,
	})
	return

}
