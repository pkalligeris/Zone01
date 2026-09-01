package input

import (
	"reflect"
	"testing"

	"ascii-art/pkg/model"
)

func TestParseArgs_Defaults(t *testing.T) {
	// Test backward compatibility: simple input should yield default config
	args := []string{"hello"}
	expected := &model.Config{
		Input:      "hello",
		BannerFile: "standard",
		Align:      "left",
		Color:      "",
		OutputFile: "",
	}

	got, err := ParseArgs(args)
	if err != nil {
		t.Fatalf("ParseArgs() error = %v", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ParseArgs() = %+v, want %+v", got, expected)
	}
}

func TestParseArgs_Flags(t *testing.T) {
	// Test parsing of all supported flags
	args := []string{
		"--color=red",
		"--align=right",
		"--output=result.txt",
		"hello",
	}

	expected := &model.Config{
		Input:      "hello",
		BannerFile: "standard",
		Align:      "right",
		Color:      "red",
		OutputFile: "result.txt",
	}

	got, err := ParseArgs(args)
	if err != nil {
		t.Fatalf("ParseArgs() error = %v", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ParseArgs() = %+v, want %+v", got, expected)
	}
}

func TestParseArgs_Empty(t *testing.T) {
	_, err := ParseArgs([]string{})
	if err == nil {
		t.Error("ParseArgs(empty) expected error, got nil")
	}
}

func TestParseArgs_Positional(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected *model.Config
		wantErr  bool
	}{
		{
			name: "String and Banner",
			args: []string{"hello", "shadow"},
			expected: &model.Config{
				Input:      "hello",
				BannerFile: "shadow",
				Align:      "left",
				Color:      "",
				OutputFile: "",
			},
		},
		{
			name: "Color Substring and String",
			args: []string{"--color=red", "he", "hello"},
			expected: &model.Config{
				Input:       "hello",
				BannerFile:  "standard",
				Align:       "left",
				Color:       "red",
				ColorSubstr: "he",
				OutputFile:  "",
			},
		},
		{
			name: "Color Input and Banner",
			args: []string{"--color=red", "hello", "shadow"},
			expected: &model.Config{
				Input:       "hello",
				BannerFile:  "shadow",
				Align:       "left",
				Color:       "red",
				ColorSubstr: "",
				OutputFile:  "",
			},
		},
		{
			name: "Color Substring, String, and Banner",
			args: []string{"--color=green", "H", "Hello", "thinkertoy"},
			expected: &model.Config{
				Input:       "Hello",
				BannerFile:  "thinkertoy",
				Align:       "left",
				Color:       "green",
				ColorSubstr: "H",
				OutputFile:  "",
			},
		},
		{
			name:    "Invalid Banner",
			args:    []string{"hello", "invalid"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ParseArgs() = %+v, want %+v", got, tt.expected)
			}
		})
	}
}