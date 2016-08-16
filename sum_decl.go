// +build !gccgo
// +build amd64 s390x

package vector

//go:noscape
func Sum(v []int16) int16
