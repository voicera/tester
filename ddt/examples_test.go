package ddt_test

import (
	"fmt"
	"testing"

	"github.com/voicera/tester/assert"
	"github.com/voicera/tester/ddt"
)

func ExampleLoadTestCasesFromDerivedJSONFile() {
	t := &testing.T{}
	answer := func(q string, timeout int, config map[string]interface{}) (string, error) { return "42", nil }
	var testCases []struct {
		ID    string `json:"id"`
		Input struct {
			Question       string                 `json:"question"`
			TimeoutInHours int                    `json:"timeoutInHours"`
			Config         map[string]interface{} `json:"config"`
		} `json:"input"`
		Expected struct {
			Answer string `json:"answer"`
			Error  error  `json:"error"`
		} `json:"expected"`
	}

	mustWriteDerivedJSONFile(vanillaContent)
	err := ddt.LoadTestCasesFromDerivedJSONFile(&testCases)
	if assert.For(t).ThatActual(err).IsNil().Passed() {
		for _, c := range testCases {
			a, err := answer(c.Input.Question, c.Input.TimeoutInHours, c.Input.Config)
			assert.For(t, c.ID).ThatActual(a).Equals(c.Expected.Answer)
			assert.For(t, c.ID).ThatActual(err).Equals(c.Expected.Error)
			fmt.Println("Actual Answer:", a)
		}
	}

	// Output:
	// Actual Answer: 42
	// Actual Answer: 42
}
