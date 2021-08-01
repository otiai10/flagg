package largo

import (
	"bytes"
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

	When(t, "NO-BREAK SPACE is included, it should be splitted", func(t *testing.T) {
		// https://www.utf8-chartable.de/
		b1 := []byte{97, 97, 32, 98, 98}       // [32] = SPACE
		b2 := []byte{97, 97, 194, 160, 98, 98} // [194, 160] == NO-BREAK SPACE
		Expect(t, Tokenize(string(b1))).ToBe([]string{"aa", "bb"})
		Expect(t, Tokenize(string(b2))).ToBe([]string{"aa", "bb"})
	})
}

func TestFlagSet_ParseLine(t *testing.T) {
	verbose := false
	fset := NewFlagSet("dosomething", PanicOnError)
	fset.BoolVar(&verbose, "verbose", false, "Show verbose log").Alias("v")
	fset.ParseLine("dosomething create aws -v ec2-instance")
	Expect(t, verbose).ToBe(true)
	Expect(t, fset.Rest()).ToBe([]string{"create", "aws", "ec2-instance"})

	When(t, "invalid input with PanicOnError", func(t *testing.T) {
		defer func() {
			re := recover()
			Expect(t, re).Not().ToBe(nil)
		}()
		fset.ParseLine("dosomething three minuses --- is invalid")
	})
}

func TestFlagSet_Parse(t *testing.T) {
	var name string
	var upper bool
	var count int
	fset := NewFlagSet("foo", ContinueOnError)

	fset.StringVar(&name, "name", "otiai10", "Name of the user").Alias("n")
	fset.BoolVar(&upper, "upper", false, "To upper case of the given name")
	fset.IntVar(&count, "c", 7, "Count to say names")

	err := fset.Parse([]string{"hoge", "-name=ochiai", "baz", "-upper", "-undefined", "-c", "2", "-foo=baa", "--"})
	Expect(t, err).ToBe(nil)
	Expect(t, fset.Lookup("name").Value.Get()).ToBe("ochiai")
	Expect(t, fset.Lookup("n").Value.Get()).ToBe("ochiai")
	Expect(t, name).ToBe("ochiai")
	Expect(t, upper).ToBe(true)
	Expect(t, count).ToBe(2)
	Expect(t, fset.Rest()).ToBe([]string{"hoge", "baz"})

	err = fset.Parse([]string{"foo", "---unko"})
	Expect(t, err).Not().ToBe(nil)

	When(t, "given empty arguments, it does nothing", func(t *testing.T) {
		err := fset.Parse([]string{})
		Expect(t, err).ToBe(nil)
	})

	When(t, "include single minus char, stop parsing following args", func(t *testing.T) {
		err := fset.Parse([]string{"hoge", "-upper=false", "-", "-name", "Hiromu"})
		Expect(t, err).ToBe(nil)
		Expect(t, upper).ToBe(false)
		Expect(t, name).Not().ToBe("Hiromu")
	})

	When(t, "help is given, print help message", func(t *testing.T) {
		fset.Output = bytes.NewBuffer(nil)
		err := fset.Parse([]string{"hoge", "-upper=false", "-h", "-name", "Hiromu"})
		Expect(t, err).ToBe(nil)
		Expect(t, (fset.Output).(*bytes.Buffer).Len()).Not().ToBe(0)
	})
	When(t, "help with custom Usage func is set", func(t *testing.T) {
		var msg string
		fset.Usage = func() {
			msg = "Hello"
		}
		err := fset.Parse([]string{"hoge", "-upper=false", "-h", "-name", "Hiromu"})
		Expect(t, err).ToBe(nil)
		Expect(t, msg).ToBe("Hello")
		Expect(t, fset.HelpRequested()).ToBe(true)
	})
}

func TestFlagSet_DefaultUsage(t *testing.T) {
	var name string
	var upper bool
	var count int
	fset := NewFlagSet("greet", ContinueOnError)
	fset.Description = "Greet is an amazing command to configure everything in the world."

	fset.StringVar(&name, "name", "otiai10", "Name of the user")
	fset.BoolVar(&upper, "upper", false, "To upper case of the given name").Alias("U")
	fset.IntVar(&count, "count", 7, "Count to say names").Alias("c")

	buf := bytes.NewBuffer(nil)
	err := fset.PrintDefaultUsage(buf)
	Expect(t, err).ToBe(nil)
	Expect(t, buf.String()).ToBe(`NAME
  greet

DESCRIPTION
  Greet is an amazing command to configure everything in the world.

OPTIONS
  -count int, -c={int}
        Count to say names
  -name string
        Name of the user
  -upper, -U
        To upper case of the given name
`)

	When(t, "without description", func(t *testing.T) {
		fset := NewFlagSet("greet", ContinueOnError)
		fset.BoolVar(&upper, "upper", false, "To upper case of the given name").Alias("U")
		buf := bytes.NewBuffer(nil)
		err := fset.PrintDefaultUsage(buf)
		Expect(t, err).ToBe(nil)
		Expect(t, buf.String()).ToBe(`NAME
  greet

OPTIONS
  -upper, -U
        To upper case of the given name
`)

	})

	Because(t, "var must NOT be nil pointer", func(t *testing.T) {
		defer func() {
			r := recover()
			Expect(t, r).Not().ToBe(nil)
		}()
		help := bytes.NewBuffer(nil)
		var animated *bool
		fset := NewFlagSet("", ContinueOnError)
		fset.Output = help
		fset.BoolVar(animated, "animated", false, "GIF画像でタイムラプス表示").Alias("a")
	})
}
