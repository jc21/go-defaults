package defaults

import check "gopkg.in/check.v1"

type FactorySuite struct{}

var _ = check.Suite(&FactorySuite{})

func (s *FactorySuite) TestSetDefaultsBasic(c *check.C) {
	foo := &ExampleBasic{}
	Factory(foo)

	s.assertTypes(c, foo)
}

func (FactorySuite) assertTypes(c *check.C, foo *ExampleBasic) {
	c.Assert(foo.String, check.HasLen, 32)
	c.Assert(foo.Integer, check.Not(check.Equals), 0)
	c.Assert(foo.Integer8, check.Not(check.Equals), int8(0))
	c.Assert(foo.Integer16, check.Not(check.Equals), int16(0))
	c.Assert(foo.Integer32, check.Not(check.Equals), int32(0))
	c.Assert(foo.Integer64, check.Not(check.Equals), int64(0))
	c.Assert(foo.UInteger, check.Not(check.Equals), uint(0))
	c.Assert(foo.UInteger8, check.Not(check.Equals), uint8(0))
	c.Assert(foo.UInteger16, check.Not(check.Equals), uint16(0))
	c.Assert(foo.UInteger32, check.Not(check.Equals), uint32(0))
	c.Assert(foo.UInteger64, check.Not(check.Equals), uint64(0))
	c.Assert(foo.String, check.Not(check.Equals), "")
	c.Assert(string(foo.Bytes), check.HasLen, 32)
	c.Assert(foo.Float32, check.Not(check.Equals), float32(0))
	c.Assert(foo.Float64, check.Not(check.Equals), float64(0))
}
