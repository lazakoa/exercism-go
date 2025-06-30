package strand

func ToRNA(dna string) string {
	var rna string
	for _, char := range dna {
		rna += string(ConvertNucleotide(char))
	}
	return rna
}

func ConvertNucleotide(nucleotide rune) rune {
	if nucleotide == 'G' {
		return 'C'
	} else if nucleotide == 'C' {
		return 'G'
	} else if nucleotide == 'T' {
		return 'A'
	} else {
		return 'U'
	}
}
