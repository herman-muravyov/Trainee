package hello 

import (
	"testing"
	"test/variables"
)

func TestHello(t *testing.T) {
	exp := variables.Str
	call := Hello()
	if call != exp {
		t.Errorf("Func's call is: %q; we want: %s", call, exp)
	}
}
