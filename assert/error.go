package assert

// AssertableError represents an under-test error that's expected to meet
// certain criteria.
type AssertableError interface {
	// Equals asserts that the specified actual error equals the expected one.
	// Returns a ValueAssertionResult that provides post-assert actions.
	Equals(expected error) ValueAssertionResult

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
// The assert library compares errors by comapring the output of Error().
type ErrorString string

func (err ErrorString) Error() string {
	return string(err)
}

func (actual *assertableError) Equals(expected error) ValueAssertionResult {
	if expected == nil {
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

func (actual *assertableError) IsNil() ValueAssertionResult {
	if actual.value != nil {
		actual.testContext.decoratedErrorf("Actual error was not <nil>.\nActual: %v\n", actual.value)
	}
	return &valueAssertionResult{bool: actual.value == nil, actual: actual.value, expected: nil}
}

func (actual *assertableError) IsNotNil() ValueAssertionResult {
	if actual.value == nil {
		actual.testContext.decoratedErrorf("Actual error was <nil>.\n")
	}
	return &valueAssertionResult{bool: actual.value != nil, actual: actual.value, expected: &anyOtherValue{}}
}
