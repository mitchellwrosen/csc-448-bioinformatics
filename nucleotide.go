package biology

import (
	"fmt"
)

const (
	Adenine  Nucleotide = 'A'
	Cytosine Nucleotide = 'C'
	Guanine  Nucleotide = 'G'
	Thymine  Nucleotide = 'T'
)

type Nucleotide byte

func NewNucleotide(b byte) (Nucleotide, error) {
	switch b {
	case 'A', 'a':
		return Adenine, nil
	case 'C', 'c':
		return Cytosine, nil
	case 'G', 'g':
		return Guanine, nil
	case 'T', 't':
		return Thymine, nil
	}

	return Nucleotide(b), fmt.Errorf("Invalid nucleotide %c", b)
}

func (n Nucleotide) Complement() Nucleotide {
	switch n {
	case Adenine:
		return Thymine
	case Cytosine:
		return Guanine
	case Guanine:
		return Cytosine
	case Thymine:
		return Adenine
	}

	panic("NOTREACHED")
}

func (n Nucleotide) String() string {
	return fmt.Sprintf("%c", n)
}
