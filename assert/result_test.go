package assert

import (
	"fmt"
	"net/mail"
)

func ExampleValueAssertionResult_ThenDiffOnFail_whenAssertionFails() {
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

func ExampleValueAssertionResult_ThenDiffOnFail_whenAssertionPasses() {
	if mockTestContextToAssert().ThatActual(42).Equals(42).ThenDiffOnFail().Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleValueAssertionResult_ThenRunOnFail_whenAssertionFails() {
	mockTestContextToAssert().ThatActualString("foo").Equals("bar").ThenRunOnFail(func(actual, expected interface{}) {
		fmt.Printf("Custom Message: %q != %q", actual, expected)
	})
	// Output:
	// file:3: String mismatch.
	// Actual: "foo"
	// Expected: "bar"
	// Custom Message: "foo" != "bar"
}

func ExampleValueAssertionResult_ThenRunOnFail_whenAssertionPasses() {
	panicker := func(actual, expected interface{}) { panic("This should have never run!") }
	if mockTestContextToAssert().ThatActual(42).Equals(42).ThenRunOnFail(panicker).Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}
