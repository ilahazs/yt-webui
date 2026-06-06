package config

import (
	"os"
)

type Config struct {
	BindAddress string
	DataDir     string
	DownloadDir string
	YtDlpPath   string
	FfmpegPath  string
}

// Load loads application configuration from environment variables with sane defaults.
func Load() *Config {
	bindAddr := os.Getenv("BIND_ADDRESS")
	if bindAddr == "" {
		port := os.Getenv("PORT")
		if port != "" {
			bindAddr = ":" + port
		} else {
			bindAddr = ":8080"
		}
	}

	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		dataDir = "./data"
	}

	downloadDir := os.Getenv("DOWNLOAD_DIR")
	if downloadDir == "" {
		downloadDir = "./downloads"
	}

	ytDlpPath := os.Getenv("YT_DLP_PATH")
	if ytDlpPath == "" {
		ytDlpPath = "yt-dlp"
	}

	ffmpegPath := os.Getenv("FFMPEG_PATH")
	if ffmpegPath == "" {
		ffmpegPath = "ffmpeg"
	}

	return &Config{
		BindAddress: bindAddr,
		DataDir:     dataDir,
		DownloadDir: downloadDir,
		YtDlpPath:   ytDlpPath,
		FfmpegPath:  ffmpegPath,
	}
}
