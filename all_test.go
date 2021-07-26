package largo

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestTokenize(t *testing.T) {
	for _, c := range []struct {
		Input    string
		Epxected []string
	}{
		{
			`foo bar baz "hoge fuga" piyo`,
			[]string{"foo", "bar", "baz", "hoge fuga", "piyo"},
		},
		{
			`foo bar 'baz hoge' fuga piyo`,
			[]string{"foo", "bar", "baz hoge", "fuga", "piyo"},
		},
		{
			`foo 'bar "baz hoge" fuga' piyo`,
			[]string{"foo", "bar \"baz hoge\" fuga", "piyo"},
		},
	} {
		Expect(t, Tokenize(c.Input)).ToBe(c.Epxected)
	}
}

func TestFlagSet_Parse(t *testing.T) {
	var name string
	var upper bool
	var count int
	fset := NewFlagSet("foo", ContinueOnError)

	fset.StringVar(&name, "name", "otiai10", "Name of the user").Alias("n")
	fset.BoolVar(&upper, "upper", false, "To upper case of the given name")
	fset.IntVar(&count, "c", 7, "Count to say names")

	err := fset.Parse([]string{"hoge", "-name", "ochiai", "baz", "-upper", "-undefined", "-c", "2", "-foo=baa", "--"})
	Expect(t, err).ToBe(nil)
	Expect(t, fset.Lookup("name").Value.Get()).ToBe("ochiai")
	Expect(t, fset.Lookup("n").Value.Get()).ToBe("ochiai")
	Expect(t, name).ToBe("ochiai")
	Expect(t, upper).ToBe(true)
	Expect(t, count).ToBe(2)
	Expect(t, fset.Rest()).ToBe([]string{"hoge", "baz"})

	err = fset.Parse([]string{"foo", "---unko"})
	Expect(t, err).Not().ToBe(nil)
}
