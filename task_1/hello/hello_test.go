package hello 

import "testing"

func TestHello(t *testing.T) {
	exp := "Hello."
	call := Hello()
	if call != exp {
		t.Errorf("Func's call is: %q; we want: %s", call, exp)
	}
}
