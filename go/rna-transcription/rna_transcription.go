// Package strand converts dan strands to rna
package strand

var dnaToRnaMap = map[string]string{
	"C": "G",
	"G": "C",
	"T": "A",
	"A": "U",
}

// ToRNA converts a DNA strand to an RNA strand
func ToRNA(dna string) string {
	rna := ""
	for i := 0; i < len(dna); i++ {
		rna += dnaToRnaMap[string(dna[i])]
	}
	return rna
}
