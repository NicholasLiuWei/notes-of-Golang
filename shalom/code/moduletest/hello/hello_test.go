package hello

import "testing"

func TestHello(t *testing.T) {
	want:="hello,world!"
	if got:=Hello();got!=want{
		t.Errorf("Hello()=%s want: %s",got,want)
	}
}
