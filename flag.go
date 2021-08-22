package largo

type (
	Flag struct {
		Name    string
		Value   Value
		Usage   string
		aliases []string
		given   bool
	}
)

// Alias defines aliases for the flag.
func (f *Flag) Alias(aliases ...string) *Flag {
	f.aliases = append(f.aliases, aliases...)
	return f
}

func (f *Flag) Aliases() []string {
	return f.aliases
}

func (f *Flag) Given() bool {
	if f == nil {
		return false
	}
	return f.given
}
