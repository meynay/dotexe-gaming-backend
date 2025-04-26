package pkg

import (
	"strings"
	"unicode"
)

func Exists[T comparable](t T, slice []T) bool {
	for _, t1 := range slice {
		if t == t1 {
			return true
		}
	}
	return false
}

func CalculateScore(query, term string) float64 {
	query = strings.ToLower(strings.TrimSpace(query))
	term = strings.ToLower(strings.TrimSpace(term))
	if query == term {
		return 1
	}
	termOverlap := termOverlapScore(query, term)
	prefixScore := prefixMatchScore(query, term)
	substringScore := substringMatchScore(query, term)
	editDistanceScore := editDistanceScore(query, term)

	totalScore := 0.4*termOverlap +
		0.3*prefixScore +
		0.2*substringScore +
		0.1*editDistanceScore

	return clamp(totalScore, 0, 1)
}

func termOverlapScore(query, product string) float64 {
	qTerms := tokenize(query)
	pTerms := tokenize(product)

	if len(qTerms) == 0 {
		return 0
	}

	matched := 0
	for _, qt := range qTerms {
		for _, pt := range pTerms {
			if qt == pt {
				matched++
				break
			}
		}
	}

	return float64(matched) / float64(len(qTerms))
}

func prefixMatchScore(query, product string) float64 {
	qTerms := tokenize(query)
	pTerms := tokenize(product)

	if len(qTerms) == 0 || len(pTerms) == 0 {
		return 0
	}

	matched := 0
	for _, qt := range qTerms {
		for _, pt := range pTerms {
			if strings.HasPrefix(pt, qt) {
				matched++
				break
			}
		}
	}

	return float64(matched) / float64(len(qTerms))
}

func substringMatchScore(query, product string) float64 {
	qTerms := tokenize(query)
	if len(qTerms) == 0 {
		return 0
	}

	matched := 0
	for _, qt := range qTerms {
		if strings.Contains(product, qt) {
			matched++
		}
	}

	return float64(matched) / float64(len(qTerms))
}

func editDistanceScore(query, product string) float64 {
	maxLen := max(len(query), len(product))
	if maxLen == 0 {
		return 0
	}

	distance := levenshteinDistance(query, product)
	return 1.0 - float64(distance)/float64(maxLen)
}

func tokenize(text string) []string {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	return strings.FieldsFunc(strings.ToLower(text), f)
}

func levenshteinDistance(s, t string) int {
	m := len(s)
	n := len(t)
	d := make([][]int, m+1)

	for i := range d {
		d[i] = make([]int, n+1)
		d[i][0] = i
	}

	for j := range d[0] {
		d[0][j] = j
	}

	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				d[i][j] = min(
					d[i-1][j]+1,
					d[i][j-1]+1,
					d[i-1][j-1]+1,
				)
			}
		}
	}

	return d[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
