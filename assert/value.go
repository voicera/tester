package assert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	jsonIndent = "  "
)

// AssertableValue represents an under-test value that's expected to meet
// certain criteria.
type AssertableValue interface {
	// Equals asserts that the specified actual value equals the expected one.
	// See https://golang.org/pkg/reflect/#DeepEqual for equality rules.
	// Returns a ValueAssertionResult that provides post-assert actions.
	Equals(expected interface{}) ValueAssertionResult

	// DoesNotEqual asserts that the specified actual value does not equal
	// the unexpected one.
	// See https://golang.org/pkg/reflect/#DeepEqual for equality rules.
	// Returns a ValueAssertionResult that provides post-assert actions.
	DoesNotEqual(unexpected interface{}) ValueAssertionResult

	// IsNil asserts that the specified actual value is nil.
	// Returns a ValueAssertionResult that provides post-assert actions.
	IsNil() ValueAssertionResult

	// IsNotNil asserts that the specified actual value is not nil.
	// Returns a ValueAssertionResult that provides post-assert actions.
	IsNotNil() ValueAssertionResult

	// IsFalse asserts that the specified actual value is false.
	// Returns a ValueAssertionResult that provides post-assert actions.
	IsFalse() ValueAssertionResult

	// IsTrue asserts that the specified actual value is true.
	// Returns a ValueAssertionResult that provides post-assert actions.
	IsTrue() ValueAssertionResult

	// MarshalsEquivalentJSON asserts that the specified actual value yields
	// a JSON encoding equivalent to that of the specified expected value.
	// See https://golang.org/pkg/encoding/json/#Marshal for encoding details.
	// Returns a ValueAssertionResult that provides post-assert actions.
	MarshalsEquivalentJSON(expected interface{}) ValueAssertionResult
}

type assertableValue struct {
	testContext *testContext
	value       interface{}
}

// anyOtherValue presents any other value but the expected.
// Since it's a private type, no value passed to DoesNotEqual can be
// equal to any instance of this type; hence, it can be used as an expected
// value for the ValueAssertionResult returned by DoesNotEqual.
type anyOtherValue struct{}

func (actual *assertableValue) Equals(expected interface{}) ValueAssertionResult {
	areEqual := reflect.DeepEqual(actual.value, expected)
	if !areEqual {
		actual.printValueMismatchError(expected)
	}
	return &valueAssertionResult{bool: areEqual, actual: actual.value, expected: expected}
}

func (actual *assertableValue) printValueMismatchError(expected interface{}) {
	if fmt.Sprint(actual.value) == fmt.Sprint(expected) {
		actual.testContext.decoratedErrorf(
			"Type mismatch.\nActual: %T=%v\nExpected: %T=%v\n", actual.value, actual.value, expected, expected)
	} else {
		actual.testContext.decoratedErrorf("Value mismatch.\nActual: %#v\nExpected: %#v\n", actual.value, expected)
	}
}

func (actual *assertableValue) DoesNotEqual(value interface{}) ValueAssertionResult {
	areEqual := reflect.DeepEqual(actual.value, value)
	if areEqual {
		actual.testContext.decoratedErrorf("Values are equal.\nActual: %#v\n", actual.value)
	}
	return &valueAssertionResult{bool: !areEqual, actual: actual.value, expected: &anyOtherValue{}}
}

func (actual *assertableValue) IsNil() ValueAssertionResult {
	return actual.Equals(nil)
}

func (actual *assertableValue) IsNotNil() ValueAssertionResult {
	return actual.DoesNotEqual(nil)
}

func (actual *assertableValue) IsFalse() ValueAssertionResult {
	return actual.Equals(false)
}

func (actual *assertableValue) IsTrue() ValueAssertionResult {
	return actual.Equals(true)
}

func (actual *assertableValue) MarshalsEquivalentJSON(expected interface{}) ValueAssertionResult {
	var expectedBytes []byte
	actualBytes, err := json.MarshalIndent(actual.value, "", jsonIndent)
	if err != nil {
		goto mismatch
	}
	expectedBytes, err = json.MarshalIndent(expected, "", jsonIndent)
	if err != nil {
		goto mismatch
	}
	if !bytes.Equal(actualBytes, expectedBytes) {
		goto mismatch
	}
	return &valueAssertionResult{bool: true, actual: actual.value, expected: expected}

mismatch:
	actual.testContext.decoratedErrorf("JSON mismatch.\nActual: %s\nExpected: %s\n", actualBytes, expectedBytes)
	return &valueAssertionResult{bool: false, actual: actual.value, expected: expected}
}
