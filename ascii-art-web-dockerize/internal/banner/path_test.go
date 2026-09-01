package banner

import (
	"errors"
	"testing"
)

// TestGetBannerPath verifies that banner names are correctly mapped to their file paths.
func TestGetBannerPath(t *testing.T) {
	tests := []struct {
		name       string
		bannerName string
		wantPath   string
		wantErr    error
	}{
		{name: "Standard", bannerName: "standard", wantPath: "assets/banners/standard.txt"},
		{name: "Shadow", bannerName: "shadow", wantPath: "assets/banners/shadow.txt"},
		{name: "Thinkertoy", bannerName: "thinkertoy", wantPath: "assets/banners/thinkertoy.txt"},
		{name: "Unknown", bannerName: "unknown", wantErr: ErrUnknownBanner},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBannerPath(tt.bannerName)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("GetBannerPath(%q) error = %v, want %v", tt.bannerName, err, tt.wantErr)
			}
			if got != tt.wantPath {
				t.Fatalf("GetBannerPath(%q) = %q, want %q", tt.bannerName, got, tt.wantPath)
			}
		})
	}
}
