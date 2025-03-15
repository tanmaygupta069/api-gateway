package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/tanmaygupta069/auth-service-go/generated"
)

type AuthService interface {
	Login(ctx *gin.Context)
	Signup(ctx *gin.Context)
}

type AuthServiceImp struct {
	client pb.AuthServiceClient
	pb.UnimplementedAuthServiceServer
}

func NewService(client pb.AuthServiceClient) AuthService {
	return &AuthServiceImp{client: client}
}

func (s *AuthServiceImp) Login(ctx *gin.Context) {
	var req pb.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// grpcCtx := ctx.Request.Context()

	res, err := s.client.Login(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error in api-gateway login": err.Error()})
		return
	}

	ctx.SetCookie("Authorization", res.Token, 3600*24, "/", "", true, true)

	ctx.JSON(int(res.Response.Code), gin.H{
		"response": res.Response,
		"token":    res.Token,
	})

	return
}

func (s *AuthServiceImp) Signup(ctx *gin.Context) {
	var req pb.SignupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	grpcCtx := ctx.Request.Context()

	res, err := s.client.Signup(grpcCtx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error in api-gateway signup": err.Error()})
		return
	}

	ctx.JSON(int(res.Response.Code), gin.H{
		"response": res.Response,
	})
	return
}
