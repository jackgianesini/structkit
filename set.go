package structkit

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type set struct {
	source     reflect.Value
	kind       reflect.Kind
	value      any
	fieldSplit []string
}

func Set(source any, field string, value any) (err error) {
	set := new(set)

	set.source = reflect.ValueOf(source)
	if !set.source.IsValid() {
		return errors.New("source is invalid")
	}

	if set.source.Kind() != reflect.Ptr {
		return errors.New("source must be a pointer")
	}

	if set.source.IsNil() {
		return errors.New("source is nil")
	}

	set.source = set.source.Elem()
	set.kind = set.source.Kind()
	set.fieldSplit = strings.Split(field, ".")
	set.value = value

	if set.kind != reflect.Struct && set.kind != reflect.Slice {
		return errors.New("source must be a struct or a slice of struct")
	}

	return set.apply()
}

func (set *set) apply() (err error) {
	/*	if set.kind == reflect.Slice {
		return set.applyToSlice()
	}*/
	return set.applyToStruct()
}

/*func (set *set) applyToSlice() (err error) {
	return
}*/

func (set *set) workingFieldName() string {
	return set.fieldSplit[0]
}

func (set *set) lenFieldSplit() int {
	return len(set.fieldSplit)
}

func (set *set) applyToStruct() (err error) {
	workingField := set.source.FieldByName(set.workingFieldName())
	if !workingField.IsValid() {
		return fmt.Errorf("field %s not found", set.workingFieldName())
	}

	if set.lenFieldSplit() > 1 {
		return Set(workingField.Addr().Interface(), strings.Join(set.fieldSplit[1:], "."), set.value)
	}

	value := reflect.ValueOf(set.value)

	if workingField.Kind() != value.Kind() {
		return errors.New("type mismatch")
	}

	workingField.Set(value)

	return
}
