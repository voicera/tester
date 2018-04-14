// Package ddt provides utilities to populate test cases for data-driven tests.
package ddt

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"runtime"
	"strings"
)

const (
	cannotGetTestFunctionNameErrorMessage = "ddt: cannot determine test function name required to load JSON file"
)

type dataDrivenTest struct {
	TestCases json.RawMessage `json:"testCases,omitempty"`
}

// LoadTestCasesFromDerivedJSONFile loads test cases from a JSON file whose path
// is derived from the caller's test function name and file. The file path is
// "<package under test>/_ddt/<basename of test file>.json"; for example,
// "hitchhiker/_ddt/question_test.json" with the following schema:
//
//  {"testCases": [{<properties of the test case to unmarshal>}, ...]}
//
// For example, the JSON content may look like the following:
//
//  {
//    "testCases": [
//      {
//        "id": "The Ultimate Question",
//        "input": {
//          "question": "What do you get when you multiply six by nine?",
//          "timeoutInHours": 65700000000,
//          "config": {"base": 13}
//        },
//        "expected": {
//          "answer": "42",
//          "error": null
//        }
//      }
//    ]
//  }
//
// The details of the test case struct are left for the tester to specify.
func LoadTestCasesFromDerivedJSONFile(testCasesToLoad interface{}) error {
	testFunctionName, err := getTestFunctionName()
	if err != nil {
		return err
	}
	testDataFilePath := "_ddt/" + testFunctionName + ".json"

	fileContent, err := ioutil.ReadFile(testDataFilePath)
	if err != nil {
		return err
	}

	test := &dataDrivenTest{}
	err = json.Unmarshal(fileContent, &test)
	if err != nil {
		return err
	}

	if len(test.TestCases) == 0 {
		return errors.New("ddt: cannot load test cases from " + testFunctionName + ".json")
	}
	return json.Unmarshal(test.TestCases, &testCasesToLoad)
}

func getTestFunctionName() (string, error) {
	for stackFramesToSkip := 2; ; stackFramesToSkip++ {
		callerProgramCounter, _, _, ok := runtime.Caller(stackFramesToSkip)
		if !ok {
			return "", errors.New(cannotGetTestFunctionNameErrorMessage)
		}
		functionFullPath := runtime.FuncForPC(callerProgramCounter).Name()
		functionName := functionFullPath[strings.LastIndex(functionFullPath, ".")+1:]
		if strings.HasPrefix(functionName, "Test") ||
			strings.HasPrefix(functionName, "Benchmark") ||
			strings.HasPrefix(functionName, "Example") {
			return functionName, nil
		}
	}
}
