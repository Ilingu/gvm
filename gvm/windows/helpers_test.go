package windows_gvm

import (
	"os"
	"testing"
)

func TestGenerateWinDownloadUrl(t *testing.T) {
	os.Setenv("TEST", "1")

	testCase := []struct {
		input    string
		expected string
	}{
		{input: "1.19", expected: "https://go.dev/dl/go1.19.windows-amd64.msi"},
		{input: "1.18.5", expected: "https://go.dev/dl/go1.18.5.windows-amd64.msi"},
		{input: "something", expected: "https://go.dev/dl/gosomething.windows-amd64.msi"},
	}

	for _, test := range testCase {
		out := generateWinDownloadUrl(test.input)
		if out != test.expected {
			t.Errorf("\ngot: %s\nwant: %s\n", out, test.expected)
		}
	}
}
