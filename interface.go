package interfaces

type Operations interface {
	Has(key int) bool
	Get(key int) (string, bool)
	Remove(key int) (string, bool)
	Set(key int, value string) (string, bool)
	String() string
}
