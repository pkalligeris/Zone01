package test

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

type goldenCase struct {
	name       string
	input      string
	goldenFile string
}

type featureGoldenCase struct {
	name       string
	args       []string
	goldenFile string
}

func TestGolden(t *testing.T) {
	testCases := []goldenCase{
		{name: "GT-01", input: "hello", goldenFile: "hello.txt"},
		{name: "GT-02", input: "HELLO", goldenFile: "HELLO.txt"},
		{name: "GT-03", input: "HeLlo WoRlD", goldenFile: "mixed_case.txt"},
		{name: "GT-04", input: "1234567890", goldenFile: "numbers.txt"},
		{name: "GT-05", input: "!@#$%^&*()", goldenFile: "special_chars.txt"},
		{name: "GT-06", input: "Hello\\nThere", goldenFile: "multiline.txt"},
		{name: "GT-07", input: "\\n", goldenFile: "newline_only.txt"},
		{name: "GT-08", input: "", goldenFile: "empty.txt"},
		{name: "GT-09", input: "Hello\\n\\nWorld", goldenFile: "double_newline.txt"},
		{name: "GT-10", input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", goldenFile: "all_upper.txt"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Execute the main program as a subprocess
			cmd := exec.Command("go", "run", "./cmd/ascii-art", tc.input)
			cmd.Dir = ".."
			var stdout bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			if err := cmd.Run(); err != nil {
				t.Fatalf("command failed: %v, stderr: %s", err, stderr.String())
			}

			// Read the expected output from the golden file
			goldenPath := filepath.Join("golden", tc.goldenFile)
			expected, err := os.ReadFile(goldenPath)
			if err != nil {
				t.Fatalf("read golden file %q: %v", goldenPath, err)
			}

			// Compare actual stdout with the content of the golden file
			if stdout.String() != string(expected) {
				t.Fatalf("output mismatch\nexpected:\n%s\nactual:\n%s", string(expected), stdout.String())
			}
		})
	}
}

func TestGoldenAlign(t *testing.T) {
	testCases := []featureGoldenCase{
		{name: "GT-13", args: []string{"--align=right", "hello"}, goldenFile: "align_right.txt"},
		{name: "GT-14", args: []string{"--align=center", "hello"}, goldenFile: "align_center.txt"},
		{name: "GT-15", args: []string{"--align=justify", "A B"}, goldenFile: "align_justify.txt"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			args := append([]string{"run", "./cmd/ascii-art"}, tc.args...)
			cmd := exec.Command("go", args...)
			cmd.Dir = ".."
			// Force terminal width to 80 to ensure deterministic output for alignment tests
			cmd.Env = append(os.Environ(), "COLUMNS=80")

			var stdout bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			if err := cmd.Run(); err != nil {
				t.Fatalf("command failed: %v, stderr: %s", err, stderr.String())
			}

			goldenPath := filepath.Join("golden", tc.goldenFile)
			expected, err := os.ReadFile(goldenPath)
			if err != nil {
				t.Fatalf("read golden file %q: %v", goldenPath, err)
			}

			if stdout.String() != string(expected) {
				t.Fatalf("output mismatch\nexpected:\n%s\nactual:\n%s", string(expected), stdout.String())
			}
		})
	}
}

func TestGoldenBannerSelection(t *testing.T) {
	testCases := []featureGoldenCase{
		{name: "GT-11", args: []string{"hello", "shadow"}, goldenFile: "shadow_hello.txt"},
		{name: "GT-12", args: []string{"hello", "thinkertoy"}, goldenFile: "thinkertoy_hello.txt"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			args := append([]string{"run", "./cmd/ascii-art"}, tc.args...)
			cmd := exec.Command("go", args...)
			cmd.Dir = ".."

			var stdout bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			if err := cmd.Run(); err != nil {
				t.Fatalf("command failed: %v, stderr: %s", err, stderr.String())
			}

			goldenPath := filepath.Join("golden", tc.goldenFile)
			expected, err := os.ReadFile(goldenPath)
			if err != nil {
				t.Fatalf("read golden file %q: %v", goldenPath, err)
			}

			if stdout.String() != string(expected) {
				t.Fatalf("output mismatch\nexpected:\n%s\nactual:\n%s", string(expected), stdout.String())
			}
		})
	}
}

// TestOutputFlag verifies that the --output flag writes to a file instead of stdout.
func TestOutputFlag(t *testing.T) {
	outputFile := "result.txt"
	defer os.Remove(outputFile)

	// Execute with --output flag
	cmd := exec.Command("go", "run", "./cmd/ascii-art", "--output="+outputFile, "hello")
	cmd.Dir = ".."
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("command failed: %v, stderr: %s", err, stderr.String())
	}

	// Assert file was created
	if _, err := os.Stat(filepath.Join("..", outputFile)); os.IsNotExist(err) {
		t.Fatalf("Output file %s was not created", outputFile)
	}

	// Read the file content
	data, err := os.ReadFile(filepath.Join("..", outputFile))
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	// Read expected golden file
	expected, err := os.ReadFile("golden/hello.txt")
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	// Compare content
	if string(data) != string(expected) {
		t.Fatalf("Output file content mismatch\nexpected:\n%s\nactual:\n%s", string(expected), string(data))
	}

	// Clean up
	os.Remove(filepath.Join("..", outputFile))
}
