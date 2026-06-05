package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// binaryPath returns the path to the compiled test binary.
// We compile once into a temp location and reuse it.
var binaryPath string

func TestMain(m *testing.M) {
	// Build the binary before running integration tests.
	bin := filepath.Join(os.TempDir(), "tetris-optimizer-test-bin")
	cmd := exec.Command("go", "build", "-o", bin, ".")
	cmd.Dir = "." // build from the package directory
	if out, err := cmd.CombinedOutput(); err != nil {
		panic("failed to build binary: " + string(out))
	}
	binaryPath = bin
	code := m.Run()
	os.Remove(bin)
	os.Exit(code)
}

// runBinary runs the compiled binary with the given args and returns stdout output and exit code.
func runBinary(t *testing.T, args ...string) (string, int) {
	t.Helper()
	cmd := exec.Command(binaryPath, args...)
	out, err := cmd.Output()
	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		}
	}
	return string(out), exitCode
}

func testdataPath(name string) string {
	return filepath.Join("..", "..", "testdata", name)
}

// TestMain_NoArgs verifies running with no arguments prints ERROR and exits 1.
func TestMain_NoArgs(t *testing.T) {
	out, code := runBinary(t)
	if !strings.HasPrefix(out, "ERROR") {
		t.Errorf("expected ERROR output, got: %q", out)
	}
	if code == 0 {
		t.Error("expected non-zero exit code, got 0")
	}
}

// TestMain_TwoArgs verifies running with two arguments prints ERROR and exits 1.
func TestMain_TwoArgs(t *testing.T) {
	out, code := runBinary(t, "a.txt", "b.txt")
	if !strings.HasPrefix(out, "ERROR") {
		t.Errorf("expected ERROR output, got: %q", out)
	}
	if code == 0 {
		t.Error("expected non-zero exit code, got 0")
	}
}

// TestMain_FileNotFound verifies a missing file prints ERROR and exits 1.
func TestMain_FileNotFound(t *testing.T) {
	out, code := runBinary(t, "this_file_does_not_exist.txt")
	if !strings.HasPrefix(out, "ERROR") {
		t.Errorf("expected ERROR output, got: %q", out)
	}
	if code == 0 {
		t.Error("expected non-zero exit code, got 0")
	}
}

// TestMain_InvalidChars verifies an invalid-character file prints ERROR.
func TestMain_InvalidChars(t *testing.T) {
	out, code := runBinary(t, testdataPath("invalid_chars.txt"))
	if !strings.HasPrefix(out, "ERROR") {
		t.Errorf("expected ERROR output, got: %q", out)
	}
	if code == 0 {
		t.Error("expected non-zero exit code, got 0")
	}
}

// TestMain_ValidSingle verifies a single-piece file produces valid output and exits 0.
func TestMain_ValidSingle(t *testing.T) {
	out, code := runBinary(t, testdataPath("valid_single.txt"))
	if strings.HasPrefix(out, "ERROR") {
		t.Errorf("unexpected ERROR output: %q", out)
	}
	if code != 0 {
		t.Errorf("expected exit 0, got %d", code)
	}
	if len(strings.TrimSpace(out)) == 0 {
		t.Error("expected non-empty output")
	}
}

// TestMain_ValidMulti verifies a multi-piece file produces valid output and exits 0.
func TestMain_ValidMulti(t *testing.T) {
	out, code := runBinary(t, testdataPath("valid_multi.txt"))
	if strings.HasPrefix(out, "ERROR") {
		t.Errorf("unexpected ERROR output: %q", out)
	}
	if code != 0 {
		t.Errorf("expected exit 0, got %d", code)
	}
	if len(strings.TrimSpace(out)) == 0 {
		t.Error("expected non-empty output")
	}
}
