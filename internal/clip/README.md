# Clip (Clipboard Access)

This package provides a cross-platform interface for accessing the system clipboard.

## Responsibilities
- Read and write clipboard contents (text, images in future).
- Poll or listen for changes and notify the app.
- Abstract away OS differences.

## Notes
- Uses [`github.com/atotto/clipboard`](https://github.com/atotto/clipboard).
- Currently text-only; future work: extend for images and files.
- Polling is the MVP method; event-based hooks may be added later.

