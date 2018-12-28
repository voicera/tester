package assert

import (
	"time"
)

// AssertableTime represents an under-test time that's expected to meet
// certain criteria.
type AssertableTime interface {
	// Equals asserts that the specified actual time equals the expected one.
	// Returns a ValueAssertionResult that provides post-assert actions.
	Equals(expected *time.Time) ValueAssertionResult

	// IsNil asserts that the specified actual time is nil.
	// Returns a ValueAssertionResult that provides post-assert actions.
	IsNil() ValueAssertionResult

	// IsNotNil asserts that the specified actual time is not nil.
	// Returns a ValueAssertionResult that provides post-assert actions.
	IsNotNil() ValueAssertionResult
}

type assertableTime struct {
	testContext *testContext
	value       *time.Time
}

func (actual *assertableTime) Equals(expected *time.Time) ValueAssertionResult {
	if expected == nil {
		return actual.IsNil()
	}
	if actual.value == nil {
		actual.testContext.decoratedErrorf("Time mismatch.\nActual was <nil>.\nExpected: %v\n", expected)
		return &valueAssertionResult{bool: false, actual: actual.value, expected: expected}
	}
	areEqual := actual.value.Equal(*expected)
	if !areEqual {
		actual.testContext.decoratedErrorf("Time mismatch.\nActual: %v\nExpected: %v\n", actual.value, expected)
	}
	return &valueAssertionResult{bool: areEqual, actual: actual.value, expected: expected}
}

func (actual *assertableTime) IsNil() ValueAssertionResult {
	if actual.value != nil {
		actual.testContext.decoratedErrorf("Actual time was not <nil>.\nActual: %v\n", actual.value)
	}
	return &valueAssertionResult{bool: actual.value == nil, actual: actual.value, expected: nil}
}

func (actual *assertableTime) IsNotNil() ValueAssertionResult {
	if actual.value == nil {
		actual.testContext.decoratedErrorf("Actual time was <nil>.\n")
	}
	return &valueAssertionResult{bool: actual.value != nil, actual: actual.value, expected: &anyOtherValue{}}
}
