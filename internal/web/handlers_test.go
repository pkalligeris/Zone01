package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAPIASCIIArtHandlerSuccess(t *testing.T) {
	body := bytes.NewBufferString(`{"text":"hi","banner":"standard","align":"left","color":"red"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/ascii-art", body)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	apiASCIIArtHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
	if got := rec.Header().Get("Content-Type"); !strings.Contains(got, "application/json") {
		t.Fatalf("expected JSON content type, got %q", got)
	}

	var resp asciiArtResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if resp.Result == "" {
		t.Fatal("expected rendered result")
	}
	if resp.Error != "" {
		t.Fatalf("unexpected error: %q", resp.Error)
	}
}

func TestAPIASCIIArtHandlerRejectsInvalidMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/ascii-art", nil)
	rec := httptest.NewRecorder()

	apiASCIIArtHandler(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, rec.Code)
	}
}

func TestAPIASCIIArtHandlerRejectsMalformedJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/ascii-art", strings.NewReader(`{"text":"hi"`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	apiASCIIArtHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestAPIASCIIArtHandlerRejectsInvalidBanner(t *testing.T) {
	body := bytes.NewBufferString(`{"text":"hi","banner":"invalid"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/ascii-art", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	apiASCIIArtHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestAPIASCIIArtHandlerRejectsInvalidASCII(t *testing.T) {
	body := bytes.NewBufferString("{\"text\":\"hi\\u00e9\",\"banner\":\"standard\"}")
	req := httptest.NewRequest(http.MethodPost, "/api/ascii-art", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	apiASCIIArtHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestAPIASCIIArtHandlerRejectsInvalidColor(t *testing.T) {
	body := bytes.NewBufferString(`{"text":"hi","banner":"standard","color":"not-a-color"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/ascii-art", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	apiASCIIArtHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestAPIASCIIArtHandlerRejectsInvalidAlign(t *testing.T) {
	body := bytes.NewBufferString(`{"text":"hi","banner":"standard","align":"diagonal"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/ascii-art", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	apiASCIIArtHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestASCIIArtHandlerHTMLStillWorks(t *testing.T) {
	form := url.Values{
		"text":   {"hi"},
		"banner": {"standard"},
		"color":  {"red"},
	}
	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	asciiArtHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "<pre") {
		t.Fatal("expected rendered HTML response")
	}
	if strings.Contains(body, "\u001b[") {
		t.Fatal("did not expect ANSI escape sequences in HTML response")
	}
	if !strings.Contains(body, `style="color: red;"`) {
		t.Fatal("expected CSS color styling in HTML response")
	}
}

func TestExportHandlerSuccess(t *testing.T) {
	form := url.Values{
		"text":   {"hi"},
		"banner": {"standard"},
	}
	req := httptest.NewRequest(http.MethodPost, "/export", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	exportHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	// Verify required HTTP headers
	ct := rec.Header().Get("Content-Type")
	if !strings.Contains(ct, "text/plain") {
		t.Fatalf("expected Content-Type text/plain, got %q", ct)
	}

	cd := rec.Header().Get("Content-Disposition")
	if !strings.Contains(cd, "attachment") || !strings.Contains(cd, "ascii_art.txt") {
		t.Fatalf("expected Content-Disposition attachment with filename, got %q", cd)
	}

	cl := rec.Header().Get("Content-Length")
	if cl == "" || cl == "0" {
		t.Fatal("expected non-zero Content-Length")
	}

	body := rec.Body.String()
	if body == "" {
		t.Fatal("expected non-empty body")
	}
}

func TestExportHandlerRejectsGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/export", nil)
	rec := httptest.NewRecorder()

	exportHandler(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, rec.Code)
	}
}

func TestExportHandlerRejectsEmptyText(t *testing.T) {
	form := url.Values{
		"text":   {""},
		"banner": {"standard"},
	}
	req := httptest.NewRequest(http.MethodPost, "/export", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	exportHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}
