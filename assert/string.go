package assert

// AssertableString represents an under-test string that's expected to meet
// certain criteria.
type AssertableString interface {
	// Equals asserts that the specified actual string equals the expected one.
	// String values are compared byte-wise (implies case-sensetivity).
	// Returns a ValueAssertionResult that provides post-assert actions.
	Equals(expected string) ValueAssertionResult

	// IsEmpty asserts that the specified actual string is empty.
	// Returns a ValueAssertionResult that provides post-assert actions.
	IsEmpty() ValueAssertionResult

	// IsNotEmpty asserts that the specified actual string is not empty.
	// Returns a ValueAssertionResult that provides post-assert actions.
	IsNotEmpty() ValueAssertionResult
}

type assertableString struct {
	testContext *testContext
	value       string
}

func (actual *assertableString) Equals(expected string) ValueAssertionResult {
	areEqual := actual.value == expected
	if !areEqual {
		actual.testContext.decoratedErrorf("String mismatch.\nActual: %q\nExpected: %q\n", actual.value, expected)
	}
	return &valueAssertionResult{bool: areEqual, actual: actual.value, expected: expected}
}

func (actual *assertableString) IsEmpty() ValueAssertionResult {
	isEmpty := actual.value == ""
	if !isEmpty {
		actual.testContext.decoratedErrorf("String is not empty.\nActual: %q\n", actual.value)
	}
	return &valueAssertionResult{bool: isEmpty, actual: actual.value, expected: ""}
}

func (actual *assertableString) IsNotEmpty() ValueAssertionResult {
	isEmpty := actual.value == ""
	if isEmpty {
		actual.testContext.decoratedErrorf("String is empty.\n")
	}
	return &valueAssertionResult{bool: !isEmpty, actual: actual.value, expected: "<any non-empty string>"}
}
