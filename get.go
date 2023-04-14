package structkit

import (
	"reflect"
	"strconv"
	"strings"
)

type get struct {
	fieldSplit []string
	field      string
	opt        *Option
	source     reflect.Value
	sourceType reflect.Type
	ptr        bool
	isSlice    bool
}

type Option struct {
	index      []int
	field      []string
	withCache  bool
	identifier string
	Delimiter  string
}

func Get(source any, field string, opt ...*Option) any {
	get := new(get)

	get.source = reflect.ValueOf(source)
	if !get.source.IsValid() {
		return nil
	}

	get.sourceType = get.source.Type()

	get.ptr = get.source.Kind() == reflect.Ptr

	if get.ptr && get.source.IsNil() {
		return nil
	}

	if get.ptr {
		get.source = get.source.Elem()
		get.sourceType = get.source.Type()
	}

	if get.source.Kind() == reflect.Slice {
		get.isSlice = true
	}

	if get.source.Kind() != reflect.Struct && !get.isSlice {
		return nil
	}

	get.opt = new(Option)
	if len(opt) > 0 {
		get.opt = opt[0]
	}

	if get.opt.Delimiter == "" {
		get.opt.Delimiter = "."
	}

	get.fieldSplit = strings.Split(field, get.Delimiter())
	get.field = field

	if get.opt.identifier == "" {
		get.opt.identifier = field
	}

	if len(get.field) > 1 {
		get.opt.withCache = true
	}

	return get.parse()
}

func (get *get) Identifier() string {
	return get.opt.identifier
}

func (get *get) parse() any {
	if get.isSlice {
		return get.parseSlice()
	}
	return get.parseStruct()
}

func (get *get) parseSlice() any {
	s := get.fieldSplit[0]
	get.opt.field = append(get.opt.field, s)

	if s[0] != '[' {
		return nil
	}
	s = s[1:]

	if s[len(s)-1] != ']' {
		return nil
	}

	s = s[:len(s)-1]

	for _, c := range s {
		if c < '0' || c > '9' {
			return nil
		}
	}

	index, _ := strconv.Atoi(s)
	get.opt.index = append(get.opt.index, index)

	if index >= get.source.Len() {
		return nil
	}

	return Get(get.source.Index(index).Interface(), strings.Join(get.fieldSplit[1:], get.Delimiter()), get.opt)
}

func (get *get) parseStruct() any {
	workingFieldName := get.fieldSplit[0]

	get.opt.field = append(get.opt.field, workingFieldName)

	if get.opt.withCache {
		index := getCache().GetField(get.Identifier(), get.field)
		if getCache().GetIdentifier(get.Identifier()) != nil && len(index) > 0 {
			for _, i := range index {
				if get.source.Kind() == reflect.Slice {
					get.source = get.source.Index(i)
					continue
				}
				if get.source.Kind() == reflect.Ptr {
					get.source = get.source.Elem()
				}
				if !get.source.IsValid() {
					return nil
				}
				get.source = get.source.Field(i)
			}
			return get.source.Interface()
		}
	}

	workingField := get.source.FieldByName(workingFieldName)
	if !workingField.IsValid() {
		return nil
	}

	if get.opt.withCache {
		field, _ := get.sourceType.FieldByName(workingFieldName)
		get.opt.index = append(get.opt.index, field.Index[0])
		getCache().SetField(get.Identifier(), strings.Join(get.opt.field, get.Delimiter()), get.opt.index)
	}

	if len(get.fieldSplit) > 1 {
		return Get(workingField.Interface(), strings.Join(get.fieldSplit[1:], get.Delimiter()), get.opt)
	}

	return workingField.Interface()
}

func (get *get) Delimiter() string {
	return get.opt.Delimiter
}
