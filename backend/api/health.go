package api

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/ilahazs/yt-webui/backend/config"
)

// DependencyInfo holds status for an external tool.
type DependencyInfo struct {
	Found   bool   `json:"found"`
	Path    string `json:"path,omitempty"`
	Version string `json:"version,omitempty"`
	Error   string `json:"error,omitempty"`
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

	ytDlpInfo := checkDependency(cfg.YtDlpPath, "--version")
	ffmpegInfo := checkDependency(cfg.FfmpegPath, "-version")

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

func checkDependency(nameOrPath, versionArg string) DependencyInfo {
	path, err := exec.LookPath(nameOrPath)
	if err != nil {
		return DependencyInfo{
			Found: false,
			Error: fmt.Sprintf("executable not found in PATH: %v", err),
		}
	}

	cmd := exec.Command(path, versionArg)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return DependencyInfo{
			Found: false,
			Path:  path,
			Error: fmt.Sprintf("failed to execute version command: %v", err),
		}
	}

	versionStr := strings.TrimSpace(string(output))
	if lines := strings.Split(versionStr, "\n"); len(lines) > 0 {
		versionStr = strings.TrimSpace(lines[0])
	}

	return DependencyInfo{
		Found:   true,
		Path:    path,
		Version: versionStr,
	}
}
