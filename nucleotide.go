package main

import (
	"fmt"
)

const (
	ADENINE  Nucleotide = 'A'
	CYTOSINE Nucleotide = 'C'
	GUANINE  Nucleotide = 'G'
	THYMINE  Nucleotide = 'T'
)

type Nucleotide byte

func (n Nucleotide) complement() Nucleotide {
	switch n {
	case ADENINE:
		return THYMINE
	case CYTOSINE:
		return GUANINE
	case GUANINE:
		return CYTOSINE
	case THYMINE:
		return ADENINE
	}

	panic("NOTREACHED")
}

func (n Nucleotide) String() string {
	return fmt.Sprintf("%c", n)
}
