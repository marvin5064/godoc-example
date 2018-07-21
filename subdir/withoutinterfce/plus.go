package withoutinterfce

func sub(a, b int) int {
	return a - b
}

func Diff(a, b int) int {
	diff := sub(a, b)
	if diff > 0 {
		return diff
	}
	return -diff
}
