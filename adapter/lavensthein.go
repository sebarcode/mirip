package adapter

import (
	"github.com/samber/lo"
	"modernc.org/mathutil"
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

func (m *Levenshtein) Compare(a, b string) float64 {
	distance, maxLen := m.distance(a, b)
	return 1 - float64(distance)/float64(maxLen)
}

// Distance returns the Levenshtein distance between a and b. Lower distances
// indicate closer matches. A distance of 0 means the strings are identical.
func (m *Levenshtein) Distance(a, b string) int {
	distance, _ := m.distance(a, b)
	return distance
}

func (m *Levenshtein) distance(a, b string) (int, int) {
	runesA, runesB := []rune(a), []rune(b)

	// Check if both terms are empty.
	lenA, lenB := len(runesA), len(runesB)
	if lenA == 0 && lenB == 0 {
		return 0, 0
	}

	// Check if one of the terms is empty.
	maxLen := mathutil.Max(lenA, lenB)
	if lenA == 0 {
		return m.InsertCost * lenB, maxLen
	}
	if lenB == 0 {
		return m.DeleteCost * lenA, maxLen
	}

	// Initialize cost slice.
	prevCol := make([]int, lenB+1)
	for i := 0; i <= lenB; i++ {
		prevCol[i] = i
	}

	// Calculate distance.
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
