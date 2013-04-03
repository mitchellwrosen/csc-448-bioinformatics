package main

import (
	"testing"
)

func TestSequenceComplement(t *testing.T) {
	seq := newSequence('A', 'T', 'C', 'G')

	comp := seq.complement()
	if comp.String() != "TAGC" {
		t.Errorf("%s complement -> %s (expected TAGC)", seq, comp)
	}

	revComp := seq.reverseComplement()
	if revComp.String() != "CGAT" {
		t.Errorf("%s reverse complement -> %s (expected CGAT)", seq, revComp)
	}
}

func TestToAminoAcidSequences(t *testing.T) {
	str := "TCTTAATCGAATCGAT"
	seq := Sequence(str)

	acids := seq.toAminoAcidSequences()
	if len(acids) != 6 {
		t.Fatalf("%d amino acid sequences (expected 6)", len(acids))
	}

	if acids[0].String() != "S#SNR" {
		t.Errorf("%s frame 0: %s (expected S#SNR)", str, acids[0].String())
	}
	if acids[1].String() != "LNRID" {
		t.Errorf("%s frame 1: %s (expected LNRID)", str, acids[1].String())
	}
	if acids[2].String() != "LIES" {
		t.Errorf("%s frame 2: %s (expected LIES)", str, acids[2].String())
	}
	if acids[3].String() != "IDSIK" {
		t.Errorf("%s reversed frame 0: %s (expected IDSIK)", str, acids[3].String())
	}
	if acids[4].String() != "SIRLR" {
		t.Errorf("%s reversed frame 1: %s (expected SIRLR)", str, acids[4].String())
	}
	if acids[5].String() != "RFD#" {
		t.Errorf("%s reversed frame 2: %s (expected RFD#)", str, acids[5].String())
	}
}
