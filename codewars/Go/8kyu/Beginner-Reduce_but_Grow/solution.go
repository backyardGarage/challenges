package kata

func Grow(arr []int) int {
	var v int = 1
	for i := 0; i <= len(arr)-1; i++ {
		v *= arr[i]
	}
	return v
}
