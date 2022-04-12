package examples

import (
	"fmt"
	"testing"
)

/**
s[begin:end:cap_end]
*/
func TestSlice(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	s2 := s1[1:3:5]
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
}
