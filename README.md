# largo - Yet another arg parser of Go in a flexible way

[![Go](https://github.com/otiai10/largo/actions/workflows/go.yaml/badge.svg)](https://github.com/otiai10/largo/actions/workflows/go.yaml)
[![codecov](https://codecov.io/gh/otiai10/largo/branch/main/graph/badge.svg?token=OrcqSORFpr)](https://codecov.io/gh/otiai10/largo)
[![Go Report Card](https://goreportcard.com/badge/github.com/otiai10/largo)](https://goreportcard.com/report/github.com/otiai10/largo)
[![Maintainability](https://api.codeclimate.com/v1/badges/e88c2bb92082919b46c0/maintainability)](https://codeclimate.com/github/otiai10/largo/maintainability)
[![Go Reference](https://pkg.go.dev/badge/github.com/otiai10/largo.svg)](https://pkg.go.dev/github.com/otiai10/largo)

# Motivation

```shell
# Let's say your are building `greet`, and can receive any command arg as message.
% greet hello
hello
% greet thanks
thanks
```

What if `greet` can accept some flags,
but would receive **unordered** way like below.

```shell
# 1) BSD style
% greet -count 3 -upper hello
# 2) GNU style
% greet hello -count 3 -upper
# 3) Mixed...
% greet -count 3 hello -upper

# All output should be
HELLO HELLO HELLO
```

The problem is that the standard `flag` package ignores all flags of `case (2)` and `-upper` flag of `case (3)`.
How can we keep flexibility to parse those flags, even if they are NOT urdered in BSD-style?

# Idea

Basically `flag.FlagSet` works fine, however, we like to keep parsing all given flags until the end of line, and retrieve all remaining args as `Rest()`.

```
greet -count 3 hello -upper
-----
 cmd
      --------       ------
      int flag      bool flag
               -----
               rest
```

# Usage

```go
import (
  "github.com/otiai10/largo"
)

var (
  count int
  upper bool
)

func main() {
  fset := largo.NewFlaggSet("greet")
  fset.IntVar(&count, "count", 1, "Number of count to say it").Alias("c")
  fset.BoolVar(&upper, "upper", false, "Say it in upper cases").Alias("u")

  // In most cases, it is given as os.Args[1:] or something.
  fset.Parse([]string{"greet", "-c", "3", "hello", "-upper"})

  fmt.Println("count:", count)   // count: 3
  fmt.Println("upper:", upper)   // upper: true
  fmt.Println("rest:", f.Rest()) // rest: [hello]
}
```

# Reference projects using `largo`

- https://github.com/otiai10/amesh-bot

# Issues & Feedbacks

Any feedbacks will be welcomed!

- https://github.com/otiai10/largo/issues
