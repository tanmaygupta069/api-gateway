package main

import (
	"fmt"

	"github.com/tanmaygupta069/api-gateway/config"
	"github.com/tanmaygupta069/api-gateway/internal/router"
	// "github.com/tanmaygupta069/api-gateway/internal/services/auth"
	// "google.golang.org/grpc"
)

func main() {
	router := router.GetRouter()
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Printf("error in getting config : %v", err)
	}
	router.Run(":" + cfg.ServerConfig.Port)
}
