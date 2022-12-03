package console_test

import (
	console "gvm/console"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func init() {
	os.Setenv("TEST", "1")
}

// High order function
func hijackStdout(toExec func()) string {
	// save std's
	stdOut := os.Stdout
	stdErr := os.Stderr

	// hijack std's
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = w
	log.SetOutput(w)

	// Exec Funcs
	toExec()

	// Close hijack
	_ = w.Close()

	// get results
	result, _ := io.ReadAll(r)
	output := string(result)

	// reset the std's
	os.Stdout = stdOut
	os.Stderr = stdErr
	log.SetOutput(os.Stderr)

	return output
}

func TestLogMsg(t *testing.T) {
	res := hijackStdout(func() {
		console.LogMsg("LogMsg test", console.NEUTRAL)
		console.Success("test Success")
		console.Log("test Log")
		console.Warn("test Warn")
		console.Error("test Error")
		console.Neutral("test Neutral")
	})

	wantedLog := []string{"[Neutral] \x1b[0m \x1b[37m LogMsg test", "[Success] \x1b[0m \x1b[32m test Success", "[Info] \x1b[0m \x1b[36m test Log", "[Warning] \x1b[0m \x1b[33m test Warn", "[Error] \x1b[0m \x1b[31m test Error", "[Neutral] \x1b[0m \x1b[37m test Neutral"}
	for _, log := range wantedLog {
		if !strings.Contains(res, log) {
			t.Error("missing", log)
		}
	}
}
