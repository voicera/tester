package assert

import "fmt"

func ExampleAssertableString_equalsPass() {
	if For(t).ThatActualString("foo").Equals("foo").Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableString_equalsFail() {
	cases := []struct {
		id       string
		actual   string
		expected string
	}{
		{"different letters", "foo", "bar"},
		{"different casing", "foo", "Foo"},
		{"empty (for logging)", "", "a"},
	}

	for _, c := range cases {
		if !mockTestContextToAssert(c.id).ThatActualString(c.actual).Equals(c.expected).Passed() {
			fmt.Println("Assertion failed successfully!")
		}
	}
	// Output:
	// file:3: [different letters] String mismatch.
	// Actual: "foo"
	// Expected: "bar"
	// Assertion failed successfully!
	// file:3: [different casing] String mismatch.
	// Actual: "foo"
	// Expected: "Foo"
	// Assertion failed successfully!
	// file:3: [empty (for logging)] String mismatch.
	// Actual: ""
	// Expected: "a"
	// Assertion failed successfully!
}

func ExampleAssertableValue_isEmptyPass() {
	if For(t).ThatActualString("").IsEmpty().Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableValue_isEmptyFail() {
	if !mockTestContextToAssert().ThatActualString("foo").IsEmpty().ThenDiffOnFail().Passed() {
		fmt.Println("Assertion failed successfully!")
	}
	// Output:
	// file:3: String is not empty.
	// Actual: "foo"
	// Diff:
	// "foo" != ""
	// Assertion failed successfully!
}
