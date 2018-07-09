# MY APP
Manage something

## Development

### Requirements

- Install [Go](https://golang.org)
- Install [go dep](https://github.com/golang/dep)

### Makefile
```sh
// Build a beta version of kongo
$ make build

// Clean up
$ make clean

// Creates folders and download dependencies
$ make configure

//Run tests and generates html coverage file
make cover

// Download project dependencies
make depend

// Format all go files
make fmt

//Run linters
make lint

// Run tests
make test
```
