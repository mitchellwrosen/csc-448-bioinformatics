package main

import (
	"testing"
)

func TestNucleotideComplement(t *testing.T) {
	nuc := ADENINE
	if nuc.complement() != THYMINE {
		t.Errorf("%c -> %c", nuc, nuc.complement())
	}

	nuc = CYTOSINE
	if nuc.complement() != GUANINE {
		t.Errorf("%c -> %c", nuc, nuc.complement())
	}

	nuc = GUANINE
	if nuc.complement() != CYTOSINE {
		t.Errorf("%c -> %c", nuc, nuc.complement())
	}

	nuc = THYMINE
	if nuc.complement() != ADENINE {
		t.Errorf("%c -> %c", nuc, nuc.complement())
	}
}
