# AGENTS.md

## Build & Run Commands
- **Build**: `go build`
- **Build (Windows GUI)**: `go build -ldflags '-H=windowsgui -s -w'`
- **Run**: `go run .` (Requires local `piper` binary in `piper/` and ONNX models in `models/`)

## Architecture & Structure
- **Entrypoint**: `main.go` initializes the local file server (`fileserver.go`), webview window, and JS-Go bindings (`readText`, `getModels`, `setModel`, etc.).
- **UI**: HTML/JS/CSS frontend in `webui/` served locally over HTTP.
- **Speech Synthesis**: `synthesis.go` pipes input text to the external Piper TTS binary and streams generated audio chunks.
- **Platform Specifics**: `hidewindow_windows.go` and `hidewindow_linux.go` handle OS-specific window management.

## Gotchas & Quirks
- No unit tests exist in this repository (`*_test.go` files are absent).
- `go.mod` specifies a local `replace` directive for `github.com/webview/webview_go`.
