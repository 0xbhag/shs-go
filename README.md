# shs-go

Yet another simple HTTP server written in Go.

## Installation & Build

Download a pre-built binary from the releases, or install/build from source:

```bash
git clone https://github.com/0xbhag/shs-go.git
cd shs-go
```

**Build locally:**
```bash
go build -ldflags="-s -w" -o shs-go ./cmd/shs-go
```

**Install system-wide:**
```bash
go install -ldflags="-s -w" ./cmd/shs-go
```

## Usage

```bash
./shs-go [flags]
```

### Flags

| Flag | Default | Description |
| :--- | :--- | :--- |
| `-port` | `8080` | Port to run the server on |
| `-path` | `.` (current dir) | Directory path to serve/share |

**Example:**

```bash
./shs-go -port=8000 -path=/path/to/share
```

## Roadmap

- [ ] File upload support (upload files directly to the shared directory)
- [ ] Improved frontend UI design

## Why bother?

Built as a practice project to explore and learn Go's `net/http` package.
