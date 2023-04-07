package interfaces

import (
	"golang.org/x/exp/constraints"
	"hash/fnv"
)

type Operations[K, V CustomType] interface {
	Has(key K) bool
	Get(key K) *V
	Remove(key K) *V
	Set(key K, value V) *V
	String() string
}

type CustomType interface {
	constraints.Ordered
	HashMe() int
}

type MyCustomInt int

func (mci MyCustomInt) HashMe() int {
	return int(mci % 50)
}

type MyCustomString string

func (mcs MyCustomString) HashMe() int {
	hash := fnv.New32a()
	hash.Write([]byte(mcs))
	hashValue := int(hash.Sum32())
	return hashValue % 50
}
