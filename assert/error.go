package assert

import "reflect"

// AssertableError represents an under-test error that's expected to meet
// certain criteria.
type AssertableError interface {
	// Equals asserts that the specified actual error equals the expected one.
	// Returns a ValueAssertionResult that provides post-assert actions.
	Equals(expected error) ValueAssertionResult

	// FormatsAs asserts that the specified actual error formats as expected.
	// The specified text is wrapped in an ErrorString object for comparison.
	// Returns a ValueAssertionResult that provides post-assert actions.
	FormatsAs(text string) ValueAssertionResult

	// IsNil asserts that the specified actual error is nil.
	// Returns a ValueAssertionResult that provides post-assert actions.
	IsNil() ValueAssertionResult

	// IsNotNil asserts that the specified actual error is not nil.
	// Returns a ValueAssertionResult that provides post-assert actions.
	IsNotNil() ValueAssertionResult
}

type assertableError struct {
	testContext *testContext
	value       error
}

// ErrorString is a trivial implementation of error. It's useful for asserting
// error messages. For example:
//
//     assert.For(t).ThatActualError(err).Equals(assert.ErrorString("foo"))
//
// We compare errors by comapring the output of Error(), which is used to format
// the error when printed. Here's a convenient way to rewrite the above:
//
//     assert.For(t).ThatActualError(err).FormatsAs("foo")
//
// The specified text is wrapped in an ErrorString object for comparison.
type ErrorString string

func (err ErrorString) Error() string {
	return string(err)
}

func (actual *assertableError) Equals(expected error) ValueAssertionResult {
	// Allow reflect to check for nil expected error as that object could have been loaded from JSON file (for DDT)
	if expected == nil || (reflect.ValueOf(expected).Kind() == reflect.Ptr && reflect.ValueOf(expected).IsNil()) {
		return actual.IsNil()
	}
	if actual.value == nil {
		actual.testContext.decoratedErrorf("Error mismatch.\nActual was <nil>.\nExpected: %v\n", expected)
		return &valueAssertionResult{bool: false, actual: actual.value, expected: expected}
	}
	// We're comparing interfaces — we only care about what Error() returns for both objects
	areEqual := actual.value.Error() == expected.Error()
	if !areEqual {
		actual.testContext.decoratedErrorf("Error mismatch.\nActual: %s\nExpected: %s\n", actual.value, expected)
	}
	return &valueAssertionResult{bool: areEqual, actual: actual.value, expected: expected}
}

func (actual *assertableError) FormatsAs(text string) ValueAssertionResult {
	return actual.Equals(ErrorString(text))
}

func (actual *assertableError) IsNil() ValueAssertionResult {
	if actual.value != nil { // no reflection here as we want to verify that the interface itself is not nil
		actual.testContext.decoratedErrorf("Actual error was not <nil>.\nActual: %v\n", actual.value)
	}
	return &valueAssertionResult{bool: actual.value == nil, actual: actual.value, expected: nil}
}

func (actual *assertableError) IsNotNil() ValueAssertionResult {
	if actual.value == nil { // no reflection here as we want to verify that the interface itself is nil
		actual.testContext.decoratedErrorf("Actual error was <nil>.\n")
	}
	return &valueAssertionResult{bool: actual.value != nil, actual: actual.value, expected: &anyOtherValue{}}
}
