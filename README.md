# mr-menubar

OSX Menubar app displaying active MRs


## Development

 - Go 1.11+
 - Dependencies managed with `go mod`

### Setup

These steps will describe how to setup this project for active development. Adjust paths to your desire.

1. Clone the repository: `git clone git@github.com:kyleshepherd/mr-menubar.git mr-menubar`
2. Build: `make build`
3. 🍻

### Dependencies

Dependencies are managed using `go mod` (introduced in 1.11), their versions
are tracked in `go.mod`.

To add a dependency:
```
go get url/to/origin
```

### Configuration

Configuration can be provided through a toml file, these are loaded
in order from:

- `/etc/mr-menubar/mr-menubar.toml`
- `$HOME/.config/mr-menubar.toml`

Alternatively a config file path can be provided through the
-c/--config CLI flag.

#### Example mr-menubar.toml
```toml
[log]
console = true
level = "debug"  # [debug|info|error]
```
