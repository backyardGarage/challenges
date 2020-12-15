package kata

func ExpressionMatter(a int, b int, c int) int {
	var arr [3]int = [3]int{a, b, c}
	var max int = 0

	for i := 0; i < 3; i++ {
		var count int = 0
		if arr[i] == 1 {
			count := count + 1
			return count
		}
		if count == 3 {
			max = a + b + c
			return max
		}

		if count == 2 || count == 1 {
			if (a+b)*c-(a+(b*c)) >= 0 {
				return (a + b) * c
			}
		} else {
			return a * (b + c)
		}

		if count == 0 {
			max = a * b * c
			return max
		}
	}
	return max
}
