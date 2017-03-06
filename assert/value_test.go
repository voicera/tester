package assert

import "fmt"

func ExampleAssertableValue_equalsPass() {
	if For(t).ThatActual([]int{42}).Equals([]int{42}).Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableValue_equalsFail() {
	cases := []struct {
		id       string
		actual   interface{}
		expected interface{}
	}{
		{"different values", 42, 13},
		{"different types", 42.0, 42},
		{"different containers", [...]int{42}, []int{42}},
	}

	for _, c := range cases {
		if !mockTestContextToAssert(c.id).ThatActual(c.actual).Equals(c.expected).Passed() {
			fmt.Println("Assertion failed successfully!")
		}
	}
	// Output:
	// file:3: [different values] Value mismatch.
	// Actual: 42
	// Expected: 13
	// Assertion failed successfully!
	// file:3: [different types] Type mismatch.
	// Actual: float64=42
	// Expected: int=42
	// Assertion failed successfully!
	// file:3: [different containers] Type mismatch.
	// Actual: [1]int=[42]
	// Expected: []int=[42]
	// Assertion failed successfully!
}

func ExampleAssertableValue_doesNotEqualPass() {
	if For(t).ThatActual(42).DoesNotEqual(13).Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableValue_doesNotEqualFail() {
	if !mockTestContextToAssert().ThatActual(42).DoesNotEqual(42).ThenDiffOnFail().Passed() {
		fmt.Println("Assertion failed successfully!")
	}
	// Output:
	// file:3: Values are equal.
	// Actual: 42
	// Diff:
	// int != *assert.anyOtherValue
	// Assertion failed successfully!
}

func ExampleAssertableValue_isNilPass() {
	if For(t).ThatActual(nil).IsNil().Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableValue_isNilFail() {
	if !mockTestContextToAssert().ThatActual(42).IsNil().Passed() {
		fmt.Println("Assertion failed successfully!")
	}
	// Output:
	// file:3: Value mismatch.
	// Actual: 42
	// Expected: <nil>
	// Assertion failed successfully!
}

func ExampleAssertableValue_isNotNilPass() {
	if For(t).ThatActual(42).IsNotNil().Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableValue_isNotNilFail() {
	if !mockTestContextToAssert().ThatActual(nil).IsNotNil().ThenDiffOnFail().Passed() {
		fmt.Println("Assertion failed successfully!")
	}
	// Output:
	// file:3: Values are equal.
	// Actual: <nil>
	// Diff:
	// nil != &assert.anyOtherValue{}
	// Assertion failed successfully!
}

func ExampleAssertableValue_isFalsePass() {
	if For(t).ThatActual(false).IsFalse().Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableValue_isFalseFail() {
	if !mockTestContextToAssert().ThatActual(true).IsFalse().Passed() {
		fmt.Println("Assertion failed successfully!")
	}
	// Output:
	// file:3: Value mismatch.
	// Actual: true
	// Expected: false
	// Assertion failed successfully!
}

func ExampleAssertableValue_isTruePass() {
	if For(t).ThatActual(true).IsTrue().Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableValue_isTrueFail() {
	if !mockTestContextToAssert().ThatActual(false).IsTrue().Passed() {
		fmt.Println("Assertion failed successfully!")
	}
	// Output:
	// file:3: Value mismatch.
	// Actual: false
	// Expected: true
	// Assertion failed successfully!
}
