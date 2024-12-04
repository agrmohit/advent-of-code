// Package mathutils contains commonly used mathematical functions
package mathutils

// Abs takes an int and returns the absolute value
func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
