package config

import (
	"os"
	"testing"
)

func TestLoadDefaults(t *testing.T) {
	// Clear environments that could interfere
	keysToClear := []string{"BIND_ADDRESS", "PORT", "DATA_DIR", "DOWNLOAD_DIR", "YT_DLP_PATH", "FFMPEG_PATH"}
	origEnv := make(map[string]string)
	for _, key := range keysToClear {
		if val, exists := os.LookupEnv(key); exists {
			origEnv[key] = val
			os.Unsetenv(key)
		}
	}
	defer func() {
		for key, val := range origEnv {
			os.Setenv(key, val)
		}
	}()

	cfg := Load()

	if cfg.BindAddress != ":8080" {
		t.Errorf("expected BindAddress to be ':8080', got %s", cfg.BindAddress)
	}
	if cfg.DataDir != "./data" {
		t.Errorf("expected DataDir to be './data', got %s", cfg.DataDir)
	}
	if cfg.DownloadDir != "./downloads" {
		t.Errorf("expected DownloadDir to be './downloads', got %s", cfg.DownloadDir)
	}
	if cfg.YtDlpPath != "yt-dlp" {
		t.Errorf("expected YtDlpPath to be 'yt-dlp', got %s", cfg.YtDlpPath)
	}
	if cfg.FfmpegPath != "ffmpeg" {
		t.Errorf("expected FfmpegPath to be 'ffmpeg', got %s", cfg.FfmpegPath)
	}
}

func TestLoadCustomEnv(t *testing.T) {
	keysToClear := []string{"BIND_ADDRESS", "PORT", "DATA_DIR", "DOWNLOAD_DIR", "YT_DLP_PATH", "FFMPEG_PATH"}
	origEnv := make(map[string]string)
	for _, key := range keysToClear {
		if val, exists := os.LookupEnv(key); exists {
			origEnv[key] = val
		}
	}
	defer func() {
		// Restore
		for _, key := range keysToClear {
			os.Unsetenv(key)
		}
		for key, val := range origEnv {
			os.Setenv(key, val)
		}
	}()

	os.Setenv("BIND_ADDRESS", "127.0.0.1:9000")
	os.Setenv("DATA_DIR", "/tmp/data")
	os.Setenv("DOWNLOAD_DIR", "/tmp/downloads")
	os.Setenv("YT_DLP_PATH", "/usr/local/bin/yt-dlp")
	os.Setenv("FFMPEG_PATH", "/usr/local/bin/ffmpeg")

	cfg := Load()

	if cfg.BindAddress != "127.0.0.1:9000" {
		t.Errorf("expected BindAddress to be '127.0.0.1:9000', got %s", cfg.BindAddress)
	}
	if cfg.DataDir != "/tmp/data" {
		t.Errorf("expected DataDir to be '/tmp/data', got %s", cfg.DataDir)
	}
	if cfg.DownloadDir != "/tmp/downloads" {
		t.Errorf("expected DownloadDir to be '/tmp/downloads', got %s", cfg.DownloadDir)
	}
	if cfg.YtDlpPath != "/usr/local/bin/yt-dlp" {
		t.Errorf("expected YtDlpPath to be '/usr/local/bin/yt-dlp', got %s", cfg.YtDlpPath)
	}
	if cfg.FfmpegPath != "/usr/local/bin/ffmpeg" {
		t.Errorf("expected FfmpegPath to be '/usr/local/bin/ffmpeg', got %s", cfg.FfmpegPath)
	}
}

func TestLoadPortFallback(t *testing.T) {
	keysToClear := []string{"BIND_ADDRESS", "PORT"}
	origEnv := make(map[string]string)
	for _, key := range keysToClear {
		if val, exists := os.LookupEnv(key); exists {
			origEnv[key] = val
			os.Unsetenv(key)
		}
	}
	defer func() {
		for key, val := range origEnv {
			os.Setenv(key, val)
		}
	}()

	os.Setenv("PORT", "3000")

	cfg := Load()
	if cfg.BindAddress != ":3000" {
		t.Errorf("expected BindAddress to be ':3000', got %s", cfg.BindAddress)
	}
}
