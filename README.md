# go-hexagonal

## Contributing

| Command         | Description                                       | Dependencies                                                  |
|-----------------|---------------------------------------------------|---------------------------------------------------------------|
| `make build`    | Build project                                     | `go`                                                          |
| `make test`     | Run unit tests                                    | `go`                                                          |
| `make fmt`      | Format files                                      | `gofmt`, `gofumpt` and `goimports`                            |
| `make lint`     | Check files                                       | `golangci-lint` and `go-arch-lint`                            |
| `make dod`      | (Definition of Done) Format files and check files | Same as `make build`, `make test`, `make fmt` and `make lint` | 
| `make install`  | Install all dependencies                          | `go`, `curl` and `git`                                        |
| `make mocks`    | Generate mocks                                    | `go` and `mockery`                                            |
| `make docs`     | Run docsify docs server local                     | `docsify`                                                     |
| `make godoc`    | Run godoc server local                            | `godoc`                                                       |