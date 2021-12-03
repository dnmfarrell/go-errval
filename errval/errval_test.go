package errval

import (
	"errors"
	"fmt"
	"testing"
)

func TestCatchFailure(t *testing.T) {
	ev := Err[string](errors.New("failure"))
	handlerCall := false
	handler := func(_ error) { handlerCall = true }
	success, val := ev.Catch(handler)
	if success {
		t.Error("Catch returns true for failure")
	}
	if val != "" {
		t.Errorf("Catch returns non-zero value: %s for failure", val)
	}
	if !handlerCall {
		t.Error("Catch does not call error handler for failure")
	}
}

func TestCatchSuccess(t *testing.T) {
	ev := Val[string]("foo")
	handlerCall := false
	handler := func(_ error) { handlerCall = true }
	success, val := ev.Catch(handler)
	if !success {
		t.Error("Catch returns false for success")
	}
	if val != "foo" {
		t.Errorf("Catch returns unexpected value: %s for success", val)
	}
	if handlerCall {
		t.Error("Catch calls error handler for success")
	}
}

func ExampleCatch() {
	logger := func(e error) { fmt.Println(e) }
	foo := func() *ErrVal[string] { return Val[string]("bar") }
	if ok, val := foo().Catch(logger); ok {
		fmt.Println(val)
		// Output: bar
	}
}
