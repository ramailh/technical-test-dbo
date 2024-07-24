package keys

import "fmt"

func CacheKeyGenerator(idx, key string) string {
	return fmt.Sprintf("%s:%s", idx, key)
}
