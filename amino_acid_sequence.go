package main

type AminoAcidSequence []AminoAcid

func (a AminoAcidSequence) String() string {
	chars := make([]byte, len(a))
	for i, acid := range a {
		chars[i] = acid.code1
	}
	return string(chars)
}
