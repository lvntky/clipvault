# Store (Encrypted Persistence)

This package manages persistent storage of clipboard history.

## Responsibilities
- Persist clipboard history securely across reboots.
- Encrypt data at rest with AES-GCM.
- Derive encryption key using Argon2id.
- Apply TTL (time-to-live) for auto-expiry.

## Notes
- Uses [`github.com/dgraph-io/badger/v4`](https://github.com/dgraph-io/badger).
- Sensitive items are skipped or marked in-memory only.
- Default storage path is OS-specific data directory, unless `--portable` is used.

