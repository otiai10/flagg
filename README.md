# largo - Yet another arg parser of Go in a loose way

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

```shell
# 1) BSD style
% say -count 3 -upper hello
# 2) GNU style
% say hello -count 3 -upper
# 3) Mixed...
% say -count 3 hello -upper

# All output should be
HELLO HELLO HELLO
```

# Idea

Basically `flag.FlagSet` works fine. The missing piece to enable what described above is `Rest()` func, to get all args which are NOT caught as flag args.

```
say -count 3 hello -upper
___
cmd
    --------       ------
      flag          flag
            _______
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
  f := largo.NewFlaggSet("say")
  f.IntVar(&count, "count", 1, "Number of count to say it").Alias("c")
  f.BoolVar(&upper, "upper", false, "Say it in upper cases").Alias("u")

  f.Parse([]string{"say", "-n", "3", "hello", "-upper"})

  fmt.Println("count:", count) // count: 3
  fmt.Println("upper:", upper) // upper: true
  fmt.Println("rest:", f.Rest()) // rest: [hello]
}
```
