package redis

const (
	KeyPrefix = "pinkacg:"
	KeyEmail  = "email:"
)

func GetRedisKey(key string) string {
	return KeyPrefix + key
}

func GetRedisEmailKey(slug, email string) string {
	return KeyPrefix + KeyEmail + slug + ":" + email
}
