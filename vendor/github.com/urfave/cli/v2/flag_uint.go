package cli

import (
	"flag"
	"fmt"
	"strconv"
)

// TakesValue returns true of the flag takes a value, otherwise false
func (f *UintFlag) TakesValue() bool {
	return true
}

// GetUsage returns the usage string for the flag
func (f *UintFlag) GetUsage() string {
	return f.Usage
}

// Apply populates the flag given the flag set and environment
func (f *UintFlag) Apply(set *flag.FlagSet) error {
	if val, ok := flagFromEnvOrFile(f.EnvVars, f.FilePath); ok {
		if val != "" {
			valInt, err := strconv.ParseUint(val, 0, 64)
			if err != nil {
				return fmt.Errorf("could not parse %q as uint value for flag %s: %s", val, f.Name, err)
			}

			f.Value = uint(valInt)
			f.HasBeenSet = true
		}
	}

	for _, name := range f.Names() {
		if f.Destination != nil {
			set.UintVar(f.Destination, name, f.Value, f.Usage)
			continue
		}
		set.Uint(name, f.Value, f.Usage)
	}

	return nil
}

// GetValue returns the flags value as string representation and an empty
// string if the flag takes no value at all.
func (f *UintFlag) GetValue() string {
	return fmt.Sprintf("%d", f.Value)
}

// GetDefaultText returns the default text for this flag
func (f *UintFlag) GetDefaultText() string {
	if f.DefaultText != "" {
		return f.DefaultText
	}
	return f.GetValue()
}

// GetEnvVars returns the env vars for this flag
func (f *UintFlag) GetEnvVars() []string {
	return f.EnvVars
}

// Get returns the flag’s value in the given Context.
func (f *UintFlag) Get(ctx *Context) uint {
	return ctx.Uint(f.Name)
}

// Uint looks up the value of a local UintFlag, returns
// 0 if not found
func (cCtx *Context) Uint(name string) uint {
	if fs := cCtx.lookupFlagSet(name); fs != nil {
		return lookupUint(name, fs)
	}
	return 0
}

func lookupUint(name string, set *flag.FlagSet) uint {
	f := set.Lookup(name)
	if f != nil {
		parsed, err := strconv.ParseUint(f.Value.String(), 0, 64)
		if err != nil {
			return 0
		}
		return uint(parsed)
	}
	return 0
}
