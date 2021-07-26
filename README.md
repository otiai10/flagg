# WIP: largo

[![Go](https://github.com/otiai10/largo/actions/workflows/go.yaml/badge.svg)](https://github.com/otiai10/largo/actions/workflows/go.yaml)
[![codecov](https://codecov.io/gh/otiai10/largo/branch/main/graph/badge.svg?token=OrcqSORFpr)](https://codecov.io/gh/otiai10/largo)

Yet another command line flag parser, for unordered options.

# Motivation

```sh
// Given your command is `say`, and can receive any command arg as message.
% say hello
hello
% say thanks
thanks
```

What if `say` can accept some flags, but **unordered**.

```sh
// 1) BSD style
% say -n 3 -upper hello
// 2) GNU style
% say hello -n 3 -upper
// 3) Mixed...
% say -n 3 hello -upper

// All output should be
HELLO HELLO HELLO
```

# Idea

```go
import (
  "github.com/otiai10/largo"
)

var (
  count int
  upper bool
)

func main() {
  f := largo.NewFlaggSet("say")
  f.IntVar(&count, "n", 1, "Number of count to say it")
  f.BoolVar(&upper, "upper", false, "Say it in upper cases")

  f.Parse([]string{"say", "-n", "3", "hello", "-upper"})

  fmt.Println("count:", count) // count: 3
  fmt.Println("upper:", upper) // upper: true
  fmt.Println("rest:", f.Rest()) // rest: [hello]
}
```
