package biology

// Amino acid names
const (
	Alanine = iota
	Arginine
	Asparagine
	AsparticAcid
	Cysteine
	GlutamicAcid
	Glutamine
	Glycine
	Histidine
	Isoleucine
	Leucine
	Lysine
	Methionine
	Phenylalanine
	Proline
	Serine
	Threonine
	Tryptophan
	Tyrosine
	Valine
)

type AminoAcid struct {
	Name         string
	Code3        string
	Code1        byte
	Degeneration int
}

func NewAminoAcid(seq Sequence) AminoAcid {
	switch seq.String() {
	case "GCT", "GCC", "GCA", "GCG":
		return AminoAcid{"Alanine", "Ala", 'A', 4}
	case "CGT", "CGC", "CGA", "CGG", "AGA", "AGG":
		return AminoAcid{"Arginine", "Arg", 'R', 6}
	case "AAT", "AAC":
		return AminoAcid{"Asparagine", "Asn", 'N', 2}
	case "GAT", "GAC":
		return AminoAcid{"Aspartic acid", "Asp", 'D', 2}
	case "TGT", "TGC":
		return AminoAcid{"Cysteine", "Cys", 'C', 2}
	case "GAA", "GAG":
		return AminoAcid{"Glutamic acid", "Glu", 'E', 2}
	case "CAA", "CAG":
		return AminoAcid{"Glutamine", "Gln", 'Q', 2}
	case "GGT", "GGC", "GGA", "GGG":
		return AminoAcid{"Glycine", "Gly", 'G', 4}
	case "CAT", "CAC":
		return AminoAcid{"Histidine", "His", 'H', 2}
	case "ATT", "ATC", "ATA":
		return AminoAcid{"Isoleucine", "Ile", 'I', 3}
	case "TTA", "TTG", "CTT", "CTC", "CTA", "CTG":
		return AminoAcid{"Leucine", "Leu", 'L', 6}
	case "AAA", "AAG":
		return AminoAcid{"Lysine", "Lys", 'K', 2}
	case "ATG":
		return AminoAcid{"Methionine", "Met", 'M', 1}
	case "TTT", "TTC":
		return AminoAcid{"Phenylalanine", "Phe", 'F', 2}
	case "CCT", "CCC", "CCA", "CCG":
		return AminoAcid{"Proline", "Pro", 'P', 4}
	case "TCT", "TCC", "TCA", "TCG", "AGT", "AGC":
		return AminoAcid{"Serine", "Ser", 'S', 6}
	case "TAA", "TAG", "TGA":
		return AminoAcid{"STOP", "Sto", '#', 3}
	case "ACT", "ACC", "ACA", "ACG":
		return AminoAcid{"Threonine", "Thr", 'T', 4}
	case "TGG":
		return AminoAcid{"Tryptophan", "Trp", 'W', 1}
	case "TAT", "TAC":
		return AminoAcid{"Tyrosine", "Tyr", 'Y', 2}
	case "GTT", "GTC", "GTA", "GTG":
		return AminoAcid{"Valine", "Val", 'V', 4}
	}

	panic("NOTREACHED")
}
