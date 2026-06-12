package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ilahazs/yt-webui/backend/config"
)

func TestHandleHealth(t *testing.T) {
	cfg := &config.Config{
		YtDlpPath:  "bash", // Use a command we know exists on the system and supports --version
		FfmpegPath: "non-existent-tool-xyz-123",
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/health", nil)

	HandleHealth(w, r, cfg)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}

	var resp Response
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Error != nil {
		t.Errorf("expected error to be nil, got %+v", resp.Error)
	}

	dataMap, ok := resp.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("expected data to be map, got %T", resp.Data)
	}

	// Status should be ok since YtDlpPath is "bash" (which is found)
	statusVal := dataMap["status"]
	if statusVal != "ok" {
		t.Errorf("expected status 'ok', got '%v'", statusVal)
	}

	deps, ok := dataMap["dependencies"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected dependencies to be map, got %T", dataMap["dependencies"])
	}

	ytDlp, ok := deps["yt_dlp"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected yt_dlp to be map, got %T", deps["yt_dlp"])
	}

	if ytDlp["found"] != true {
		t.Errorf("expected yt_dlp to be found, got %v", ytDlp["found"])
	}

	versionVal, ok := ytDlp["version"].(string)
	if !ok || versionVal == "" {
		t.Errorf("expected yt_dlp version to be non-empty string, got %T (%v)", ytDlp["version"], ytDlp["version"])
	}
	if !strings.Contains(strings.ToLower(versionVal), "bash") {
		t.Errorf("expected yt_dlp version to contain 'bash', got %q", versionVal)
	}

	ffmpeg, ok := deps["ffmpeg"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected ffmpeg to be map, got %T", deps["ffmpeg"])
	}

	if ffmpeg["found"] != false {
		t.Errorf("expected ffmpeg to not be found, got %v", ffmpeg["found"])
	}

	errVal, ok := ffmpeg["error"].(string)
	if !ok || errVal == "" {
		t.Errorf("expected ffmpeg to have non-empty error field, got %T (%v)", ffmpeg["error"], ffmpeg["error"])
	}
}

func TestHandleHealthMethodNotAllowed(t *testing.T) {
	cfg := &config.Config{
		YtDlpPath:  "bash",
		FfmpegPath: "ffmpeg",
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/api/health", nil)

	HandleHealth(w, r, cfg)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405, got %d", res.StatusCode)
	}

	var resp Response
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Data != nil {
		t.Errorf("expected data to be nil, got %+v", resp.Data)
	}

	if resp.Error == nil {
		t.Fatal("expected error to be non-nil")
	}

	if resp.Error.Code != "method_not_allowed" {
		t.Errorf("expected error code 'method_not_allowed', got '%s'", resp.Error.Code)
	}
}
