package assert

import (
	"errors"
	"fmt"
)

func ExampleAssertableValue_Equals_pass() {
	if For(t).ThatActual([]int{42}).Equals([]int{42}).Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableValue_Equals_fail() {
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

func ExampleAssertableValue_Equals_withErrors() {
	if !mockTestContextToAssert().ThatActual(errors.New("foo")).Equals(ErrorString("foo")).Passed() {
		fmt.Println("Don't use ThatActual with errors — use ThatActualError instead!")
	}
	// Output:
	// file:3: Type mismatch.
	// Actual: *errors.errorString=foo
	// Expected: assert.ErrorString=foo
	// Don't use ThatActual with errors — use ThatActualError instead!
}

func ExampleAssertableValue_DoesNotEqual_pass() {
	if For(t).ThatActual(42).DoesNotEqual(13).Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableValue_DoesNotEqual_fail() {
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

func ExampleAssertableValue_IsNil_pass() {
	if For(t).ThatActual(nil).IsNil().Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableValue_IsNil_fail() {
	if !mockTestContextToAssert().ThatActual(42).IsNil().Passed() {
		fmt.Println("Assertion failed successfully!")
	}
	// Output:
	// file:3: Value mismatch.
	// Actual: 42
	// Expected: <nil>
	// Assertion failed successfully!
}

func ExampleAssertableValue_IsNotNil_pass() {
	if For(t).ThatActual(42).IsNotNil().Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableValue_IsNotNil_fail() {
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

func ExampleAssertableValue_IsFalse_pass() {
	if For(t).ThatActual(false).IsFalse().Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableValue_IsFalse_fail() {
	if !mockTestContextToAssert().ThatActual(true).IsFalse().Passed() {
		fmt.Println("Assertion failed successfully!")
	}
	// Output:
	// file:3: Value mismatch.
	// Actual: true
	// Expected: false
	// Assertion failed successfully!
}

func ExampleAssertableValue_IsTrue_pass() {
	if For(t).ThatActual(true).IsTrue().Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableValue_IsTrue_fail() {
	if !mockTestContextToAssert().ThatActual(false).IsTrue().Passed() {
		fmt.Println("Assertion failed successfully!")
	}
	// Output:
	// file:3: Value mismatch.
	// Actual: false
	// Expected: true
	// Assertion failed successfully!
}

func ExampleAssertableValue_MarshalsEquivalentJSON_pass() {
	foo := &struct {
		Key string `json:"key"`
	}{
		Key: "value",
	}
	equivalentToFoo := &struct {
		Key string `json:"key"`
	}{
		Key: "value",
	}

	if For(t).ThatActual(foo).MarshalsEquivalentJSON(equivalentToFoo).Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableValue_MarshalsEquivalentJSON_fail() {
	if !mockTestContextToAssert().ThatActual(nil).MarshalsEquivalentJSON("").Passed() {
		fmt.Println("Assertion failed successfully!")
	}
	// Output:
	// file:3: JSON mismatch.
	// Actual: null
	// Expected: ""
	// Assertion failed successfully!
}
