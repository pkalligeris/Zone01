package output

import (
	"os"
	"testing"
)

// TestWriteOutput verifies that WriteOutput creates a file with the correct content.
func TestWriteOutput(t *testing.T) {
	filename := "test_output.txt"
	content := "content"

	// Clean up before and after test
	defer os.Remove(filename)
	os.Remove(filename)

	// Call WriteOutput
	err := WriteOutput(filename, content)
	if err != nil {
		t.Fatalf("WriteOutput failed: %v", err)
	}

	// Assert file exists
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		t.Fatalf("File %s was not created", filename)
	}

	// Assert file permissions are 0644 or 0664 (depending on the system's umask)
	perm := info.Mode().Perm()
	if perm != 0644 && perm != 0664 {
		t.Errorf("Expected permissions 0644 or 0664, got %o", perm)
	}

	// Assert file contains expected content
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(data) != content {
		t.Errorf("Expected content %q, got %q", content, string(data))
	}
}
