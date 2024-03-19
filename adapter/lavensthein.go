package adapter

import (
	"github.com/samber/lo"
)

type Levenshtein struct {
	InsertCost  int
	DeleteCost  int
	ReplaceCost int
}

func NewLevenshtein() *Levenshtein {
	return &Levenshtein{
		InsertCost:  1,
		DeleteCost:  1,
		ReplaceCost: 1,
	}
}

// Compare menghasilkan angka antara 0 - 1, dimana 1 = identik
func (m *Levenshtein) Compare(a, b string) float64 {
	distance, maxLen := m.distance(a, b)
	return 1 - float64(distance)/float64(maxLen)
}

// Distance menghitung jarak levenstein
func (m *Levenshtein) Distance(a, b string) int {
	distance, _ := m.distance(a, b)
	return distance
}

func (m *Levenshtein) distance(a, b string) (int, int) {
	runesA, runesB := []rune(a), []rune(b)

	lenA, lenB := len(runesA), len(runesB)
	if lenA == 0 && lenB == 0 {
		return 0, 0
	}

	maxLen := lo.Max([]int{lenA, lenB})
	if lenA == 0 {
		return m.InsertCost * lenB, maxLen
	}
	if lenB == 0 {
		return m.DeleteCost * lenA, maxLen
	}

	prevCol := make([]int, lenB+1)
	for i := 0; i <= lenB; i++ {
		prevCol[i] = i
	}

	col := make([]int, lenB+1)
	for i := 0; i < lenA; i++ {
		col[0] = i + 1
		for j := 0; j < lenB; j++ {
			delCost := prevCol[j+1] + m.DeleteCost
			insCost := col[j] + m.InsertCost

			subCost := prevCol[j]
			if runesA[i] != runesB[j] {
				subCost += m.ReplaceCost
			}

			col[j+1] = lo.Min([]int{delCost, insCost, subCost})
		}

		col, prevCol = prevCol, col
	}

	return prevCol[lenB], maxLen
}
