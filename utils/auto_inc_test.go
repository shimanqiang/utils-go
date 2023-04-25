package utils

import (
	"fmt"
	"testing"
)

func TestAutoInc(t *testing.T) {
	inc := NewAutoInc(3, 100)

	for i := 0; i < 100; i++ {
		fmt.Println(inc.NextId())
	}
}
