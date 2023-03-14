package qnap

import "testing"

func TestEncodePwd(t *testing.T) {
	if encoded := encodePwd(""); encoded != "" {
		t.Fatal(encoded)
	}
	if encoded := encodePwd("b"); encoded != "Yg==" {
		t.Fatal(encoded)
	}
	if encoded := encodePwd("中文"); encoded != "LYc=" {
		t.Fatal(encoded)
	}
}
