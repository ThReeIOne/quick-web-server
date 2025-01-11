package config

import "os"

const (
	Version              = "VERSION"
	IsDev                = "DEBUG"
	EnableNetwork        = "ENABLE_NETWORK"
	LogLevel             = "LOG_LEVEL"
	ApiPathPrefix        = "API_PATH_PREFIX"
	DBHost               = "DB_HOST"
	DBPort               = "DB_PORT"
	DBDatabase           = "DB_DATABASE"
	DBUsername           = "DB_USERNAME"
	DBPassword           = "DB_PASSWORD"
	RedisAddr            = "REDIS_ADDR"
	RedisPassword        = "REDIS_PASSWORD"
	GatewayAddress       = "GATEWAY_ADDRESS"
	SessionLifeDay       = "SESSION_LIFE_DAY"
	JwtSignKey           = "JWT_SIGN_KEY"
	JwtBufferTime        = "JWT_BUFFER_TIME"
	JwtExpiredTime       = "JWT_EXPIRED_TIME"
	JwtIssuer            = "JWT_ISSUER"
	TencentCosSecretKey  = "TENCENT_COS_SECRET_KEY"
	TencentCosSecretId   = "TENCENT_COS_SECRET_ID"
	TencentCosBasePath   = "TENCENT_COS_BASE_PATH"
	TencentCosBucket     = "TENCENT_COS_BUCKET"
	TencentCosRegion     = "TENCENT_COS_REGION"
	TencentSMSSecretId   = "TENCENT_SMS_SECRET_ID"
	TencentSMSSecretKey  = "TENCENT_SMS_SECRET_KEY"
	TencentSmsAppId      = "TENCENT_SMS_APP_ID"
	TencentSmsTemplateId = "TENCENT_SMS_TEMPLATE_ID"
	TencentSmsSignId     = "TENCENT_SMS_SIGN_ID"
	TencentSmsRegion     = "TENCENT_SMS_REGION"
	RedeemCodeLength     = "REDEEM_CODE_LENGTH"
)

var defaults = map[string]string{
	IsDev:            "true",
	EnableNetwork:    "false",
	LogLevel:         "info",
	ApiPathPrefix:    "/api",
	DBHost:           "127.0.0.1",
	DBPort:           "3306",
	DBDatabase:       "resource",
	DBUsername:       "root",
	DBPassword:       "123456",
	RedisAddr:        "127.0.0.1:6379",
	GatewayAddress:   "0.0.0.0:3000",
	SessionLifeDay:   "1",
	JwtSignKey:       "CeMetaResource",
	JwtBufferTime:    "1d",
	JwtExpiredTime:   "7d",
	JwtIssuer:        "CeMeta",
	TencentCosRegion: "ap-beijing",
	TencentSmsRegion: "ap-beijing",
	RedeemCodeLength: "16",
}

func Get(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	defValue, ok := defaults[key]
	if !ok {
		return ""
	}

	return defValue
}
