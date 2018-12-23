package assert

import (
	"errors"
	"fmt"
)

func ExampleAssertableError_Equals_pass() {
	cases := []struct {
		id       string
		actual   error
		expected error
	}{
		{"both are nil", nil, nil},
		{"same type and message", errors.New("foo"), errors.New("foo")},
		{"different struct types, same message", ErrorString("foo"), errors.New("foo")},
	}

	for _, c := range cases {
		if For(t, c.id).ThatActualError(c.actual).Equals(c.expected).Passed() {
			fmt.Println("Passed: " + c.id)
		}
	}
	// Output:
	// Passed: both are nil
	// Passed: same type and message
	// Passed: different struct types, same message
}

func ExampleAssertableError_Equals_fail() {
	cases := []struct {
		id       string
		actual   error
		expected error
	}{
		{"expected is nil while actual isn't", ErrorString("foo"), nil},
		{"actual is nil while expected isn't", nil, ErrorString("foo")},
		{"different messages", ErrorString("foo"), ErrorString("bar")},
	}

	for _, c := range cases {
		if !mockTestContextToAssert(c.id).ThatActualError(c.actual).Equals(c.expected).Passed() {
			fmt.Println("Assertion failed successfully!")
		}
	}
	// Output:
	// file:3: [expected is nil while actual isn't] Actual error was not <nil>.
	// Actual: foo
	// Assertion failed successfully!
	// file:3: [actual is nil while expected isn't] Error mismatch.
	// Actual was <nil>.
	// Expected: foo
	// Assertion failed successfully!
	// file:3: [different messages] Error mismatch.
	// Actual: foo
	// Expected: bar
	// Assertion failed successfully!
}

func ExampleAssertableError_IsNil_pass() {
	if For(t).ThatActualError(nil).IsNil().Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableError_IsNil_fail() {
	if !mockTestContextToAssert().ThatActualError(errors.New("foo")).IsNil().Passed() {
		fmt.Println("Assertion failed successfully!")
	}
	// Output:
	// file:3: Actual error was not <nil>.
	// Actual: foo
	// Assertion failed successfully!
}

func ExampleAssertableError_IsNotNil_pass() {
	if For(t).ThatActualError(errors.New("foo")).IsNotNil().Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableError_IsNotNil_fail() {
	if !mockTestContextToAssert().ThatActualError(nil).IsNotNil().ThenDiffOnFail().Passed() {
		fmt.Println("Assertion failed successfully!")
	}
	// Output:
	// file:3: Actual error was <nil>.
	// Diff:
	// nil != &assert.anyOtherValue{}
	// Assertion failed successfully!
}
