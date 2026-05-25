package flagext

import (
	"flag"
	"fmt"
	"reflect"
	"strings"
)

type FlagDefinition struct {
	Long            string
	Short           string
	HelpText        string
	OptionFieldName string
	FlagBinder      FlagBinder
}

// Create single instances to reuse
var (
	StringBinder = &StringFlagBinder{}
	BoolBinder   = &BoolFlagBinder{}
	UintBinder   = &UintFlagBinder{}
)

type FlagBinder interface {
	BindFlag(*FlagDefinition, *flag.FlagSet, any) error
}

// StringFlagBinder binds string flags to struct fields

type StringFlagBinder struct{}

func (sb *StringFlagBinder) BindFlag(fd *FlagDefinition, fs *flag.FlagSet, target any) error {
	targetVal := reflect.ValueOf(target).Elem()
	field := targetVal.FieldByName(fd.OptionFieldName)
	if !field.IsValid() {
		return fmt.Errorf("field %s not found", fd.OptionFieldName)
	}

	// Ensure it's initialized
	if field.IsNil() {
		field.Set(reflect.ValueOf(new(string)))
	}

	fs.StringVar(field.Interface().(*string), fd.Long, "", fd.HelpText)
	return nil
}

// BoolFlagBinder binds boolean flags to struct fields

type BoolFlagBinder struct{}

func (bb *BoolFlagBinder) BindFlag(fd *FlagDefinition, fs *flag.FlagSet, target any) error {
	targetVal := reflect.ValueOf(target).Elem()
	field := targetVal.FieldByName(fd.OptionFieldName)
	if !field.IsValid() {
		return fmt.Errorf("field %s not found", fd.OptionFieldName)
	}

	if field.IsNil() {
		field.Set(reflect.ValueOf(new(bool)))
	}

	fs.BoolVar(field.Interface().(*bool), fd.Long, false, fd.HelpText)
	return nil
}

// Reusable UintBinder
type UintFlagBinder struct{}

func (ub *UintFlagBinder) BindFlag(fd *FlagDefinition, fs *flag.FlagSet, target any) error {
	targetVal := reflect.ValueOf(target).Elem()
	field := targetVal.FieldByName(fd.OptionFieldName)
	if !field.IsValid() {
		return fmt.Errorf("field %s not found", fd.OptionFieldName)
	}

	if field.IsNil() {
		field.Set(reflect.ValueOf(new(uint)))
	}

	fs.UintVar(field.Interface().(*uint), fd.Long, 0, fd.HelpText)
	return nil
}

func isFlag(arg string) bool {
	return strings.HasPrefix(arg, "-") && arg != "-"
}
