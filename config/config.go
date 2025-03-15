package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	SIGNUP_KILL_SWITCH = "SIGNUP_KILL_SWITCH"
	LOGIN_KILL_SWITCH = "LOGIN_KILL_SWITCH"
	PLACE_ORDER_KILL_SWITCH = "PLACE_ORDER_KILL_SWITCH"
	CANCEL_ORDER_KILL_SWITCH = "CANCEL_ORDER_KILL_SWITCH"
	GET_ORDER_HISTORY_KILL_SWITCH = "GET_ORDER_HISTORY_KILL_SWITCH"
	GET_CUR_PRICE_KILL_SWITCH = "GET_CUR_PRICE_KILL_SWITCH"
	COMPLETE_ORDER_KILL_SWITCH = "COMPLETE_ORDER_KILL_SWITCH"
	GET_HOLDING_KILL_SWITCH = "GET_HOLDING_KILL_SWITCH"
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
		KillSwitchMap: map[string]bool{
			LOGIN_KILL_SWITCH : getEnvBool(LOGIN_KILL_SWITCH),
			SIGNUP_KILL_SWITCH: getEnvBool(SIGNUP_KILL_SWITCH),
			CANCEL_ORDER_KILL_SWITCH:getEnvBool(CANCEL_ORDER_KILL_SWITCH),
			COMPLETE_ORDER_KILL_SWITCH:getEnvBool(COMPLETE_ORDER_KILL_SWITCH),
			GET_CUR_PRICE_KILL_SWITCH:getEnvBool(GET_CUR_PRICE_KILL_SWITCH),
			GET_HOLDING_KILL_SWITCH:getEnvBool(GET_HOLDING_KILL_SWITCH),
			GET_ORDER_HISTORY_KILL_SWITCH:getEnvBool(GET_ORDER_HISTORY_KILL_SWITCH),
			PLACE_ORDER_KILL_SWITCH:getEnvBool(PLACE_ORDER_KILL_SWITCH),
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

func getEnvBool(key string)bool{
	val, _ := os.LookupEnv(key)
		if val == "true"{
			return true
		}else{
			return false
		}
	}

func getEnvInt(key string, Default int) int {
	if val, exisit := os.LookupEnv(key); exisit {
		res, _ := strconv.Atoi(val)
		return res
	}
	fmt.Printf("\n%s not found in .env\n", key)
	return Default
}
