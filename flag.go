package largo

// Alias defines aliases for the flag.
func (f *Flag) Alias(aliases ...string) *Flag {
	f.aliases = append(f.aliases, aliases...)
	return f
}
