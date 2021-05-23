# Echo

The echo project implements a TCP echo server like it is specified in the [RFC 862](https://datatracker.ietf.org/doc/html/rfc862).

## Project structure

- **cmd/echo**: contains the main executable of the project.
- **config**: contains the general configuration of the project.
- **internal/app**: contains the logic of the echo process (reads data and writes it back).
- **internal/infra**: contains the logic to accept connections.

## Configuration

By default, the echo project is listening in port 7 in the interface 127.0.0.1, and it is using a buffer of 100 bytes to read information and write it back.

All these parameters can be modified using the following flags:

- `-port`: changes the port where the echo project is listening to.
- `-host`: changes the interface to which is the echo project doing the port binding.
- `-buffer-size`: changes the size of the buffer used to read information and write it back.

## How to run the project

The following dependencies are required for the project to be executed, `make`, and `go`.

```bash
make run
```

## How to test the project

Besides, the mandatory dependencies to run the project we will need the following extra dependencies to run test and run linting checks in the project. `golangci-lint`, and `moq`. They can be installed with the following:

```bash
make tools
```

To run the linting checks it is similar to the previous command:

```bash
make lint
```

### Unit test

In order to run the unit test you just need the following command:

```bash
make test
```
