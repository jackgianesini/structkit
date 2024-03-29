package structkit

import (
	"reflect"
	"strings"
)

type cp struct {
	fieldsFocused []string
	source        reflect.Value
	destination   []reflect.StructField
	mapping       []reflect.Value
	bypass        map[string]any
	slice         bool
	ptr           bool
}

func Copy(source any, fields ...string) any {
	m := new(cp)

	m.source = reflect.ValueOf(source)
	if !m.source.IsValid() {
		return source
	}

	m.ptr = m.source.Kind() == reflect.Ptr
	if m.ptr {
		m.source = m.source.Elem()
	}

	if m.source.Kind() == reflect.Slice && m.source.Type().Elem().Kind() != reflect.Struct {
		return source
	}

	if m.source.Kind() != reflect.Struct && m.source.Kind() != reflect.Slice {
		return source
	}

	if len(fields) == 0 {
		return source
	}

	m.fieldsFocused = fields
	m.bypass = make(map[string]any)
	m.slice = m.source.Kind() == reflect.Slice

	return m.parse()
}

func (m *cp) extract(parent string) (res []string) {
	for _, value := range m.fieldsFocused {
		s := strings.Split(value, ".")
		if parent == s[0] {
			res = append(res, strings.Join(s[1:], "."))
		}
	}
	return
}

func (m *cp) buildField(field string, parent bool) {
	field = strings.TrimSpace(field)
	if _, ok := m.bypass[field]; ok {
		return
	}

	fieldType, found := m.source.Type().FieldByName(field)
	if !found {
		return
	}

	m.bypass[field] = field

	var newStructField reflect.StructField
	if parent {
		result := Copy(m.source.FieldByName(field).Interface(), m.extract(field)...)
		res := reflect.ValueOf(result)
		newStructField = reflect.StructField{
			Name:      fieldType.Name,
			Type:      res.Type(),
			Tag:       fieldType.Tag,
			Anonymous: fieldType.Anonymous,
		}
		m.destination = append(m.destination, newStructField)
		m.mapping = append(m.mapping, res)
		return
	}

	newStructField = reflect.StructField{
		Name:      fieldType.Name,
		Type:      m.source.FieldByName(field).Type(),
		Tag:       fieldType.Tag,
		Anonymous: fieldType.Anonymous,
	}

	m.destination = append(m.destination, newStructField)
	m.mapping = append(m.mapping, m.source.FieldByName(field))
}

func (m *cp) parse() any {
	if m.slice {
		var data any
		data = m.parseSlice()
		if len(data.([]any)) == 0 {
			data = []string{}
		}

		return data
	}

	return m.parseStruct()
}

func (m *cp) parseSlice() []any {
	var slice []any

	for i := 0; i < m.source.Len(); i++ {
		slice = append(slice, Copy(m.source.Index(i).Interface(), m.fieldsFocused...))
	}

	return slice
}

func (m *cp) parseStruct() any {
	for _, field := range m.fieldsFocused {
		splitByPoint := strings.Split(field, ".")
		if len(splitByPoint) > 1 {
			m.buildField(splitByPoint[0], true)
			continue
		}
		m.buildField(field, false)
	}

	st := reflect.StructOf(m.destination)
	newStruct := reflect.New(st).Elem()

	for i := 0; i < newStruct.NumField(); i++ {
		newStruct.Field(i).Set(m.mapping[i])
	}

	if m.ptr {
		return newStruct.Addr().Interface()
	}

	return newStruct.Interface()
}
