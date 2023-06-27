package redis

type RedisCache interface {
	SetCache(key string, value interface{})
	GetCache(key string, requested string) interface{}
}
