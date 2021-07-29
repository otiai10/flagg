package largo

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/otiai10/largo/values"
)

type (
	FlagSet struct {
		Name      string
		errhandle ErrorHandling
		args      []string
		rest      []string
		dict      map[string]*Flag

		// Usage ...
		// FIXME: Write something here
		Usage         func()
		Output        io.Writer
		Description   string
		helpRequested bool
	}
	ErrorHandling int
)

// https://cs.opensource.google/go/go/+/refs/tags/go1.16.6:src/flag/flag.go;l=314-318;drc=refs%2Ftags%2Fgo1.16.6
const (
	ContinueOnError ErrorHandling = iota // Return a descriptive error.
	ExitOnError                          // Call os.Exit(2) or for -h/-help Exit(0).
	PanicOnError                         // Call panic with a descriptive error.
)

var ErrHelp = errors.New("flag: help requested")

func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet {
	return &FlagSet{
		Name:      name,
		errhandle: errorHandling,
	}
}

func (fset *FlagSet) ParseLine(line string) error {
	return fset.Parse(Tokenize(line))
}

func (fset *FlagSet) Parse(arguments []string) error {
	if len(arguments) == 0 {
		return nil
	}
	if arguments[0] == fset.Name {
		arguments = arguments[1:]
	}
	fset.args = arguments
	for i := 0; i < len(arguments); {
		next, err := fset.parseSingle(i)
		if err != nil {
			if err == ErrHelp {
				return nil
			}
			return fset.onError(err)
		}
		i = next
	}
	return nil
}

func (fset *FlagSet) parseSingle(i int) (next int, err error) {
	s := fset.args[i]
	if s[0] != '-' {
		fset.rest = append(fset.rest, s)
		return i + 1, nil
	}
	if len(s) == 1 {
		return len(fset.args), nil // Finish immediate
	}
	numMinuses := 1
	if s[1] == '-' {
		if len(s) == 2 {
			return len(fset.args), nil // Finish immediate
		}
		numMinuses++
	}

	name := s[numMinuses:]
	if name[0] == '-' {
		return i + 1, fmt.Errorf("invalid arg name: %s", name)
	}

	if name == "help" || name == "h" {
		fset.helpRequested = true
		fset.usage()
		return len(fset.args), ErrHelp
	}

	givenByEqual := false
	rawval := ""

	if kvpair := strings.Split(name, "="); len(kvpair) > 1 {
		givenByEqual = true
		name = kvpair[0]
		rawval = kvpair[1]
	}

	f := fset.Lookup(name)
	if f == nil {
		// if fset.strict {
		// 	return i + 1, fmt.Errorf("unkonwn flag: %v", name)
		// } else {
		return i + 1, nil
		// }
	}

	if bv, ok := f.Value.(*values.BoolValue); ok {
		if !givenByEqual {
			rawval = "true"
		}
		return i + 1, bv.Set(rawval)
	}

	if !givenByEqual { // Use next arg and skip it
		rawval = fset.args[i+1]
		i++
	}

	err = f.Value.Set(rawval)

	return i + 1, err
}

func (fset *FlagSet) Rest() []string {
	return fset.rest
}

func (fset *FlagSet) StringVar(dest *string, name string, defaultval string, usage string) *Flag {
	*dest = defaultval
	sv := (*values.StringValue)(dest)
	return fset.Var(sv, name, usage)
}

func (fset *FlagSet) BoolVar(dest *bool, name string, defaultval bool, usage string) *Flag {
	*dest = defaultval
	bv := (*values.BoolValue)(dest)
	return fset.Var(bv, name, usage)
}

func (fset *FlagSet) IntVar(dest *int, name string, defaultval int, usage string) *Flag {
	*dest = defaultval
	iv := (*values.IntValue)(dest)
	return fset.Var(iv, name, usage)
}

func (fset *FlagSet) Var(value Value, name string, usage string) *Flag {
	flag := &Flag{Name: name, Value: value, Usage: usage}
	if fset.dict == nil {
		fset.dict = make(map[string]*Flag)
	}
	fset.dict[name] = flag
	return flag
}

func (fset *FlagSet) Lookup(name string) *Flag {
	found, ok := fset.dict[name]
	if ok {
		return found
	}
	for _, f := range fset.dict {
		for _, a := range f.aliases {
			if a == name {
				return f
			}
		}
	}
	return nil
}

func (fset *FlagSet) onError(err error) error {
	// if err == nil {
	// 	return nil
	// }
	switch fset.errhandle {
	case ContinueOnError:
		return err
	case ExitOnError:
		fmt.Println(err.Error())
		os.Exit(1)
	case PanicOnError:
		panic(err)
	}
	return err
}

func (fset *FlagSet) usage() {
	if fset.Usage != nil {
		fset.Usage()
	}
	fset.printDefaultUsage()
}

func (fset *FlagSet) printDefaultUsage() {
	if fset.Output == nil {
		fset.Output = os.Stderr
	}
	fset.PrintDefaultUsage(fset.Output)
}

func (fset *FlagSet) PrintDefaultUsage(w io.Writer) error {
	tpl := template.Must(template.New("usage").Parse(string(defaultUsageTemplate)))
	return tpl.Execute(w, fset)
}

func (fset *FlagSet) List() []*Flag {
	result := []*Flag{}
	for _, f := range fset.dict {
		result = append(result, f)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result
}

func (fset *FlagSet) HelpRequested() bool {
	return fset.helpRequested
}
