package main

// Amino acid names
const (
	ALANINE = iota
	ARGININE
	ASPARAGINE
	ASPARTIC_ACID
	CYSTEINE
	GLUTAMIC_ACID
	GLUTAMINE
	GLYCINE
	HISTIDINE
	ISOLEUCINE
	LEUCINE
	LYSINE
	METHIONINE
	PHENYLALANINE
	PROLINE
	SERINE
	THREONINE
	TRYPTOPHAN
	TYROSINE
	VALINE
)

type AminoAcid struct {
	name  string
	code3 string
	code1 byte
}

func newAminoAcid(seq Sequence) AminoAcid {
	switch seq.String() {
	case "TTT":
		return AminoAcid{"Phenylalanine", "Phe", 'F'}
	case "TTC":
		return AminoAcid{"Phenylalanine", "Phe", 'F'}
	case "TTA":
		return AminoAcid{"Leucine", "Leu", 'L'}
	case "TTG":
		return AminoAcid{"Leucine", "Leu", 'L'}

	case "TCT":
		return AminoAcid{"Serine", "Ser", 'S'}
	case "TCC":
		return AminoAcid{"Serine", "Ser", 'S'}
	case "TCA":
		return AminoAcid{"Serine", "Ser", 'S'}
	case "TCG":
		return AminoAcid{"Serine", "Ser", 'S'}

	case "TAT":
		return AminoAcid{"Tyrosine", "Tyr", 'Y'}
	case "TAC":
		return AminoAcid{"Tyrosine", "Tyr", 'Y'}
	case "TAA":
		return AminoAcid{"STOP", "Sto", '#'}
	case "TAG":
		return AminoAcid{"STOP", "Sto", '#'}

	case "TGT":
		return AminoAcid{"Cysteine", "Cys", 'C'}
	case "TGC":
		return AminoAcid{"Cysteine", "Cys", 'C'}
	case "TGA":
		return AminoAcid{"STOP", "Sto", '#'}
	case "TGG":
		return AminoAcid{"Tryptophan", "Trp", 'W'}

	case "CTT":
		return AminoAcid{"Leucine", "Leu", 'L'}
	case "CTC":
		return AminoAcid{"Leucine", "Leu", 'L'}
	case "CTA":
		return AminoAcid{"Leucine", "Leu", 'L'}
	case "CTG":
		return AminoAcid{"Leucine", "Leu", 'L'}

	case "CCT":
		return AminoAcid{"Proline", "Pro", 'P'}
	case "CCC":
		return AminoAcid{"Proline", "Pro", 'P'}
	case "CCA":
		return AminoAcid{"Proline", "Pro", 'P'}
	case "CCG":
		return AminoAcid{"Proline", "Pro", 'P'}

	case "CAT":
		return AminoAcid{"Histidine", "His", 'H'}
	case "CAC":
		return AminoAcid{"Histidine", "His", 'H'}
	case "CAA":
		return AminoAcid{"Glutamine", "Gln", 'Q'}
	case "CAG":
		return AminoAcid{"Glutamine", "Gln", 'Q'}

	case "CGT":
		return AminoAcid{"Arginine", "Arg", 'R'}
	case "CGC":
		return AminoAcid{"Arginine", "Arg", 'R'}
	case "CGA":
		return AminoAcid{"Arginine", "Arg", 'R'}
	case "CGG":
		return AminoAcid{"Arginine", "Arg", 'R'}

	case "ATT":
		return AminoAcid{"Isoleucine", "Ile", 'I'}
	case "ATC":
		return AminoAcid{"Isoleucine", "Ile", 'I'}
	case "ATA":
		return AminoAcid{"Isoleucine", "Ile", 'I'}
	case "ATG":
		return AminoAcid{"Methionine", "Met", 'M'}

	case "ACT":
		return AminoAcid{"Threonine", "Thr", 'T'}
	case "ACC":
		return AminoAcid{"Threonine", "Thr", 'T'}
	case "ACA":
		return AminoAcid{"Threonine", "Thr", 'T'}
	case "ACG":
		return AminoAcid{"Threonine", "Thr", 'T'}

	case "AAT":
		return AminoAcid{"Asparagine", "Asn", 'N'}
	case "AAC":
		return AminoAcid{"Asparagine", "Asn", 'N'}
	case "AAA":
		return AminoAcid{"Lysine", "Lys", 'K'}
	case "AAG":
		return AminoAcid{"Lysine", "Lys", 'K'}

	case "AGT":
		return AminoAcid{"Serine", "Ser", 'S'}
	case "AGC":
		return AminoAcid{"Serine", "Ser", 'S'}
	case "AGA":
		return AminoAcid{"Arginine", "Arg", 'R'}
	case "AGG":
		return AminoAcid{"Arginine", "Arg", 'R'}

	case "GTT":
		return AminoAcid{"Valine", "Val", 'V'}
	case "GTC":
		return AminoAcid{"Valine", "Val", 'V'}
	case "GTA":
		return AminoAcid{"Valine", "Val", 'V'}
	case "GTG":
		return AminoAcid{"Valine", "Val", 'V'}

	case "GCT":
		return AminoAcid{"Alanine", "Ala", 'A'}
	case "GCC":
		return AminoAcid{"Alanine", "Ala", 'A'}
	case "GCA":
		return AminoAcid{"Alanine", "Ala", 'A'}
	case "GCG":
		return AminoAcid{"Alanine", "Ala", 'A'}

	case "GAT":
		return AminoAcid{"Aspartic acid", "Asp", 'D'}
	case "GAC":
		return AminoAcid{"Aspartic acid", "Asp", 'D'}
	case "GAA":
		return AminoAcid{"Glutamic acid", "Glu", 'E'}
	case "GAG":
		return AminoAcid{"Glutamic acid", "Glu", 'E'}

	case "GGT":
		return AminoAcid{"Glycine", "Gly", 'G'}
	case "GGC":
		return AminoAcid{"Glycine", "Gly", 'G'}
	case "GGA":
		return AminoAcid{"Glycine", "Gly", 'G'}
	case "GGG":
		return AminoAcid{"Glycine", "Gly", 'G'}
	}

	panic("NOTREACHED")
}
