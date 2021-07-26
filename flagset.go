package largo

import (
	"flag"
	"fmt"
	"strings"

	"github.com/otiai10/largo/values"
)

type FlagSet struct {
	Name      string
	errhandle flag.ErrorHandling
	args      []string
	rest      []string
	dict      map[string]*Flag
	strict    bool // Error on Unkonw flag
	// *flag.FlagSet
}

func NewFlagSet(name string, errorHandling flag.ErrorHandling) *FlagSet {
	return &FlagSet{
		Name:      name,
		errhandle: errorHandling,
	}
}

func (fset *FlagSet) Parse(arguments []string) error {
	fset.args = arguments
	for i := 0; i < len(arguments); {
		next, err := fset.parseSingle(i)
		if err != nil {
			return err
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
	givenByEqual := false
	rawval := ""

	if kvpair := strings.Split(name, "="); len(kvpair) > 1 {
		givenByEqual = true
		name = kvpair[0]
		rawval = kvpair[1]
	}

	f, found := fset.dict[name]
	if !found && fset.strict {
		return i + 1, fmt.Errorf("unkonwn flag: %v", name)
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

func (fset *FlagSet) StringVar(dest *string, name string, defaultval string, usage string) {
	*dest = defaultval
	sv := (*values.StringValue)(dest)
	flag := &Flag{Name: name, Value: sv, Usage: usage}
	if fset.dict == nil {
		fset.dict = make(map[string]*Flag)
	}
	fset.dict[name] = flag
}

func (fset *FlagSet) Lookup(name string) *Flag {
	return fset.dict[name]
}
