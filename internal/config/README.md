# Config (Configuration & Paths)

This package manages application configuration and storage paths.

## Responsibilities
- Resolve system-appropriate data directories:
  - Linux: `~/.local/share/clipvault/`
  - macOS: `~/Library/Application Support/ClipVault/`
  - Windows: `%APPDATA%\ClipVault\`
- Support **portable mode**: keep data alongside the executable.
- Load/save app settings (e.g., TTL, max history size).

## Notes
- Portable mode is activated with `--portable` or `.portable` marker file.
- Ensures data directories have proper permissions (0700).

