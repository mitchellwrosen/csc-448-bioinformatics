package biology

import (
	. "launchpad.net/gocheck"
	"testing"
)

// Gocheck boiler plate
func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

////////////////////////////////////////////////////////////////////////////////

func (s *S) TestNucleotideComplement(c *C) {
	c.Check(ADENINE.Complement(), Equals, THYMINE)
	c.Check(CYTOSINE.Complement(), Equals, GUANINE)
	c.Check(THYMINE.Complement(), Equals, ADENINE)
	c.Check(GUANINE.Complement(), Equals, CYTOSINE)
}
