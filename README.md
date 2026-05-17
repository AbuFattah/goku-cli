# goku

A fast, lightweight CLI for file format conversion and document persistence. Convert between JSON and YAML instantly, or save documents directly to a PostgreSQL database.

---

## Features

- **Format Conversion** — Convert JSON ↔ YAML from the command line
- **Document Persistence** — Save documents to a PostgreSQL database with a single command
- **Lazy Initialization** — Database connections are only established when a command that requires persistence is invoked
- **Minimal Footprint** — Single self-contained binary, no runtime dependencies

---

## Installation

### via `go install` (requires Go 1.21+)

```bash
go install github.com/abufattah/goku-cli/cmd/goku@latest
```

Make sure `$GOPATH/bin` is in your `$PATH`:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

### via Binary Release

Download the latest binary for your platform from the [Releases](https://github.com/abufattah/goku-cli/releases) page and move it to your `$PATH`:

```bash
curl -L https://github.com/abufattah/goku-cli/releases/latest/download/goku_linux_amd64.tar.gz | tar -xz
sudo mv goku /usr/local/bin/
```

### Build from Source

```bash
git clone https://github.com/abufattah/goku-cli.git
cd goku-cli
make build
# Binary at: bin/app/goku

# Or install directly to /usr/local/bin
sudo make buildcmd
```

---

## Requirements

To use the `save` command, a running PostgreSQL instance is required. The quickest way to spin one up is via Docker:

```bash
make docker-up
make migrateup
```

---

## Usage

### Convert a file

Converts between supported formats. The format is inferred from the file extension.

```bash
goku -i input.json -o output.yaml
goku -i input.yaml -o output.json
```

**Flags:**

| Flag | Short | Description |
|------|-------|-------------|
| `--input` | `-i` | Path to the input file |
| `--output` | `-o` | Path to the output file |

### Save a document

Reads a file and persists it to the configured PostgreSQL database.

```bash
goku save -i ./data/report.json
goku save -i ./config/settings.yaml
```

**Flags:**

| Flag | Short | Description |
|------|-------|-------------|
| `--input` | `-i` | Path to the file to save |

### Help

```bash
goku --help
goku save --help
```

---

## Supported Formats

| Format | Extension |
|--------|-----------|
| JSON   | `.json`   |
| YAML   | `.yaml`, `.yml` |

---

## Configuration

goku is configured via environment variables. Copy `.env.example` to `.env` and adjust as needed:

```env
DATABASE_URL=postgres://postgres:postgres@localhost:5448/goku?sslmode=disable
APP_ENV=development
LOG_LEVEL=info
```

---

## Development

### Prerequisites

- [Go 1.21+](https://go.dev/dl/)
- [Docker](https://docs.docker.com/get-docker/)
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [sqlc](https://sqlc.dev/)
- [golangci-lint](https://golangci-lint.run/) *(optional, for linting)*

### Available Make Targets

```bash
make build         # Build binary to bin/app/goku
make buildcmd      # Build and install to /usr/local/bin/goku
make run           # Run without building
make docker-up     # Start PostgreSQL via Docker Compose
make docker-down   # Stop Docker Compose services
make migrateup     # Run database migrations (up)
make migratedown   # Roll back database migrations (down)
make sqlc-gen      # Regenerate sqlc query code
make test          # Run all tests with race detection
make lint          # Run golangci-lint
make tidy          # Run go mod tidy
```

### Project Structure

```
.
├── cmd/                        # Application entry point
├── config/                     # Configuration loading
├── internal/
│   ├── cli/
│   │   ├── root/               # Root command (conversion)
│   │   └── save/               # Save command (persistence)
│   ├── gokuapp/
│   │   ├── domain/             # Domain types
│   │   ├── repository/         # Database access layer
│   │   └── service/            # Business logic
│   └── platform/
│       ├── container/          # Dependency wiring
│       ├── database/           # Migrations, queries, sqlc output
│       └── fileutil/           # File I/O and format detection
└── docker-compose.yml
```

---

## License

[MIT](LICENSE)