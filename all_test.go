package largo

import (
	"flag"
	"testing"

	. "github.com/otiai10/mint"
)

func TestTokenize(t *testing.T) {

	Expect(t, Tokenize(`foo bar baz "hoge fuga" piyo`)).
		ToBe([]string{"foo", "bar", "baz", "hoge fuga", "piyo"})

	Expect(t, Tokenize(`foo 'bar baz' hoge`)).
		ToBe([]string{"foo", "bar baz", "hoge"})

	Expect(t, Tokenize(`foo "bar 'baz hoge' fuga" piyo`)).
		ToBe([]string{"foo", "bar 'baz hoge' fuga", "piyo"})

	Expect(t, Tokenize(`foo 'bar "baz hoge" fuga' piyo`)).
		ToBe([]string{"foo", "bar \"baz hoge\" fuga", "piyo"})

}

func TestFlagSet_Parse(t *testing.T) {
	fset := NewFlagSet("foo", flag.ExitOnError)
	Expect(t, fset).TypeOf("*largo.FlagSet")

	var name string
	fset.StringVar(&name, "name", "otiai10", "Name of the user")
	err := fset.Parse([]string{"hoge", "-name", "ochiai", "baz"})
	Expect(t, err).ToBe(nil)
	Expect(t, fset.Lookup("name").Value.Get()).ToBe("ochiai")
	Expect(t, name).ToBe("ochiai")
	Expect(t, fset.Rest()).ToBe([]string{"hoge", "baz"})
}
