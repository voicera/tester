package assert

import (
	"fmt"
	"net/mail"
)

func ExampleValueAssertionResult_ThenDiffOnFail_assertionFailed() {
	address := &mail.Address{Name: "Richard Hendricks", Address: "richard@pp.io"}
	expected := &mail.Address{Name: "Erlich Bachman", Address: "erlich@pp.io"}
	mockTestContextToAssert().ThatActual(address).Equals(expected).ThenDiffOnFail()
	// Output:
	// 	file:3: Value mismatch.
	// Actual: &mail.Address{Name:"Richard Hendricks", Address:"richard@pp.io"}
	// Expected: &mail.Address{Name:"Erlich Bachman", Address:"erlich@pp.io"}
	// Diff:
	// Name: "Richard Hendricks" != "Erlich Bachman"
	// Address: "richard@pp.io" != "erlich@pp.io"
}

func ExampleValueAssertionResult_ThenDiffOnFail_assertionPassed() {
	if mockTestContextToAssert().ThatActual(42).Equals(42).ThenDiffOnFail().Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleValueAssertionResult_ThenPrettyPrintOnFail_assertionFailed() {
	address := &mail.Address{Name: "Richard Hendricks", Address: "richard@pp.io"}
	expected := &mail.Address{Name: "Erlich Bachman", Address: "erlich@pp.io"}
	mockTestContextToAssert().ThatActual(address).Equals(expected).ThenPrettyPrintOnFail()
	// Output:
	// file:3: Value mismatch.
	// Actual: &mail.Address{Name:"Richard Hendricks", Address:"richard@pp.io"}
	// Expected: &mail.Address{Name:"Erlich Bachman", Address:"erlich@pp.io"}
	// Pretty:
	// Actual: "Richard Hendricks" <richard@pp.io>
	// Expected: "Erlich Bachman" <erlich@pp.io>
}

func ExampleValueAssertionResult_ThenRunOnFail_assertionFailed() {
	mockTestContextToAssert().ThatActualString("foo").Equals("bar").ThenRunOnFail(func(actual, expected interface{}) {
		fmt.Printf("Custom Message: %q != %q", actual, expected)
	})
	// Output:
	// file:3: String mismatch.
	// Actual: "foo"
	// Expected: "bar"
	// Custom Message: "foo" != "bar"
}

func ExampleValueAssertionResult_ThenRunOnFail_assertionPassed() {
	panicker := func(actual, expected interface{}) { panic("This should have never run!") }
	if mockTestContextToAssert().ThatActual(42).Equals(42).ThenRunOnFail(panicker).Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}
