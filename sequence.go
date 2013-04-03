package main

import (
//"fmt"
)

type Sequence []Nucleotide

func (s Sequence) complement() Sequence {
	comp := make(Sequence, len(s))
	for i, val := range s {
		comp[i] = val.complement()
	}

	return comp
}

func (s Sequence) reverseComplement() Sequence {
	comp := make(Sequence, len(s))
	length := len(s)
	for i, val := range s {
		comp[length-1-i] = val.complement()
	}

	return comp
}

func (s Sequence) equals(s2 Sequence) bool {
	if len(s) != len(s2) {
		return false
	}

	for i := range s {
		if s[i] != s2[i] {
			return false
		}
	}

	return true
}

func (s Sequence) toAminoAcidSequences() []AminoAcidSequence {
	acids := make([]AminoAcidSequence, 6)
	revSeq := s.reverseComplement()

	// Three frames - offset 0, 1, 2
	for i := 0; i < 3; i++ {
		seqLen := (len(s) - i) / 3
		acids[i] = make([]AminoAcid, seqLen)
		for j := 0; j < seqLen; j++ {
			acids[i][j] = newAminoAcid(s[(j*3)+i : (j*3)+i+3])
		}
	}

	for i := 0; i < 3; i++ {
		seqLen := (len(s) - i) / 3
		acids[i+3] = make([]AminoAcid, seqLen)
		for j := 0; j < seqLen; j++ {
			acids[i+3][j] = newAminoAcid(revSeq[(j*3)+i : (j*3)+i+3])
		}
	}

	return acids
}

func (s Sequence) String() string {
	b := make([]byte, len(s))
	for i := range s {
		b[i] = byte(s[i])
	}
	return string(b)
}

func newSequence(ns ...Nucleotide) Sequence {
	return ns
}
