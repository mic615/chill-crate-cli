# chill

Command-line client for [Chill Crate](../chill-crate-api), a simple S3-style object store.

## Install

```bash
go install github.com/mic615/chill-crate-cli/cmd/chill@latest
```

Or build from source:

```bash
go build -o chill ./cmd/chill
```

## Configuration

On first use, `chill` reads `~/.chill-crate.yaml`. Pass `--config <path>` to override.

| Key | Default | Description |
|-----|---------|-------------|
| `api_url` | `http://localhost:8081` | Chill Crate API base URL |
| `user` | — | Active user (set by `chill login`) |
| `current_group_id` | — | The current group id (set by `chill groups use`) |
| `current_group_name` | — | The current group name (set by `chill groups use`) |

## Usage

### Authentication

```bash
chill login <user>
```

Saves `user` to your local config. Token-based auth is coming, but for now this sets the stub user sent on every request.

### Groups

```bash
chill groups list              # list your groups
chill groups create <name>     # create a new group (you're added as a member)
chill groups use               # allows you to select one of your groups to work from
```

### Buckets

```bash
chill buckets list             # list all buckets in your current group
chill buckets create <name>    # create a new bucket in your current group
```

### Objects

```bash
chill objects list <bucket>                       # list all objects in a bucket
chill objects upload <bucket> <filePath>          # upload a file to a bucket
chill objects download <bucket> <filename> <dest> # download an object to a local path
```

## Command Reference

```
chill
├── login <user>
├── groups
│   ├── list
│   ├── create <name>
│   └── use
├── buckets
│   ├── list
│   └── create <name>
└── objects
    ├── list <bucket>
    ├── upload <bucket> <filePath>
    └── download <bucket> <filename> <destination>
```

## Development

Requires Go 1.26+ and a running [chill-crate-api](../chill-crate-api) instance.

```bash
# start the API dependencies
cd ../chill-crate-api && docker-compose up -d

# run the CLI against the local API
go run ./cmd/chill groups list
```

## License

Apache 2.0 — see [LICENSE](LICENSE).
