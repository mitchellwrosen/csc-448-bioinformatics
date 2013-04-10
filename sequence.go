package biology

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Sequence []Nucleotide

func SequenceFromFile(filename string) (Sequence, error) {
	var err error

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	s := make(Sequence, 0, 100)
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		partialSeq, err := sequenceFromBytes(line[:len(line)-1])
		if err != nil {
			return nil, err
		}

		s = append(s, partialSeq...)
	}

	return s, nil
}

func sequenceFromBytes(bs []byte) (Sequence, error) {
	s := make(Sequence, len(bs))
	var err error

	for i, b := range bs {
		if s[i], err = NewNucleotide(b); err != nil {
			return nil, err
		}
	}

	return s, nil
}

// Gets the complement of this Sequence.
//
// returns	The complement
func (s Sequence) Complement() Sequence {
	comp := make(Sequence, len(s))
	for i, val := range s {
		comp[i] = val.Complement()
	}

	return comp
}

// Gets the reverse complement of this Sequence.
//
// returns	The reverse complement
func (s Sequence) ReverseComplement() Sequence {
	comp := make(Sequence, len(s))
	length := len(s)
	for i, val := range s {
		comp[length-1-i] = val.Complement()
	}

	return comp
}

// Determines whether this Sequence equals the given Sequence.
//
// param	s2	The Sequence to compare against
//
// returns	Whether or not this Sequence equals |s2|
func (s Sequence) Equals(s2 Sequence) bool {
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

// Determines the six amino acid sequences this Sequence encodes.
//
// returns	The six amino acid sequences this Sequence encodes
func (s Sequence) ToAminoAcidSequences() []AminoAcidSequence {
	acids := make([]AminoAcidSequence, 6)
	revSeq := s.ReverseComplement()

	// Three frames - offset 0, 1, 2
	for i := 0; i < 3; i++ {
		seqLen := (len(s) - i) / 3
		acids[i] = make([]AminoAcid, seqLen)
		for j := 0; j < seqLen; j++ {
			acids[i][j] = NewAminoAcid(s[(j*3)+i : (j*3)+i+3])
		}
	}

	for i := 0; i < 3; i++ {
		seqLen := (len(s) - i) / 3
		acids[i+3] = make([]AminoAcid, seqLen)
		for j := 0; j < seqLen; j++ {
			acids[i+3][j] = NewAminoAcid(revSeq[(j*3)+i : (j*3)+i+3])
		}
	}

	return acids
}

// Gets all the DNA palindromes of length 2 or more that occur in this Sequence,
// where a DNA palindrome is a nucleotide sequence that is equal to its reverse
// complement.
//
// returns	A slice of palindromes (Sequences)
func (s Sequence) findPalindromes() []Sequence {
	palindromes := make([]Sequence, 0, 10)

	// For each Sequence of length >= 2, compare it to its reverse complement
	for i := 2; i <= len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			seq := s[j : j+i]
			revSeq := seq.ReverseComplement()
			if seq.Equals(revSeq) {
				palindromes = append(palindromes, seq)
			}
		}
	}

	return palindromes
}

func (s Sequence) gcContentHistogram(windowSize, shiftLen int) ([]float32, error) {
	if (len(s)-windowSize)%shiftLen != 0 {
		return nil, fmt.Errorf("Cannot create GC-content histogram for a "+
			"sequence of length %d using a window size of %d and a shift "+
			"length of %d", len(s), windowSize, shiftLen)
	}

	l := (len(s)-windowSize)/shiftLen + 1
	gcContent := make([]float32, l)

	for i := range gcContent {
		gcContent[i] = s[i*shiftLen : i*shiftLen+windowSize].gcContent()
	}

	return gcContent, nil
}

func (s Sequence) gcContent() float32 {
	isGC := func(n Nucleotide) bool {
		return n == Guanine || n == Cytosine
	}

	return float32(len(s.Filter(isGC))) / float32(len(s))
}

func (s Sequence) Filter(f func(n Nucleotide) bool) Sequence {
	filtered := make(Sequence, 0, len(s)/2)
	for _, n := range s {
		if f(n) {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

type LevelOfAgreement struct {
	a float32
	b float32
	c float32
	d float32
}

//func levelOfAgreement(ss []Sequence) ([]LevelOfAgreement, error) {
//if len(ss) == 0 {
//return nil, errors.New("Sequence has length of 0")
//}

//l := len(ss[0])
//for _, s := range ss[1:] {
//if len(s) != l {
//return nil, fmt.Errorf("Not all sequences have length %d\n", l)
//}
//}

//levels := make([]LevelOfAgreement, l)
//for _, s := range ss {
//for j := range s {
//switch s[j] {
//case ADENINE:

//case CYTOSINE:
//case GUANINE:
//case THYMINE:
//}
//}
//}
//}

func (s Sequence) String() string {
	b := make([]byte, len(s))
	for i := range s {
		b[i] = byte(s[i])
	}
	return string(b)
}
