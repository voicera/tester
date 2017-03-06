package assert

func ExampleAssertableCall_panicsReportingPass() {
	For(t).ThatCalling(func() { panic("error") }).PanicsReporting("error")
	// Output:
}

func ExampleAssertableCall_panicsReportingFailPanicDidNotOccur() {
	mockTestContextToAssert().ThatCalling(func() {}).PanicsReporting("expected")
	// Output:
	// file:3: Function call did not panic as expected.
	// Expected: expected
}

func ExampleAssertableCall_panicsReportingFailMessageMismatch() {
	mockTestContextToAssert().ThatCalling(func() { panic("actual") }).PanicsReporting("expected")
	// Output:
	// file:3: Panic message mismatch.
	// Actual: actual
	// Expected: expected
}
