# MY APP
Manage something

## Requirements
App requires Golang 1.8 or later and Glide Package Manager.

## Installation
- Install [Golang](https://golang.org/doc/install)
- Install [Glide](https://glide.sh)

## Build
For building binaries please use make, look at the commands bellow:

```
// Build the binary in your environment
$ make build

// Build using Docker
$ make BUILD=docker build

// Build with another OS. Default Linux
$ make OS=darwin build

// Build with custom version.
$ make APP_VERSION=0.1.0 build

// Build with custom app name.
$ make APP_NAME=myapp build

// Passing all flags
$ make BUILD=docker OS=darwin APP_NAME=myapp APP_VERSION=0.1.0 build

// Clean Up
$ make clean

// Create a docker image with application
$ make pack

// Pack with custom Docker namespace. Default gfgit
$ make DOCKER_NS=mydockerns pack

// Pack with custom version.
$ make APP_VERSION=0.1.0 pack

// Pack with custom app name.
$ make APP_NAME=myapp pack

// Pack passing all flags
$ make APP_NAME=myapp APP_VERSION=0.1.0 DOCKER_NS=mydockerns pack
```

## Develpoment
```
// Running tests
$ make test

// Running tests with coverage. Output coverage file: coverage.html
$ make test-coverage
```
