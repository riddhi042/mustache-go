package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/riddhi042/mustache-go/internal/parser"
	"github.com/riddhi042/mustache-go/internal/renderer"
)

const specPath = "../../../../mustache.js/test/spec/specs/interpolation.json"

type SpecFile struct {
	Overview string     `json:"overview"`
	Tests    []SpecTest `json:"tests"`
}

type SpecTest struct {
	Name     string         `json:"name"`
	Desc     string         `json:"desc"`
	Template string         `json:"template"`
	Data     map[string]any `json:"data"`
	Expected string         `json:"expected"`
}

func RunSpecTests() error {
	data, err := os.ReadFile(specPath)
	if err != nil {
		return err
	}

	var spec SpecFile
	if err := json.Unmarshal(data, &spec); err != nil {
		return err
	}

	fmt.Println("===================================")
	fmt.Println("Official Mustache Spec")
	fmt.Println("===================================")
	fmt.Printf("Overview : %s\n", spec.Overview)
	fmt.Printf("Tests    : %d\n", len(spec.Tests))
	fmt.Println("===================================")

	passed := 0
	failed := 0

	for _, test := range spec.Tests {
		tokens := parser.Parse(test.Template)
		output := renderer.Render(tokens, test.Data)

		if output == test.Expected {
			fmt.Printf("✅ PASS  %s\n", test.Name)
			passed++
		} else {
			fmt.Printf("❌ FAIL  %s\n", test.Name)
			fmt.Println("Template:", test.Template)
			fmt.Println("Expected:", test.Expected)
			fmt.Println("Got     :", output)
			fmt.Println()

			failed++
		}
	}

	fmt.Println("\n---------------------------")
	fmt.Printf("Passed : %d\n", passed)
	fmt.Printf("Failed : %d\n", failed)
	fmt.Println("---------------------------")

	return nil
}
