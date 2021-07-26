package largo

import (
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
