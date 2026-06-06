package api

import (
	"net/http"
	"os/exec"

	"github.com/ilahazs/yt-webui/backend/config"
)

// DependencyInfo holds status for an external tool.
type DependencyInfo struct {
	Found bool   `json:"found"`
	Path  string `json:"path,omitempty"`
	Error string `json:"error,omitempty"`
}

// HealthData holds the main structure for the health response data field.
type HealthData struct {
	Status       string                    `json:"status"`
	Dependencies map[string]DependencyInfo `json:"dependencies"`
}

// HandleHealth handles health check requests and checks dependencies.
func HandleHealth(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	if r.Method != "GET" {
		SendError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Method not allowed")
		return
	}

	ytDlpInfo := checkDependency(cfg.YtDlpPath)
	ffmpegInfo := checkDependency(cfg.FfmpegPath)

	status := "ok"
	// yt-dlp is a critical dependency, degraded if not found
	if !ytDlpInfo.Found {
		status = "degraded"
	}

	data := HealthData{
		Status: status,
		Dependencies: map[string]DependencyInfo{
			"yt_dlp": ytDlpInfo,
			"ffmpeg": ffmpegInfo,
		},
	}

	SendSuccess(w, http.StatusOK, data)
}

func checkDependency(nameOrPath string) DependencyInfo {
	path, err := exec.LookPath(nameOrPath)
	if err != nil {
		return DependencyInfo{
			Found: false,
			Error: err.Error(),
		}
	}
	return DependencyInfo{
		Found: true,
		Path:  path,
	}
}
