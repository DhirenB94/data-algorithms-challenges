package interfaces

type Operations[T any] interface {
	Has(key int) bool
	Get(key int) *T
	Remove(key int) *T
	Set(key int, value T) *T
	String() string
}
