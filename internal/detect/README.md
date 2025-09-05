# Detect (Sensitive Content Filtering)

This package detects sensitive clipboard entries and prevents them from being written to disk.

## Responsibilities
- Identify secrets (passwords, API keys, JWTs, credit cards, SSH keys).
- Flag entries as `Sensitive` so they can be handled differently.
- Provide both regex-based and entropy-based detection.

## Notes
- Detection helps avoid accidentally persisting secrets.
- Defaults to **RAM-only** storage for sensitive content.
- Users may override via settings in the future.

