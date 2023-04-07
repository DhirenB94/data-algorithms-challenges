package interfaces

import "golang.org/x/exp/constraints"

type Operations[K, V CustomType] interface {
	Has(key K) bool
	Get(key K) *V
	Remove(key K) *V
	Set(key K, value V) *V
	String() string
}

type CustomType interface {
	constraints.Ordered
}
