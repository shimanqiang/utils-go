package utils

import (
	"fmt"
	"testing"
)

func TestAutoInc(t *testing.T) {
	inc, err := NewAutoInc(3, 100)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	for i := 0; i < 100; i++ {
		fmt.Println(inc.NextId())
	}
}
