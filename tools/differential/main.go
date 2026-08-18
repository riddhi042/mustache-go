package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/riddhi042/mustache-go/internal/parser"
	"github.com/riddhi042/mustache-go/internal/renderer"
)

type TestCase struct {
	Name     string
	Template string
	DataJSON string
	Data     map[string]any
}

func Compare(tc TestCase) (bool, error) {
	// Go implementation
	tokens := parser.Parse(tc.Template)
	goOutput := renderer.Render(tokens, tc.Data)

	// JavaScript implementation
	cmd := exec.Command(
		"node",
		"render.js",
		tc.Template,
		tc.DataJSON,
	)

	cmd.Dir = "tools/differential"

	jsBytes, err := cmd.Output()
	if err != nil {
		return false, err
	}

	jsOutput := strings.TrimSpace(string(jsBytes))
	goOutput = strings.TrimSpace(goOutput)

	if goOutput == jsOutput {
		return true, nil
	}

	fmt.Printf("❌ FAIL  %-20s\n", tc.Name)
	fmt.Println("Template:")
	fmt.Println(tc.Template)

	fmt.Println("Go:")
	fmt.Println(goOutput)

	fmt.Println("JS:")
	fmt.Println(jsOutput)
	fmt.Println()

	return false, nil
}

func main() {
	err := RunSpecTests()
	if err != nil {
		fmt.Println("Error running spec tests:", err)
	}
}

func runManualTests() {
	tests := []TestCase{
		{
			Name:     "Variable",
			Template: "Hello {{name}}",
			DataJSON: `{"name":"Riddhi"}`,
			Data: map[string]any{
				"name": "Riddhi",
			},
		},
		{
			Name:     "Escaping",
			Template: "Hello {{name}}",
			DataJSON: `{"name":"<b>Riddhi</b>"}`,
			Data: map[string]any{
				"name": "<b>Riddhi</b>",
			},
		},
		{
			Name:     "Triple Mustache",
			Template: "Hello {{{name}}}",
			DataJSON: `{"name":"<b>Riddhi</b>"}`,
			Data: map[string]any{
				"name": "<b>Riddhi</b>",
			},
		},
	}

	passed := 0
	failed := 0

	fmt.Println("Running Manual Differential Tests...\n")

	for _, tc := range tests {
		ok, err := Compare(tc)
		if err != nil {
			fmt.Printf("[ERROR] %s\n%v\n\n", tc.Name, err)
			failed++
			continue
		}

		if ok {
			fmt.Printf("✅ PASS  %-20s\n", tc.Name)
			passed++
		} else {
			failed++
		}
	}

	fmt.Println("\n---------------------------")
	fmt.Printf("Passed : %d\n", passed)
	fmt.Printf("Failed : %d\n", failed)
	fmt.Println("---------------------------")
}
