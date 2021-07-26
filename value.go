package largo

type Value interface {
	Set(s string) error
	Get() interface{}
}
