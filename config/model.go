package config

type Config struct{
	ServerConfig ServerConfig
	RateLimiterConfig RateLimiterConfig
	GrpcServerConfig GrpcServerConfig
	KillSwitchMap map[string]bool
}

type ServerConfig struct{
	Port string
}

type RateLimiterConfig struct{
	RateLimit int
	BucketSize int
}

type GrpcServerConfig struct{
	AuthConfig AuthConfig
	OrderConfig OrderConfig
}

type AuthConfig struct{
	Port string
	Host string
}

type OrderConfig struct{
	Port string
	Host string
}