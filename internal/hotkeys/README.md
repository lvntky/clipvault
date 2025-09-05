# Hotkeys (Global Keyboard Shortcuts)

This package registers and manages system-wide hotkeys.

## Responsibilities
- Capture global hotkeys to trigger actions (e.g., show history).
- Map cross-platform equivalents (`Ctrl+Shift+V`, etc.).
- Allow customization in future.

## Notes
- Uses [`golang.design/x/hotkey`](https://pkg.go.dev/golang.design/x/hotkey).
- Works on Windows, macOS, and Linux (X11).
- Wayland support is limited — fallback is tray/menu activation.

