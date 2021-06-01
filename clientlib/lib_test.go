package clientlib

import "testing"

func TestHello(t *testing.T) {
	hello := Hello()
	want := "clientlib hello"

	if hello != want {
		t.Errorf("got %v, want %s", hello, want)
	}
}
