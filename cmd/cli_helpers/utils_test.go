package cli_helpers

import (
	"os"
	"testing"
)

func TestIsArgsValids(t *testing.T) {
	os.Setenv("TEST", "1")

	testCases := []struct {
		input    []string
		expected bool
	}{
		{input: []string{"1.18.5", "1.19"}, expected: false},
		{input: []string{"1.18.5"}, expected: true},
		{input: []string{"latest"}, expected: true},
		{input: []string{"1.19"}, expected: true},
		{input: []string{"1.19-alpha"}, expected: false},
		{input: []string{"go1.1.6"}, expected: false},
		{input: []string{"not_a_go_version"}, expected: false},
		{input: []string{"1456.258.88369"}, expected: true},
		{input: []string{"1,18,3"}, expected: false},
		{input: []string{""}, expected: false},
		{input: []string{}, expected: false},
		{input: []string{"1.18.2.6"}, expected: false},
		{input: []string{"1."}, expected: false},
		{input: []string{"2811"}, expected: false},
	}

	for i, test := range testCases {
		out := IsArgsValids(test.input)
		if out != test.expected {
			t.Errorf("Test #%d\ngot: %t\nwant: %t", i, out, test.expected)
		}
	}
}

func TestGetLatestGoVersion(t *testing.T) {
	os.Setenv("TEST", "1")

	latestV, ok := GetLatestGoVersion()
	if !ok {
		t.Fatal("couldn't get latest Go Version")
	}

	if latestV != "1.19" {
		t.Fatalf("\ngot: %s\nwant: %s", latestV, "1.19")
	}
}
