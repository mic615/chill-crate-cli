# chill

Command-line client for [Chill Crate](../chill-crate-api), a simple S3-style object store.

## Install

```bash
go install github.com/mic615/chill@latest
```

Or build from source:

```bash
go build -o chill .
```

## Configuration

On first use, `chill` reads `~/.chill-crate.yaml`. Pass `--config <path>` to override.

| Key | Default | Description |
|-----|---------|-------------|
| `api_url` | `http://localhost:8081` | Chill Crate API base URL |
| `user` | — | Active user (set by `chill login`) |

## Usage

### Authentication

```bash
chill login <user>
```

Saves `user` to your local config. Token-based auth is coming — for now this sets the stub user sent on every request.

### Groups

```bash
chill groups list              # list your groups
chill groups create <name>     # create a new group (you're added as a member)
```

## Command Reference

```
chill
├── login <user>
└── groups
    ├── list
    └── create <name>
```

## Development

Requires Go 1.26+ and a running [chill-crate-api](../chill-crate-api) instance.

```bash
# start the API dependencies
cd ../chill-crate-api && docker-compose up -d

# run the CLI against the local API
go run . groups list
```

## License

Apache 2.0 — see [LICENSE](LICENSE).
