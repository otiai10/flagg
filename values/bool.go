package values

import "strconv"

type BoolValue bool

func (b *BoolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	*b = BoolValue(v)
	return nil
}

func (b *BoolValue) Get() interface{} {
	return bool(*b)
}

func (b *BoolValue) Type() string {
	return "bool"
}
