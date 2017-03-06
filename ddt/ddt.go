// Package ddt provides utilities to populate test cases for data-driven tests.
package ddt

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path"
	"runtime"
	"strings"
)

// LoadTestCasesFromDerivedJSONFile loads test cases from a JSON file whose path
// is derived from the caller's test function name and file. The file path is
// "<package under test>/_ddt/<basename of test file>.json"; for example,
// "hitchhiker/_ddt/question_test.json" with the following schema:
//
//  {
//    "testFunctions": {
//        "<name of test function>": [
//        {
//            <properties of the test case to unmarshal>
//        }
//      ]
//    }
//  }
//
// For example, the JSON content may look like the following:
//
//  {
//    "testFunctions": {
//      "TestDeepThought": [
//        {
//          "id": "The Ultimate Question",
//          "input": {
//            "question": "What do you get when you multiply six by nine?",
//            "timeoutInHours": 65700000000,
//            "config": {"base": 13}
//          },
//          "expected": {
//            "answer": "42",
//            "error": null
//          }
//        }
//      ]
//    }
//  }
//
// The details of the test case struct are left for the tester to specify.
func LoadTestCasesFromDerivedJSONFile(testCasesToLoad interface{}) error {
	callerProgramCounter, _, _, _ := runtime.Caller(1)
	callerSegments := strings.Split(path.Base(runtime.FuncForPC(callerProgramCounter).Name()), ".")
	testDataFilePath := "_ddt/" + callerSegments[0] + ".json"

	testData, err := ioutil.ReadFile(testDataFilePath)
	if err != nil {
		return err
	}

	// Be verbose with parsing jsonData to throw better errors when needed (rather than a generic unmarshalling error)
	jsonData := map[string]map[string]json.RawMessage{}
	err = json.Unmarshal(testData, &jsonData)
	if err != nil {
		return err
	}

	testFunctions, ok := jsonData["testFunctions"]
	if !ok {
		return errors.New("Cannot find the testFunctions property")
	}

	rawTestCases, ok := testFunctions[callerSegments[1]]
	if !ok {
		return errors.New("Cannot find test cases for the caller function: " + callerSegments[1])
	}

	return json.Unmarshal(rawTestCases, &testCasesToLoad)
}
