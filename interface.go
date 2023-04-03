package interfaces

type Operations[K comparable, V any] interface {
	Has(key K) bool
	Get(key K) *V
	Remove(key K) *V
	Set(key K, value V) *V
	String() string
}
