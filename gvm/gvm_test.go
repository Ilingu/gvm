package gvm

import (
	"fmt"
	appos "gvm/app_os"
	"gvm/gvm/core"
	"gvm/utils"
	"os"
	"testing"
	"time"
)

func init() {
	os.Setenv("TEST", "1")
}

const version = "1.19"

// Test Cmd (no timeout + no cache) --> go test -run ^TestGVMApp$ gvmgvm -v -count=1
func TestGVMWinApp(t *testing.T) {
	appFolder, err := utils.GenerateAppDataPath()
	if err != nil {
		return
	}

	/* Test Scan */
	f, err := os.Create(appFolder + "/1.txt")
	if err != nil {
		t.Fatal("cannot create txt file")
	}
	f.Close()

	noOfDeletes, err := core.ScanAndDelete(time.Now().UnixMilli() + 100) // delete all files
	if err != nil {
		t.Fatal("scan failed")
	}
	if noOfDeletes != 1 {
		t.Fatalf("scan failed\ngot: %d\nwant: %d", noOfDeletes, 1)
	}

	/* Test No Cache Download */
	Godl := MakeGoDownloader(version)
	GoInstaller, err := Godl.DownloadInTemp()
	if err != nil {
		t.Fatal("DownloadTemp failed")
	}
	defer os.Remove(GoInstaller.GetPath())

	tempFileInfo, err := os.Stat(GoInstaller.GetPath())
	if err != nil {
		t.Fatal("cannot read tempFile")
	}
	if os.IsNotExist(err) || tempFileInfo.IsDir() || tempFileInfo.Size() <= 10 {
		t.Fatal("invalid tempFile: DownloadTemp failed")
	}

	/* Test normal Download */
	var GoCachePath string
	err = appos.ExecAccording(
		func() { GoCachePath = appFolder + fmt.Sprintf("/go%s-test.tar.gz", version) }, // Linux
		func() { GoCachePath = appFolder + fmt.Sprintf("/go%s-test.msi", version) },    // Windows
	)
	if err != nil {
		t.Fatal("GoCachePath failed")
		return
	}

	GoInstaller, err = Godl.DownloadInCache()
	if err != nil {
		t.Fatal("DownloadCache failed")
	}
	defer os.Remove(GoInstaller.GetPath())

	if GoInstaller.GetPath() != GoCachePath {
		t.Fatalf("Files paths does not correspond\ngot: %s\nwant: %s", GoCachePath, GoInstaller)
	}

	FileInfo, err := os.Stat(GoCachePath)
	if err != nil {
		t.Fatal("cannot read File")
	}
	if os.IsNotExist(err) || FileInfo.IsDir() || FileInfo.Size() <= 10 {
		t.Fatal("invalid File: Download failed")
	}

	/* Test Installer */
	GoInstaller = MakeGoInstaller(GoCachePath, version)
	err = GoInstaller.Install()
	if err != nil {
		t.Fatal("Install() failed")
	}

	// Verify installation --> Sadly enough but this cannot be executed because when installing the new version it uninstall the Go (old version) but this test still runs on the old go version, so when it's unistalled it kills this test, and so It fails... (but don't worry, my app works, I can verify myself that gox.y.z was successfully installed!)
	userVersion, ok := utils.GetUserGoVersion()
	if !ok {
		t.Fatal("couldn't verify user version")
	}

	if userVersion != "go"+version {
		t.Fatalf("Go Version hasn't been changed.\ngot: %s\nwant: %s", userVersion, "go"+version)
	}
}
