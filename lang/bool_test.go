package lang

import "testing"

func TestBool_X(t *testing.T) {
	newBool := NewBool(true)
	if newBool.Load() != true {
		t.Fail()
	}
}
