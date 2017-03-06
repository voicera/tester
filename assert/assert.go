package assert

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"testing"

	"github.com/kr/pretty"
)

// TestContext provides methods to assert what the test actually got.
type TestContext interface {
	// ThatCalling adapts the specified call to an assertable one that's
	// expected to meet certain criteria.
	ThatCalling(func()) AssertableCall

	// ThatActual adapts the specified value to an assertable one that's
	// expected to meet certain criteria.
	ThatActual(value interface{}) AssertableValue

	// ThatActualString adapts the specified string to an assertable one that's
	// expected to meet certain criteria.
	// TODO rename to ThatActualString
	ThatActualString(value string) AssertableString

	// ThatType adapts the specified type to an assertable one that's
	// expected to meet certain criteria.
	ThatType(t reflect.Type) AssertableType
}

// testContext decorates and extends testing.TB that's passed to test functions
// to manage test state and support formatted test logs.
type testContext struct {
	testing.TB
	parameters []interface{}
	caller     func() (string, int) `test-hook:"verify-unexported"`
	fail       func()               `test-hook:"verify-unexported"`
}

const (
	testFileNameSuffix     = "_test.go"
	noCallerInfoLineNumber = -1
)

var printLock sync.Locker = &sync.Mutex{} // ensures that output is serialized

// For adapts from testing.TB to TestContext in order to allow
// the latter to assert on behalf of the former.
// The optional parameter(s) can be used to identify a specific test case
// in a data-driven test.
func For(t testing.TB, parameters ...interface{}) TestContext {
	return &testContext{t, parameters, caller, t.Fail}
}

func (testContext *testContext) ThatCalling(call func()) AssertableCall {
	return &assertableCall{testContext: testContext, call: call}
}

func (testContext *testContext) ThatActual(value interface{}) AssertableValue {
	return &assertableValue{testContext: testContext, value: value}
}

func (testContext *testContext) ThatActualString(value string) AssertableString {
	return &assertableString{testContext: testContext, value: value}
}

func (testContext *testContext) ThatType(t reflect.Type) AssertableType {
	return &assertableType{testContext: testContext, Type: t}
}

// PrintDiff prints a pretty diff of the specified actual and expected values,
// in that order.
func PrintDiff(actual interface{}, expected interface{}) {
	printLock.Lock()
	defer printLock.Unlock()

	fmt.Println("Diff:")
	pretty.Fdiff(os.Stdout, actual, expected)
}

func (testContext *testContext) decoratedErrorf(format string, args ...interface{}) {
	file, line := testContext.caller()
	testContext.errorf(file, line, format, args...)
}

func (testContext *testContext) errorf(file string, line int, format string, args ...interface{}) {
	printLock.Lock()
	defer printLock.Unlock()

	if line != noCallerInfoLineNumber {
		fmt.Printf("%s:%d: ", file, line) // because t.Errorf prints out the wrong file and line info
	}

	if len(testContext.parameters) > 0 {
		fmt.Print(testContext.parameters, " ")
	}

	fmt.Printf(format, args...)
	testContext.fail()
}

func caller() (file string, line int) {
	skip := 1
	ok := true

	for ok {
		_, file, line, ok = runtime.Caller(skip)
		if strings.HasSuffix(file, "_test.go") {
			return
		}
		skip++
	}
	return "", noCallerInfoLineNumber
}
