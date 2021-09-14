package redis

const (
	KeyPrefix = "pinkacg:"
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
