package comp_test

import "testing"
import "github.com/morriship/TravisCITest"

func TestZeroGen(t *testing.T) {
	if comp.ZeroGen() != 0 {
		t.Error("Result should be zero")
	}
}

func TestOneGen(t *testing.T) {
	if comp.OneGen() != 1 {
		t.Error("Result should be one")
	}
}
