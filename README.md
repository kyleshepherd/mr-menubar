# mr-menubar

OSX Menubar app displaying active MRs

## Development

- Go 1.11+
- Dependencies managed with `go mod`

### Setup

These steps will describe how to setup this project for active development. Adjust paths to your desire.

1. Clone the repository: `git clone git@github.com:kyleshepherd/mr-menubar.git mr-menubar`
2. Build: `make`
3. Run `cp -R MRBar.app /Applications/`
4. In Applications, double-click MRBar to open
5. Click Gitlab icon in menu bar
6. Click "Set Token" in menu
7. Enter your Gitlab Access Token (must have `read_api` scope at minimum)

### Dependencies

Dependencies are managed using `go mod` (introduced in 1.11), their versions
are tracked in `go.mod`.

To add a dependency:

```
go get url/to/origin
```
