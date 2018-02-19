// Package gobasics contains basics conepts from the go guide
package gobasics

// NakedReturn his is an example of nakedReturn
func NakedReturn(x, y string, z int) (a, b string, c int) {
	a, b, c = x, y, z
	return
}
