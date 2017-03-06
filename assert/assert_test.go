package assert

import (
	"reflect"
	"runtime"
	"testing"
)

func TestHooksAreHidden(t *testing.T) {
	For(t).ThatType(reflect.TypeOf(testContext{})).HidesTestHooks()
}

func TestCaller(t *testing.T) {
	unpacked := append(pack(caller()), pack(runtime.Caller(0))...) // make both calls fit on the same line
	file, line := unpacked[0], unpacked[1]
	expectedFile, expectedLine := unpacked[3], unpacked[4]
	For(t).ThatActual(file).Equals(expectedFile)
	For(t).ThatActual(line).Equals(expectedLine)
}

func pack(elements ...interface{}) []interface{} {
	return elements
}

var t = &testing.T{}

// mockTestContextToAssert mocks a test context to use for assertions.
// The optional parameter(s) can be used to identify a specific test case
// in a data-driven test.
func mockTestContextToAssert(parameters ...interface{}) *testContext {
	mock := For(t).(*testContext)
	mock.parameters = parameters
	mock.caller = func() (string, int) { return "file", 3 }
	mock.fail = func() {}
	return mock
}
