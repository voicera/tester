# Contributing, Developer Guide, and Coding Conventions
Thank you for taking the time to improve tester. Below are a few guidelines
to check before contributing.

## Tenets
* Tester is opinionated
* Tester is lightweight
* Tester is intuitive
* Tester is consistent

## Dependency Management
Copy dependencies to the `vendor` folder.
See <https://golang.org/cmd/go/#hdr-Vendor_Directories> for more details.

## Makefile
It's there to take care of some best practices for you; please use it.

## Code Review Guidelines
See <https://github.com/golang/go/wiki/CodeReviewComments>.

### Addenda
The rule of thumb is: if `go fmt` allows it, so be it. The guidelines below
intend to make the coding style more consistent.

#### Vertical Ordering
From _Clean Code, Chapter 5: Formatting_
> We would like a source file to be like a newspaper article. The name should be
simple but explanatory. The name, by itself, should be sufficient to tell us
whether we are in the right module or not. The topmost parts of the source file
should provide the high-level concepts and algorithms. Detail should increase as
we move downward, until at the end we find the lowest level functions
and details in the source file... In general we want function call dependencies
to point in the downward direction. That is, a function that is called should be
below a function that does the calling. This creates a nice flow down the source
code module from high level to low level.

Also, see <https://talks.golang.org/2013/bestpractices.slide#14>,
which recommends the following order:
1. Header: license information, build tags, package documentation.
1. Imports: related groups separated by blank lines.
1. Body: the rest of the code starting with the most significant types,
and ending with helper functions and helper types.

#### Line Length
Keep line length under 120 characters for code and 80 for documentation.
Break lines in a way that improves readability but still keeps the code compact;
for example,
```go
callable.testContext.errorf(
    file, line, "Function call did not panic as expected.\nExpected: %s\n", expectedError)
```
is readable and more compact than
```go
callable.testContext.errorf(
    file,
    line,
    "Function call did not panic as expected.\nExpected: %s\n",
    expectedError)
```
Either is definitely more readable than
```go
callable.testContext.errorf(file, line,
    "Function call did not panic as expected.\nExpected: %s\n", expectedError)
```

#### Variable Names
Do not prefer `c` to `lineCount`; see _Clean Code, Chapter 2: Meaningful Names_.
> There can be no worse reason for using the [variable] name c than because a
and b were already taken.

That said, find the shortest name that's self-explanatory:

> Shorter names are generally better than longer ones, so long as they are
clear. Add no more context to a name than is necessary.

Loop variables can be a single letter (given that the loop block is short).
