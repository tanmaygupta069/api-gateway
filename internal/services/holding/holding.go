package holding

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	HoldingPb "github.com/tanmaygupta069/order-service-go/generated/holding"
	"google.golang.org/grpc/metadata"
)

type HoldingService interface {
	GetCurrentHoldings(ctx *gin.Context)
}

type HoldingServiceImp struct{
	HoldingClient HoldingPb.HoldingServiceClient
	HoldingPb.UnimplementedHoldingServiceServer
}

func NewHoldingService(HoldingClient HoldingPb.HoldingServiceClient)HoldingService{
	return &HoldingServiceImp{HoldingClient: HoldingClient}
}

func(s *HoldingServiceImp)GetCurrentHoldings(ctx *gin.Context){
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

	contx := metadata.NewOutgoingContext(context.Background(), md)	

	var req HoldingPb.CurrentHoldingsRequest

	res, err := s.HoldingClient.GetCurrentHoldings(contx, &req)
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
		"Holdings":   res.Holdings,
	})
	return
}
