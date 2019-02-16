package cache

// Cache is an interface that any concrete type that are supposed to provide caching functionality, should implement it
type Cache interface {
	Get(key string) ([]byte, error)
	Set(key string, content []byte, duration string)
}