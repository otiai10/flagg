package largo

import "fmt"

func ExampleFlag_Alias() {
	var verbose bool
	fset := NewFlagSet("foo", ContinueOnError)
	fset.BoolVar(&verbose, "verbose", false, "Show verbose log").Alias("V")
	fset.Parse([]string{"baa", "-V", "baz"})
	fmt.Println("verbose?", verbose)
	// Output:
	// verbose? true
}

func ExampleFlagSet_Rest() {
	var upper bool
	var count int
	fset := NewFlagSet("say", ContinueOnError)
	fset.BoolVar(&upper, "upper", false, "Make chars uppercase").Alias("U")
	fset.IntVar(&count, "count", 1, "Count to say that word").Alias("c")

	// In most cases, it comes from `os.Args`.
	args := []string{"say", "-count", "3", "hello", "-upper", "hiromu"}

	fset.Parse(args)
	fmt.Printf("upper: %v\ncount: %v\nrest: %v\n", upper, count, fset.Rest())
	// Output:
	// upper: true
	// count: 3
	// rest: [hello hiromu]
}
