// Package withoutinterfce provides a test on pub/private func for checking
// godoc documentation.
package withoutinterfce

// sub.
//
// description will be hidden from the godoc
func sub(a, b int) int {
	return a - b
}

// Diff returns absolute difference between two integer
//
// a private sub function will be used inside
func Diff(a, b int) int {
	diff := sub(a, b)
	if diff > 0 {
		return diff
	}
	return -diff
}
