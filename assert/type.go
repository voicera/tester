package assert

import (
	"go/ast"
	"reflect"
)

// AssertableType represents an under-test type that's expected to meet
// certain criteria.
type AssertableType interface {
	// HidesTestHooks asserts that the under-test type hides test hooks.
	// It asserts that fields tagged with TestHookTagKey; for example:
	//     type sleeper struct {
	//         sleep func(time.Duration) `test-hook:"verify-unexported"`
	//     }
	// are unexported. If the field is anonymous, it asserts that its type is
	// unexported. An empty test-hook tag value is equivalent to no test-hook
	// tag; in which case, HidesTestHooks does not check the field.
	HidesTestHooks()
}

type assertableType struct {
	testContext *testContext
	reflect.Type
}

const (
	// TestHookTagKey denotes the tag key to use to tag a field as a test hook.
	TestHookTagKey = "test-hook"
)

func (actual *assertableType) HidesTestHooks() {

	if len(actual.Name()) == 0 || !ast.IsExported(actual.Name()) { // anonymous or unexported type
		return
	}

	failedFields := []reflect.StructField{}
	for i := 0; i < actual.NumField(); i++ {
		field := actual.Field(i)
		if tag := field.Tag.Get(TestHookTagKey); tag != "" && ast.IsExported(field.Name) {
			failedFields = append(failedFields, field)
		}
	}

	if len(failedFields) > 0 {
		actual.testContext.decoratedErrorf("Type %s exports test-hook fields: %+v\n", actual.Name(), failedFields)
	}
}
