# Tester: Test More, Type Less

[![Build Status](https://travis-ci.org/voicera/tester.svg?branch=master)](https://travis-ci.org/voicera/tester)
[![Go Report Card](https://goreportcard.com/badge/github.com/voicera/tester)](https://goreportcard.com/report/github.com/voicera/tester)
[![GoDoc](https://godoc.org/github.com/voicera/tester?status.svg)](https://godoc.org/github.com/voicera/tester)
[![Maintainability](https://api.codeclimate.com/v1/badges/5a77a77b9efb9bacb55d/maintainability)](https://codeclimate.com/github/voicera/tester/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/5a77a77b9efb9bacb55d/test_coverage)](https://codeclimate.com/github/voicera/tester/test_coverage)

Lightweight test utilities to use with Go's testing package.

## Features
* Assertions that make tests easier to read, write, and debug
* Streamlined data providers for data-driven testing (DDT)
* Test hooks' hygiene check

## Motivations
Testing is an integral part of Go; the language provides strong and opinionated
support for testing. However, developers who moved from languages like C#, Java,
and Python miss the convenience of test utilities like assertions and data
providers for data-driven tests. That's why we started Tester: lightweight test
utilities to use with Go's testing package.

Most tests follow the same pattern: set up, invoke the unit under test, assert,
then clean up (if need be); said pattern encourages test code reuse
and consistency. By using test utilities, you can spend more time thinking about
test strategies and less time typing boilerplate code.

## Quick Start
Get the latest version (`go get -u github.com/voicera/tester`) then test away:

```go
package hitchhiker

import (
    "testing"

    "github.com/voicera/tester/assert"
)

func TestDeepThought(t *testing.T) {
    computer := NewDeepThoughtComputer()
    answer, err := computer.AnswerTheUltimateQuestion()
    if assert.For(t).ThatActual(err).IsNil().Passed() {
        assert.For(t).ThatActual(answer).Equals(42)
    }
}
```

## Learn More
The following can also be found at <https://godoc.org/github.com/voicera/tester>

### Assertions
Package `assert` provides a more readable way to assert in test cases;
for example:

```go
assert.For(t).ThatCalling(fn).PanicsReporting("expected error")
```

This way, the assert statement reads well; it flows like a proper sentence.

Also, one can easily tell which value is the test case got (actual)
and which it wanted (expected); this is key to printing the values correctly
to make debugging a bit easier. In Go, the actual value is usually printed
first; for example:

```go
assert.For(t).ThatActual(foo).Equals(expected)
```

The above enforces said order in both reading the code and the assertion failure
message (if any).

For convenience (that also improves readability), there are methods for special
cases like:

```go
assert.For(t).ThatActual(foo).IsTrue()
assert.For(t).ThatActualString(bar).IsEmpty()
```

Which are equivalent to:

```go
assert.For(t).ThatActual(foo).Equals(true)
assert.For(t).ThatActual(len(bar)).Equals(0)
```

To identify a test case in a table-driven test, optional parameters can be
specified and will be included in failure messages:

```go
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
```

After an assertion is performed, a `ValueAssertionResult` is returned to allow
for post-assert actions to be performed; for example:

```go
assert.For(t).ThatActual(value).Equals(expected).ThenDiffOnFail()
```

That will pretty-print a detailed recursive diff of both objects on failure.

It also can be used as a condition to perform extra test steps:

```go
value := GetValue()
if assert.For(t).ThatActual(value).IsNotNil().Passed() {
    assert.For(t).ThatActual(value.GetFoo()).Equals(expected)
}
```

Or to perform a deeper analysis of the test values:

```go
if !assert.For(t).ThatActual(value).Equals(expected).Passed() {
    analyze(value, expected) // e.g., analyze may look at common bugs
}
```

Conveniently, the last example above can be rewritten as:

```go
assert.For(t).ThatActual(value).Equals(expected).ThenRunOnFail(analyze)
```

Another use is to print a custom failure message:

```go
assert.For(t).ThatActual(foo).Equals(bar).ThenRunOnFail(func(actual, expected interface{}) {
    fmt.Printf("JSON: %q != %q", actual.ToJSON(), expected.ToJSON())
})
```

The above pattern allows for reuse of post-failure analysis and cleanup.

The interfaces in this package are still a work-in-progress, and are subject
to change.

### Test Hooks
What exists merely for test code to see shall not be exported to the world.
You can tag test-hook fields like the following:

```go
type sleeper struct {
    sleep func(time.Duration) `test-hook:"verify-unexported"`
}
```

Using tags, instead of comments, enables you to search the codebase for test
hooks and validate, via reflection, that they're not exported.
A test case should be added to verify that test hooks are hidden:

```go
func TestHooksAreHidden(t *testing.T) {
    assert.For(t).ThatType(reflect.TypeOf(sleeper{})).HidesTestHooks()
}
```

### Data-Driven Testing (DDT)
When the number of test cases in a table-driven test gets out of hand and they
cannot fit neatly in structs anymore, the use of a data provider is in order.
Package `ddt` provides a way to load test cases from a JSON file whose path
is derived from the caller's test function name and file. The file path is
`<package under test>/_ddt/<basename of test file>.json`; for example,
`hitchhiker/_ddt/question_test.json` with the following schema:

```json
{
  "testFunctions": {
    "<name of test function>": [
      {
        <properties of the test case to unmarshal>
      }
    ]
  }
}
```

For example, the JSON content may look like the following:

```json
{
  "testFunctions": {
    "TestDeepThought": [
      {
        "id": "The Ultimate Question",
        "input": {
          "question": "What do you get when you multiply six by nine?",
          "timeoutInHours": 65700000000,
          "config": {"base": 13}
        },
        "expected": {
          "answer": "42",
          "error": null
        }
      }
    ]
  }
}
```

The details of the test case struct are left for the tester to specify.
