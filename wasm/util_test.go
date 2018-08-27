package wasm

import "testing"

func TestAdd(t *testing.T) {
	if Madd(2, 3) != 5 {
		t.Fail()
	}

}

func TestString(t *testing.T) {
	if Mtext(3) != "This is a string: 3\n" {
		t.Fail()
	}
	print(Mtext(3))

}
