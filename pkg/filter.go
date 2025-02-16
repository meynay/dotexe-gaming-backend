package pkg

func Exists[T comparable](t T, slice []T) bool {
	for _, t1 := range slice {
		if t == t1 {
			return true
		}
	}
	return false
}

func CalculateScore(s1, s2 string) float64 {
	return 1
}
