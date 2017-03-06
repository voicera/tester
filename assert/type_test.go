package assert

import "reflect"

func ExampleAssertableType_hidesTestHooksPass() {
	type EmptyType struct{}
	type ExportedType struct{}
	type ExportedPointerType struct{}
	type unexportedType struct{}
	type unexportedPointerType struct{}
	type PassingFieldsType struct {
		ExportedTestHook func() `test-hook:""`
		ExportedType
		unexportedTestHook     func() `test-hook:"verify-unexported"`
		unexportedType         `test-hook:"verify-unexported"`
		*unexportedPointerType `test-hook:"verify-unexported"`
	}

	cases := []struct {
		id     string
		object interface{}
	}{
		{"no fields", EmptyType{}},
		{"passing fields", PassingFieldsType{}},
		{"unexported test hooks", struct { // the enclosing type is unexported
			ExportedTestHook     func() `test-hook:"verify-unexported"`
			ExportedType         `test-hook:"verify-unexported"`
			*ExportedPointerType `test-hook:"verify-unexported"`
		}{}},
	}

	for _, c := range cases {
		For(t, c.id, c.object).ThatType(reflect.TypeOf(c.object)).HidesTestHooks()
	}
	// Output:
}

func ExampleAssertableType_hidesTestHooksExportedTestHooks() {
	type ExportedPointerType struct{}
	type ExportedType struct{}
	type ExposedFieldsType struct {
		ExportedTestHook     func() `test-hook:"verify-unexported"`
		ExportedType         `test-hook:"verify-unexported"`
		*ExportedPointerType `test-hook:"verify-unexported"`
	}

	mockTestContextToAssert().ThatType(reflect.TypeOf(ExposedFieldsType{})).HidesTestHooks()
	// Output:
	// file:3: Type ExposedFieldsType exports test-hook fields: [{Name:ExportedTestHook PkgPath: Type:func() Tag:test-hook:"verify-unexported" Offset:0 Index:[0] Anonymous:false} {Name:ExportedType PkgPath: Type:assert.ExportedType Tag:test-hook:"verify-unexported" Offset:8 Index:[1] Anonymous:true} {Name:ExportedPointerType PkgPath: Type:*assert.ExportedPointerType Tag:test-hook:"verify-unexported" Offset:8 Index:[2] Anonymous:true}]
}
