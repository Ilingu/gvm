package utils

import (
	"os"
	"runtime"
	"testing"
)

func TestGetGOROOT(t *testing.T) {
	os.Setenv("TEST", "1")

	goroot, ok := GetGOROOT()
	if !ok {
		t.Fatal("couldn't get goroot")
	}

	if goroot != `C:\Program Files\Go` {
		t.Fatalf("\ngot: %s\nwant: %s", goroot, `C:\Program Files\Go`)
	}
}

func TestGetUserGoVersion(t *testing.T) {
	os.Setenv("TEST", "1")

	goversion, ok := GetUserGoVersion()
	if !ok {
		t.Fatal("couldn't get goversion")
	}

	if goversion != runtime.Version() {
		t.Fatalf("\ngot: %s\nwant: %s", goversion, runtime.Version())
	}
}

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

func TestIsTestEnv(t *testing.T) {
	os.Setenv("TEST", "1")
	if !IsTestEnv() {
		t.Fatal("failed, expected true, got false")
	}
}
