package values

type StringValue string

func (s *StringValue) Set(val string) error {
	*s = StringValue(val)
	return nil
}

func (s *StringValue) Get() interface{} {
	return string(*s)
}

func (s *StringValue) Type() string {
	return "string"
}
