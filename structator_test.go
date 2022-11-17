package structator

import (
	"github.com/stretchr/testify/suite"
	"log"
	"reflect"
	"testing"
)

type Foo struct {
	Value   string `json:"value"`
	Struct  Bar    `json:"struct"`
	Slice   []Bar
	Counter int
}

type Bar struct {
	Value string `json:"value"`
}

type StructatorTestSuite struct {
	suite.Suite
}

func (test *StructatorTestSuite) TestWithStruct() {
	payload := Foo{Value: "foo", Struct: Bar{Value: "bar"}}
	result := From(payload, "Value", "Value", "Struct.Bar", "NotFound")

	val := reflect.ValueOf(result)
	test.Equal(val.NumField(), 2)
	test.True(val.FieldByName("Value").IsValid())
	test.False(val.FieldByName("Message").IsValid())
	test.False(val.FieldByName("Counter").IsValid())

	log.Printf("%+v", result)
}

func (test *StructatorTestSuite) TestWithPtrStruct() {
	payload := Foo{Value: "foo", Struct: Bar{Value: "bar"}}
	result := From(&payload, "Value", "Value", "Struct.Bar", "NotFound")

	val := reflect.ValueOf(result)
	test.Equal(val.Elem().NumField(), 2)
	test.True(val.Elem().FieldByName("Value").IsValid())
	test.False(val.Elem().FieldByName("Message").IsValid())
	test.False(val.Elem().FieldByName("Counter").IsValid())

	log.Printf("%+v", result)
}

func (test *StructatorTestSuite) TestWithSlice() {
	payload := []Foo{{Value: "foo", Struct: Bar{Value: "bar"}}}
	result := From(payload, "Value", "Value", "Struct.Bar", "Slice.Value")

	rf := reflect.ValueOf(result)
	test.Equal(rf.Index(0).Elem().FieldByName("Value").Interface(), "foo")
	test.False(rf.Index(0).Elem().FieldByName("Counter").IsValid())
	test.Equal(rf.Len(), 1)
}

func (test *StructatorTestSuite) TestWithStructInvalid() {
	test.Nil(From(nil))
	test.Equal(From("hello", "invalid"), "hello")
	test.Equal(From([]string{"hello"}, "invalid"), []string{"hello"})
}

func (test *StructatorTestSuite) TestWithNoField() {
	test.Equal(Foo{}, From(Foo{}))
}

func TestStructatorTestSuite(t *testing.T) {
	suite.Run(t, new(StructatorTestSuite))
}
