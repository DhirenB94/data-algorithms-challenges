package interfaces

type Operations interface {
	Has(key interface{}) bool
	Get(key interface{}) (interface{}, bool)
	Remove(key interface{}) (interface{}, bool)
	Set(key interface{}, value interface{}) (interface{}, bool)
	String() string
}
