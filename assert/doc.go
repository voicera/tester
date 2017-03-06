/*
Package assert provides a more readable way to assert in test cases;
for example:

    assert.For(t).ThatCalling(fn).PanicsReporting("expected error")

This way, the assert statement reads well; it flows like a proper sentence.

Also, one can easily tell which value is the test case got (actual)
and which it wanted (expected); this is key to printing the values correctly
to make debugging a bit easier. In Go, the actual value is usually printed
first; for example:

    assert.For(t).ThatActual(foo).Equals(expected)

The above enforces said order in both reading the code and the assertion failure
message (if any).

For convenience (that also improves readability), there are methods for special
cases like:

    assert.For(t).ThatActual(foo).IsTrue()
    assert.For(t).ThatActualString(bar).IsEmpty()

Which are equivalent to:

    assert.For(t).ThatActual(foo).Equals(true)
    assert.For(t).ThatActual(len(bar)).Equals(0)

To identify a test case in a table-driven test, optional ID parameters can be
specified and will be included in failure messages:

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
        assert.For(t, c.id).ThatActual(c.actual).Equals(c.expected)
    }

After an assertion is performed, a ValueAssertionResult is returned to allow
for post-assert actions to be performed; for example:

    assert.For(t).ThatActual(value).Equals(expected).ThenDiffOnFail()

Which will pretty-print a detailed diff of both objects recursively on failure.

It also can be used as a condition to perform extra test steps:

    value := GetValue()
    if assert.For(t).ThatActual(value).IsNotNil().Passed() {
        assert.For(t).ThatActual(value.GetFoo()).Equals(expected)
    }

Or to perform a deeper analysis of the test values:

    if !assert.For(t).ThatActual(value).Equals(expected).Passed() {
        analyze(value, expected) // e.g., analyze may look at common bugs
    }

Conveniently, the last example above can be rewritten as:

    assert.For(t).ThatActual(value).Equals(expected).ThenRunOnFail(analyze)

Another use is to print a custom failure message:

    assert.For(t).ThatActual(foo).Equals(bar).ThenRunOnFail(func(actual, expected interface{}) {
        fmt.Printf("JSON: %q != %q", actual.ToJSON(), expected.ToJSON())
    })

The above pattern allows for reuse of post-failure analysis and cleanup.

The interfaces in this package are still a work-in-progress, and are subject
to change.
*/
package assert
