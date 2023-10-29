package piscine

func TrimAtoi(s string) int {
	values := []rune(s)
	var arr []int
	t := 1

	for _, i := range values {

		if i <= '9' && i >= '0' {
			arr = append(arr, int(i-'0'))
		}
		if arr == nil && i == '-' {
			t = -1
		}
	}
	if len(arr) == 0 {
		return 0
	}
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = result*10 + arr[i]
	}
	return result * t
}
