package kata

func OddCount(n int) int {
	var count int

	for i := 0; i < n-1; i = i + 2 {
		count += 1
	}
	return count
}
