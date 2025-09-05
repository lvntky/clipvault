# ClipVault

🚀 **ClipVault** is a secure, cross-platform clipboard manager written in Go.  
It keeps a **persistent, encrypted history of your clipboard** across restarts — without leaking sensitive data.

---

## ✨ Features

- 🔒 **Encrypted local storage** (Argon2id + AES-GCM via BadgerDB)
- 💾 **Persistent history** — survive reboots and resume where you left off
- 🕵️ **Sensitive content detection** (passwords, tokens, private keys) → stored in RAM only
- ⌨️ **Global hotkeys** to open history, paste, or wipe
- 🧹 **Panic wipe** — instantly clears history + system clipboard
- 🖥️ **Cross-platform GUI** using [Fyne](https://fyne.io/) with tray support
- ⚡ Prebuilt binaries for Linux, macOS, and Windows

---

## 📂 Architecture

ClipVault follows a layered design with clear responsibilities:

```
clipvault/
├── cmd/clipvault/       # Entrypoint (main.go)
└── internal/
    ├── clip/            # Clipboard polling and read/write
    ├── store/           # Encrypted persistent storage (BadgerDB)
    ├── detect/          # Sensitive content detection (regex/entropy)
    ├── ui/              # GUI windows and system tray (Fyne)
    ├── hotkeys/         # Global hotkey registration
    └── config/          # Data directory resolution (portable/system)
```

- **`cmd/clipvault/`** — wires everything together (main app entrypoint).
- **`internal/clip/`** — provides a unified clipboard API using `atotto/clipboard`.
- **`internal/store/`** — handles encrypted persistence (BadgerDB + Argon2id key).
- **`internal/detect/`** — filters sensitive items (passwords, tokens, private keys).
- **`internal/ui/`** — GUI with system tray, history list, paste/wipe actions.
- **`internal/hotkeys/`** — manages global hotkeys (`Ctrl+Shift+V`, etc.).
- **`internal/config/`** — manages OS-specific storage locations and portable mode.

---

## 📦 Dependencies

ClipVault relies on a small set of stable, well-maintained Go libraries:

### 1. [fyne.io/fyne/v2](https://github.com/fyne-io/fyne)
- Provides the **cross-platform GUI** toolkit and **system tray integration**.
- Pure Go, no external runtime required.
- Used for history window, list UI, and tray icon.

### 2. [github.com/atotto/clipboard](https://github.com/atotto/clipboard)
- Cross-platform **clipboard access** (text).
- Simple, widely adopted library.
- Used for reading/writing clipboard content.

### 3. [github.com/dgraph-io/badger/v4](https://github.com/dgraph-io/badger)
- Embedded **key-value database** in pure Go.
- Supports **encryption at rest** with a user-supplied key.
- Used to persist clipboard history securely across reboots.

### 4. [golang.design/x/hotkey](https://pkg.go.dev/golang.design/x/hotkey)
- Provides **global hotkey support** on Linux (X11), macOS, and Windows.
- Used for `Ctrl+Shift+V` (show history), `Ctrl+Shift+C` (clear), and `Ctrl+Shift+X` (panic wipe).

### 5. [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto)
- Provides modern cryptographic primitives.
- Specifically used for **Argon2id** (key derivation) and **AES-GCM/XChaCha20** encryption.

---

## ⚙️ Installation

### From source
```bash
git clone https://github.com/lvntky/clipvault.git
cd clipvault
go build -o clipvault ./cmd/clipvault
```

### Prebuilt binaries
Download from [Releases](https://github.com/lvntky/clipvault/releases).

---

## 🚀 Usage

Start ClipVault:
```bash
./clipvault
```

### Default hotkeys
- `Ctrl+Shift+V` → show history window
- `Ctrl+Shift+C` → clear history
- `Ctrl+Shift+X` → panic wipe (history + clipboard)

### Portable mode
Store history alongside the executable instead of system data dirs:
```bash
./clipvault --portable
```

---

## 🛡️ Security

- **Encryption at rest:** history is encrypted with AES-GCM using a key derived from Argon2id.
- **Sensitive detection:** secrets (passwords, JWTs, SSH keys, credit cards) are not persisted by default.
- **Panic wipe:** instant clearing of history and system clipboard.
- **Portable mode:** optional, but data is still encrypted on disk.

---

## 🗺️ Roadmap

- [ ] Image/file clipboard support
- [ ] Fuzzy search in history
- [ ] Sync across devices (end-to-end encrypted)
- [ ] Customizable hotkeys
- [ ] Export/import history

---

## 🤝 Contributing

Contributions are welcome!  
- Report bugs via [Issues](https://github.com/lvntky/clipvault/issues).  
- Submit PRs for features, fixes, or docs improvements.  

---

## 📜 License

MIT License © 2025 [Levent Kaya](https://github.com/lvntky)
