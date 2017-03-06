package assert

import (
	"fmt"
	"reflect"
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
}

type assertableValue struct {
	testContext *testContext
	value       interface{}
}

func (actual *assertableValue) Equals(expected interface{}) ValueAssertionResult {
	areEqual := reflect.DeepEqual(actual.value, expected)
	if !areEqual {
		if fmt.Sprint(actual.value) == fmt.Sprint(expected) {
			actual.testContext.decoratedErrorf(
				"Type mismatch.\nActual: %T=%v\nExpected: %T=%v\n", actual.value, actual.value, expected, expected)
		} else {
			actual.testContext.decoratedErrorf("Value mismatch.\nActual: %#v\nExpected: %#v\n", actual.value, expected)
		}
	}
	return &valueAssertionResult{bool: areEqual, actual: actual.value, expected: expected}
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

// anyOtherValue presents any other value but the expected.
// Since it's a private type, no value passed to DoesNotEqual can be
// equal to any instance of this type; hence, it can be used as an expected
// value for the ValueAssertionResult returned by DoesNotEqual.
type anyOtherValue struct{}
