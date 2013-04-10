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

func (s *S) TestSequenceComplement(c *C) {
	seq := newSequence('A', 'T', 'C', 'G')
	c.Check(seq.complement().String(), Equals, "TAGC")
	c.Check(seq.reverseComplement().String(), Equals, "CGAT")
}

func (s *S) TestToAminoAcidSequences(c *C) {
	str := "TCTTAATCGAATCGAT"
	seq := Sequence(str)

	acids := seq.toAminoAcidSequences()
	c.Assert(len(acids), Equals, 6)
	c.Check(acids[0].String(), Equals, "S#SNR")
	c.Check(acids[1].String(), Equals, "LNRID")
	c.Check(acids[2].String(), Equals, "LIES")
	c.Check(acids[3].String(), Equals, "IDSIK")
	c.Check(acids[4].String(), Equals, "SIRLR")
	c.Check(acids[5].String(), Equals, "RFD#")
}

func (s *S) TestFindPalindromes(c *C) {
	// TA, CG, GC, CG, TA, CGCG, ACGCGT, TACGCGTA
	str := "TACGCGTA"
	seq := Sequence(str)

	palindromes := seq.findPalindromes()
	c.Check(len(palindromes), Equals, 8)
}
