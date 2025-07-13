package defaults

import (
	"fmt"
	"reflect"

	check "gopkg.in/check.v1"
)

type FillerSuite struct{}

var _ = check.Suite(&FillerSuite{})

type FixtureTypeInt int

func (FillerSuite) TestFuncByNameIsEmpty(c *check.C) {
	calledA := false
	calledB := false

	f := &Filler{
		FuncByName: map[string]FillerFunc{
			"Foo": func(_ *FieldData) {
				calledA = true
			},
		},
		FuncByKind: map[reflect.Kind]FillerFunc{
			reflect.Int: func(_ *FieldData) {
				calledB = true
			},
		},
	}

	f.Fill(&struct{ Foo int }{})
	c.Assert(calledA, check.Equals, true)
	c.Assert(calledB, check.Equals, false)
}

func (FillerSuite) TestFuncByTypeIsEmpty(c *check.C) {
	calledA := false
	calledB := false

	t := GetTypeHash(reflect.TypeOf(new(FixtureTypeInt)))
	f := &Filler{
		FuncByType: map[TypeHash]FillerFunc{
			t: func(_ *FieldData) {
				calledA = true
			},
		},
		FuncByKind: map[reflect.Kind]FillerFunc{
			reflect.Int: func(_ *FieldData) {
				calledB = true
			},
		},
	}

	f.Fill(&struct{ Foo FixtureTypeInt }{})
	c.Assert(calledA, check.Equals, true)
	c.Assert(calledB, check.Equals, false)
}

func (FillerSuite) TestFuncByKindIsNotEmpty(c *check.C) {
	called := false
	f := &Filler{FuncByKind: map[reflect.Kind]FillerFunc{
		reflect.Int: func(_ *FieldData) {
			called = true
		},
	}}

	f.Fill(&struct{ Foo int }{Foo: 42})
	c.Assert(called, check.Equals, false)
}

func (FillerSuite) TestFuncByKindSlice(_ *check.C) {
	fmt.Println(GetTypeHash(reflect.TypeOf(new([]string))))
}

func (FillerSuite) TestFuncByKindTag(c *check.C) {
	var called string
	f := &Filler{Tag: "foo", FuncByKind: map[reflect.Kind]FillerFunc{
		reflect.Int: func(field *FieldData) {
			called = field.TagValue
		},
	}}

	f.Fill(&struct {
		Foo int `foo:"qux"`
	}{})
	c.Assert(called, check.Equals, "qux")
}

func (FillerSuite) TestFuncByKindIsEmpty(c *check.C) {
	called := false
	f := &Filler{FuncByKind: map[reflect.Kind]FillerFunc{
		reflect.Int: func(_ *FieldData) {
			called = true
		},
	}}

	f.Fill(&struct{ Foo int }{})
	c.Assert(called, check.Equals, true)
}
