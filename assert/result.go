package assert

// ValueAssertionResult represents operations that may be performed on
// the result of value assertion. For example:
//
//     assert.For(t).ThatActual(value).Equals(expected).ThenDiffOnFail()
//
// It also can be used as a condition to perform extra test steps:
//
//     value := GetValue()
//     if assert.For(t).ThatActual(value).IsNotNil().Passed() {
//         assert.For(t).ThatActual(value.GetFoo()).Equals(expected)
//     }
//
// Or to perform a deeper analysis of the test values:
//
//     if !assert.For(t).ThatActual(value).Equals(expected).Passed() {
//         analyze(value, expected) // e.g., analyze may look at common bugs
//     }
//
// Conveniently, the last example above can be rewritten as:
//
//     assert.For(t).ThatActual(value).Equals(expected).ThenRunOnFail(analyze)
//
// The above pattern allows for reuse of post-failure analysis and cleanup.
type ValueAssertionResult interface {
	// Passed returns true if the assertion passed.
	Passed() bool

	// ThenDiffOnFail performed a diff of asserted values on assertion failure;
	// it prints a pretty diff of the actual and expected values used in
	// the failed assertion, in that order.
	// Returns the current ValueAssertionResult to allow for call-chaining.
	ThenDiffOnFail() ValueAssertionResult

	// ThenRunOnFail performed the specified action on assertion failure;
	// in which case, it passes the actual and expected values used in
	// the failed assertion as parameters to the specified function.
	// Returns the current ValueAssertionResult to allow for call-chaining.
	ThenRunOnFail(action func(actual, expected interface{})) ValueAssertionResult
}

type valueAssertionResult struct {
	bool
	actual   interface{}
	expected interface{}
}

func (result *valueAssertionResult) Passed() bool {
	return result.bool
}

func (result *valueAssertionResult) ThenDiffOnFail() ValueAssertionResult {
	return result.ThenRunOnFail(PrintDiff)
}

func (result *valueAssertionResult) ThenRunOnFail(action func(actual, expected interface{})) ValueAssertionResult {
	if !result.Passed() {
		action(result.actual, result.expected)
	}
	return result
}
