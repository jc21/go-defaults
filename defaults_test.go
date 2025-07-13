package defaults

import (
	"testing"
	"time"

	"bou.ke/monkey"
	check "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.

func Test(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		t, _ := time.Parse("2006-01-02 15:04:05", "2020-06-10 12:00:00")
		return t
	})

	check.TestingT(t)
}

type DefaultsSuite struct{}

var _ = check.Suite(&DefaultsSuite{})

type Parent struct {
	Children []Child
}

type Child struct {
	Name string
	Age  int `default:"10"`
}

type ExampleBasic struct {
	Bool       bool    `default:"true"`
	Integer    int     `default:"33"`
	Integer8   int8    `default:"8"`
	Integer16  int16   `default:"16"`
	Integer32  int32   `default:"32"`
	Integer64  int64   `default:"64"`
	UInteger   uint    `default:"11"`
	UInteger8  uint8   `default:"18"`
	UInteger16 uint16  `default:"116"`
	UInteger32 uint32  `default:"132"`
	UInteger64 uint64  `default:"164"`
	String     string  `default:"foo"`
	Bytes      []byte  `default:"bar"`
	Float32    float32 `default:"3.2"`
	Float64    float64 `default:"6.4"`
	// nolint: revive
	Struct struct {
		Bool    bool `default:"true"`
		Integer int  `default:"33"`
	}
	Duration          time.Duration `default:"3s"`
	Children          []Child
	Second            time.Duration `default:"2s"`
	Integer64Duration time.Duration `default:"6000000000"`
	StringSlice       []string      `default:"[1,2,3,4]"`
	IntSlice          []int         `default:"[1,2,3,4]"`
	IntSliceSlice     [][]int       `default:"[[1],[2],[3],[4]]"`
	StringSliceSlice  [][]string    `default:"[[1],[]]"`

	DateTime string `default:"{{date:1,-10,0}} {{time:1,-5,10}}"`
}

func (s *DefaultsSuite) TestSetDefaultsBasic(c *check.C) {
	foo := &ExampleBasic{}
	SetDefaults(foo)
	s.assertTypes(c, foo)
}

type ExampleNested struct {
	Struct ExampleBasic
}

func (s *DefaultsSuite) TestSetDefaultsNested(c *check.C) {
	foo := &ExampleNested{}
	SetDefaults(foo)
	s.assertTypes(c, &foo.Struct)
}

func (DefaultsSuite) assertTypes(c *check.C, foo *ExampleBasic) {
	c.Assert(foo.Bool, check.Equals, true)
	c.Assert(foo.Integer, check.Equals, 33)
	c.Assert(foo.Integer8, check.Equals, int8(8))
	c.Assert(foo.Integer16, check.Equals, int16(16))
	c.Assert(foo.Integer32, check.Equals, int32(32))
	c.Assert(foo.Integer64, check.Equals, int64(64))
	c.Assert(foo.UInteger, check.Equals, uint(11))
	c.Assert(foo.UInteger8, check.Equals, uint8(18))
	c.Assert(foo.UInteger16, check.Equals, uint16(116))
	c.Assert(foo.UInteger32, check.Equals, uint32(132))
	c.Assert(foo.UInteger64, check.Equals, uint64(164))
	c.Assert(foo.String, check.Equals, "foo")
	c.Assert(string(foo.Bytes), check.Equals, "bar")
	c.Assert(foo.Float32, check.Equals, float32(3.2))
	c.Assert(foo.Float64, check.Equals, 6.4)
	c.Assert(foo.Struct.Bool, check.Equals, true)
	c.Assert(foo.Duration, check.Equals, 3*time.Second)
	c.Assert(foo.Children, check.IsNil)
	c.Assert(foo.Second, check.Equals, 2*time.Second)
	c.Assert(foo.Integer64Duration, check.Equals, 6*time.Second)
	c.Assert(foo.StringSlice, check.DeepEquals, []string{"1", "2", "3", "4"})
	c.Assert(foo.IntSlice, check.DeepEquals, []int{1, 2, 3, 4})
	c.Assert(foo.IntSliceSlice, check.DeepEquals, [][]int{{1}, {2}, {3}, {4}})
	c.Assert(foo.StringSliceSlice, check.DeepEquals, [][]string{{"1"}, {}})
	c.Assert(foo.DateTime, check.Equals, "2020-08-10 12:55:10")
}

func (DefaultsSuite) TestSetDefaultsWithValues(c *check.C) {
	foo := &ExampleBasic{
		Integer:  55,
		UInteger: 22,
		Float32:  9.9,
		String:   "bar",
		Bytes:    []byte("foo"),
		Children: []Child{{Name: "alice"}, {Name: "bob", Age: 2}},
	}

	SetDefaults(foo)

	c.Assert(foo.Integer, check.Equals, 55)
	c.Assert(foo.UInteger, check.Equals, uint(22))
	c.Assert(foo.Float32, check.Equals, float32(9.9))
	c.Assert(foo.String, check.Equals, "bar")
	c.Assert(string(foo.Bytes), check.Equals, "foo")
	c.Assert(foo.Children[0].Age, check.Equals, 10)
	c.Assert(foo.Children[1].Age, check.Equals, 2)
}

func (DefaultsSuite) BenchmarkLogic(c *check.C) {
	for i := 0; i < c.N; i++ {
		foo := &ExampleBasic{}
		SetDefaults(foo)
	}
}
