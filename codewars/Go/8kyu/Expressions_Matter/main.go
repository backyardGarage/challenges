package kata

func ExpressionMatter(a int, b int, c int) int {
	var x1 int = a * b * c
	var x2 int = a + b + c
	var x3 int = a*b + c
	var x4 int = a * (b + c)
	var x5 int = a + (b * c)
	var x6 int = (a + b) * c

	var arr [6]int = [6]int{x1, x2, x3, x4, x5, x6}
	var i, j int
	var key int

	for i = 1; i < 6; i++ {
		key = arr[i]

		for j = i - 1; j >= 0; j-- {
			if arr[j] > key {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = key
	}
	return arr[5]
}
