package ddt_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/voicera/tester/assert"
	"github.com/voicera/tester/ddt"
)

const vanillaContent = `{
  "testFunctions": {
    "ExampleLoadTestCasesFromDerivedJSONFile": [
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
      },
      {
        "id": "Ask Again",
        "input": {
          "question": "?",
          "timeoutInHours": 65700000000
        },
        "expected": {
          "answer": "42"
        }
      }
    ]
  }
}`

func TestLoadTestCasesFromDerivedJSONFileWhenCannotReadFile(t *testing.T) {
	defer mustWriteDerivedJSONFile(vanillaContent)
	mustWriteDerivedJSONFile("")
	err := ddt.LoadTestCasesFromDerivedJSONFile(nil)
	if assert.For(t).ThatActual(err).IsNotNil().Passed() {
		assert.For(t).ThatActualString(err.Error()).Equals("unexpected end of JSON input")
	}
}

func TestLoadTestCasesFromDerivedJSONFileWhenSchemaIsInvalid(t *testing.T) {
	defer mustWriteDerivedJSONFile(vanillaContent)
	mustWriteDerivedJSONFile("{}")
	err := ddt.LoadTestCasesFromDerivedJSONFile(nil)
	if assert.For(t).ThatActual(err).IsNotNil().Passed() {
		assert.For(t).ThatActualString(err.Error()).Equals("Cannot find the testFunctions property")
	}
}

func TestLoadTestCasesFromDerivedJSONFileWhenTestNotFound(t *testing.T) {
	mustWriteDerivedJSONFile(vanillaContent)
	expected := "Cannot find test cases for the caller function: TestLoadTestCasesFromDerivedJSONFileWhenTestNotFound"
	err := ddt.LoadTestCasesFromDerivedJSONFile(nil)
	if assert.For(t).ThatActual(err).IsNotNil().Passed() {
		assert.For(t).ThatActualString(err.Error()).Equals(expected)
	}
}

func mustWriteDerivedJSONFile(content string) {
	if err := ioutil.WriteFile("_ddt/ddt_test.json", []byte(content), os.ModePerm); err != nil {
		panic(err)
	}
}
