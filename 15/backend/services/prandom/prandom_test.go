package prandom

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	r := NewRandomer()
	start := r.seed
	r.roll()
	v1 := r.seed
	r.roll()
	v2 := r.seed
	r.roll()
	v3 := r.seed
	fmt.Println(start, v1, v2, v3, r.RandN(26), r.RandN(33))
}
