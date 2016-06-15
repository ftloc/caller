package caller_test

import (
	"github.com/ftloc/caller"
	"testing"
)

func TestCaller(t *testing.T) {
	called := false
	caller.CallWith(func(s string) {
		if s != "test" {
			t.Fail()
		}
		called = true
	}, "test")
	if !called {
		t.Fail()
	}

	thrown := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				thrown = true
			}
		}()
		caller.CallWith(1)
	}()

	if !thrown {
		t.Fail()
	}
}
