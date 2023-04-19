package lang

import (
	"fmt"
	"testing"
)

func TestSet_X(t *testing.T) {
	a := NewSet()
	a.Add(1)
	a.Add(3)

	for i := range a.Iter() {
		fmt.Println(i)
	}

	a.Remove(3)
	fmt.Println("------", a.Size())

	for i := range a.Iter() {
		fmt.Println(i)
	}
}
