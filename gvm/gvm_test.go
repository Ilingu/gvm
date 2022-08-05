package gvm

import (
	"fmt"
	cmdUtils "gvm-windows/cmd/utils"
	"gvm-windows/gvm/utils"
	"os"
	"testing"
	"time"
)

// Test Cmd (no timeout + no cache) --> go test -run ^TestGVMApp$ gvm-windows/gvm -v -count=1
func TestGVMApp(t *testing.T) {
	os.Setenv("TEST", "1")

	version := "1.18.5"
	appFolder, err := utils.GenerateAppDataPath()
	if err != nil {
		return
	}

	/* Test Scan */
	f, err := os.Create(appFolder + "\\1.txt")
	if err != nil {
		t.Fatal("cannot create txt file")
	}
	f.Close()

	noOfDeletes, err := ScanAndDelete(time.Now().UnixMilli() + 100) // delete all files
	if err != nil {
		t.Fatal("scan failed")
	}
	if noOfDeletes != 1 {
		t.Fatalf("scan failed\ngot: %d\nwant: %d", noOfDeletes, 1)
	}

	/* Test No Cache Download */
	Godl := MakeGoDownloader(version)
	GoMsiExecutableTest, ok := Godl.DownloadTempMSI()
	if !ok {
		t.Fatal("DownloadTempMSI failed")
	}
	defer os.Remove(GoMsiExecutableTest)

	tempFileInfo, err := os.Stat(GoMsiExecutableTest)
	if err != nil {
		t.Fatal("cannot read tempFile")
	}
	if os.IsNotExist(err) || tempFileInfo.IsDir() || tempFileInfo.Size() <= 10 {
		t.Fatal("invalid tempFile: DownloadTempMSI failed")
	}

	/* Test normal Download */
	GoMsiExecutableTest = appFolder + fmt.Sprintf("\\go%s-test.msi", version)
	GoMsiExecutable, ok := Godl.DownloadMSI()
	if !ok {
		t.Fatal("DownloadMSI failed")
	}

	defer os.Remove(GoMsiExecutable)
	if GoMsiExecutableTest != GoMsiExecutable {
		t.Fatalf("Files paths does not correspond\ngot: %s\nwant: %s", GoMsiExecutable, GoMsiExecutableTest)
	}

	msiFileInfo, err := os.Stat(GoMsiExecutable)
	if err != nil {
		t.Fatal("cannot read msiFile")
	}
	if os.IsNotExist(err) || msiFileInfo.IsDir() || msiFileInfo.Size() <= 10 {
		t.Fatal("invalid msiFile: DownloadMSI failed")
	}

	/* Test Installer */
	GoInstaller := MakeGoInstaller(GoMsiExecutable)
	ok = GoInstaller.InstallAsMSI()
	if !ok {
		t.Fatal("InstallAsMSI failed")
	}

	// Verify installation --> Sadly enough but this cannot be executed because when installing the new version it uninstall the Go (old version) but this test still runs on the old go version, so when it's unistalled it kills this test, and so It fails... (but don't worry, my app works, I can verify myself that gox.y.z was successfully installed!)
	userVersion, ok := cmdUtils.GetUserGoVersion()
	if !ok {
		t.Fatal("couldn't verify user version")
	}

	if userVersion != "go"+version {
		t.Fatalf("Go Version hasn't been changed.\ngot: %s\nwant: %s", userVersion, "go"+version)
	}
}
