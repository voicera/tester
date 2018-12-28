package assert

import (
	"fmt"
	"time"
)

var (
	zero  = &time.Time{}
	epoch = time.Unix(0, 0).UTC()
)

func ExampleAssertableTime_Equals_pass() {
	var nilTime *time.Time
	cases := []struct {
		id       string
		actual   *time.Time
		expected *time.Time
	}{
		{"nils", nil, nilTime},
		{"same object (wall and monotonic clocks)", &epoch, &epoch},
		{"same wall-clock value", zero, &time.Time{}},
	}

	for _, c := range cases {
		if For(t, c.id).ThatActual(c.actual).Equals(c.expected).Passed() {
			fmt.Println("Passed: " + c.id)
		}
	}
	// Output:
	// Passed: nils
	// Passed: same object (wall and monotonic clocks)
	// Passed: same wall-clock value
}

func ExampleAssertableTime_Equals_fail() {
	cases := []struct {
		id       string
		actual   *time.Time
		expected *time.Time
	}{
		{"expected is nil while actual isn't", &epoch, nil},
		{"actual is nil while expected isn't", nil, &epoch},
		{"different values", &epoch, zero},
	}

	for _, c := range cases {
		if !mockTestContextToAssert(c.id).ThatActualTime(c.actual).Equals(c.expected).Passed() {
			fmt.Println("Assertion failed successfully!")
		}
	}
	// Output:
	// file:3: [expected is nil while actual isn't] Actual time was not <nil>.
	// Actual: 1970-01-01 00:00:00 +0000 UTC
	// Assertion failed successfully!
	// file:3: [actual is nil while expected isn't] Time mismatch.
	// Actual was <nil>.
	// Expected: 1970-01-01 00:00:00 +0000 UTC
	// Assertion failed successfully!
	// file:3: [different values] Time mismatch.
	// Actual: 1970-01-01 00:00:00 +0000 UTC
	// Expected: 0001-01-01 00:00:00 +0000 UTC
	// Assertion failed successfully!
}

func ExampleAssertableTime_IsNil_pass() {
	if For(t).ThatActualTime(nil).IsNil().Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableTime_IsNil_fail() {
	if !mockTestContextToAssert().ThatActualTime(zero).IsNil().Passed() {
		fmt.Println("Assertion failed successfully!")
	}
	// Output:
	// file:3: Actual time was not <nil>.
	// Actual: 0001-01-01 00:00:00 +0000 UTC
	// Assertion failed successfully!
}

func ExampleAssertableTime_IsNotNil_pass() {
	if For(t).ThatActualTime(zero).IsNotNil().Passed() {
		fmt.Println("Passed!")
	}
	// Output: Passed!
}

func ExampleAssertableTime_IsNotNil_fail() {
	if !mockTestContextToAssert().ThatActualTime(nil).IsNotNil().ThenDiffOnFail().Passed() {
		fmt.Println("Assertion failed successfully!")
	}
	// Output:
	// file:3: Actual time was <nil>.
	// Diff:
	// *time.Time != *assert.anyOtherValue
	// Assertion failed successfully!
}
