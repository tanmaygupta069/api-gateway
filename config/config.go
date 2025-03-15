package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("No .env file found : %v", err)
		return nil, err
	}
	config := &Config{
		ServerConfig: ServerConfig{
			Port: getEnv("PORT"),
		},
		RateLimiterConfig: RateLimiterConfig{
			RateLimit:  getEnvInt("RATE_LIMIT", 2),
			BucketSize: getEnvInt("BUCKET_SIZE", 10),
		},
		GrpcServerConfig: GrpcServerConfig{
			AuthConfig: AuthConfig{
				Port: getEnv("AUTH_GRPC_PORT"),
				Host: getEnv("AUTH_GRPC_HOST"),
			},
			OrderConfig: OrderConfig{
				Port: getEnv("ORDER_GRPC_PORT"),
				Host: getEnv("ORDER_GRPC_HOST"),
			},
		},
	}
	return config, nil
}

func getEnv(key string) string {
	if val, exisit := os.LookupEnv(key); exisit {
		return val
	}
	fmt.Printf("\n%s not found in .env\n", key)
	return ""
}

func getEnvInt(key string, Default int) int {
	if val, exisit := os.LookupEnv(key); exisit {
		res, _ := strconv.Atoi(val)
		return res
	}
	fmt.Printf("\n%s not found in .env\n", key)
	return Default
}
