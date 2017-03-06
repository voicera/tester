package assert

import "fmt"

// AssertableCall represents an under-test function call that's expected
// to meet certain criteria.
type AssertableCall interface {
	// PanicsReporting asserts that calling the specified callable causes
	// a panic with the specified expected error.
	PanicsReporting(interface{})
}

type assertableCall struct {
	testContext *testContext
	call        func()
}

func (callable *assertableCall) PanicsReporting(expectedError interface{}) {
	file, line := callable.testContext.caller() // must be set here to capture the right stack frame

	defer func() {
		if err := recover(); err == nil {
			callable.testContext.errorf(
				file, line, "Function call did not panic as expected.\nExpected: %s\n", expectedError)
		} else if fmt.Sprint(err) != fmt.Sprint(expectedError) {
			callable.testContext.errorf(
				file, line, "Panic message mismatch.\nActual: %s\nExpected: %s\n", err, expectedError)
		}
	}()

	callable.call()
}
