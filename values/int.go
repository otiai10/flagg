package values

import "strconv"

type IntValue int

func (i *IntValue) Set(s string) error {
	v, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*i = IntValue(v)
	return nil
}

func (i *IntValue) Get() interface{} {
	return int(*i)
}

func (i *IntValue) Type() string {
	return "int"
}
