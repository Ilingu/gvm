package utils

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func TestGetUserDir(t *testing.T) {
	os.Setenv("TEST", "1")

	homeDir, err := GetUserDir()
	if err != nil {
		t.Fatal("couldn't get user's home dir", err)
	}

	if homeDir != `C:\Users\Iling` {
		t.Fatalf("\ngot: %s\nwant: %s", homeDir, `C:\Users\Iling`)
	}
}

func TestGenerateAppDataPath(t *testing.T) {
	os.Setenv("TEST", "1")

	appDir, err := GenerateAppDataPath()
	if err != nil {
		t.Fatal("couldn't get appDir", err)
	}

	expected := `C:\Users\Iling\AppData\Roaming\gvm-windows`
	if appDir != expected {
		t.Fatalf("\ngot: %s\nwant: %s", appDir, expected)
	}
}

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
		out := GenerateWinDownloadUrl(test.input)
		if out != test.expected {
			t.Errorf("\ngot: %s\nwant: %s\n", out, test.expected)
		}
	}
}

func TestGenerateSourceDownloadUrl(t *testing.T) {
	os.Setenv("TEST", "1")

	testCase := []struct {
		input    string
		expected string
	}{
		{input: "1.19", expected: "https://go.dev/dl/go1.19.src.tar.gz"},
		{input: "1.18.5", expected: "https://go.dev/dl/go1.18.5.src.tar.gz"},
		{input: "something", expected: "https://go.dev/dl/gosomething.src.tar.gz"},
	}

	for _, test := range testCase {
		out := GenerateSourceDownloadUrl(test.input)
		if out != test.expected {
			t.Errorf("\ngot: %s\nwant: %s\n", out, test.expected)
		}
	}
}

// Test Cmd (no timeout + no cache) --> go test -run ^TestUntar$ gvm-windows/gvm/utils -v -count=1
func TestUntar(t *testing.T) {
	os.Setenv("TEST", "1")

	// Mock (prepare the ground)
	GoOneDotNineteenUrl := GenerateSourceDownloadUrl("1.19")
	resp, err := http.Get(GoOneDotNineteenUrl)
	if err != nil {
		t.Fatal("Couldn't fetch source: ", err)
	}
	defer resp.Body.Close()

	tempFile, err := os.CreateTemp("", "go1.19-*.src.tar.gz")
	if err != nil {
		t.Fatal("Couldn't create temp file: ", err)
	}

	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		t.Fatal("Couldn't populate temp source file: ", err)
	}

	defer os.Remove(tempFile.Name()) // Second Remove temp source file
	defer tempFile.Close()           // First Close temp source file

	dirPath, err := os.MkdirTemp("", "go1.19-source-*")
	if err != nil {
		t.Fatal("Couldn't create temp dir: ", err)
	}
	defer os.RemoveAll(dirPath)

	t.Logf("temp file: %s", tempFile.Name())
	t.Logf("temp dir: %s", dirPath)

	// Actual Test
	err = Untar(tempFile.Name(), dirPath)
	if err != nil {
		t.Fatal("Couldn't untar file: ", err)
	}

	// See If untar really succeed
	fileInfo, err := os.Stat(dirPath + "\\VERSION")
	if os.IsNotExist(err) || err != nil || fileInfo.IsDir() {
		t.Fatal("Couldn't find version file: ", err)
	}

	versionFile, err := os.Open(dirPath + "\\VERSION")
	if err != nil {
		t.Fatal("Couldn't find version file: ", err)
	}

	rawVersion, err := io.ReadAll(versionFile)
	if err != nil || len(rawVersion) == 0 {
		t.Fatal("Couldn't read version file: ", err)
	}
	version := string(rawVersion)

	if version != "go1.19" {
		t.Errorf("\ngot: %s\nwant: %s\n", version, "go1.19")
	}
}

func TestIsTestEnv(t *testing.T) {
	os.Setenv("TEST", "1")
	if !IsTestEnv() {
		t.Fatal("failed, expected true, got false")
	}
}
