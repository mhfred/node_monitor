package config

var Flags FlagsConfig

type FlagsConfig struct {
	RedisHost string
	RedisPort string
	CleosHost string
	CleosPort string
}
