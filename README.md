# computer-benchmark
[![Build](https://github.com/Suhaibinator/computer-benchmark/actions/workflows/go-build.yml/badge.svg)](https://github.com/Suhaibinator/computer-benchmark/actions/workflows/go-build.yml)

Some tools for benchmarking computer performance.

This repository contains a set of Go programs used to stress different components of a computer. Each benchmark focuses on a specific workload such as matrix multiplication, memory bandwidth (STREAM), FFT, sorting and cryptography. There are both single-threaded and multithreaded variants.

Run the benchmarks using `go run .` and choose which test to execute from the interactive menu.

## Building the benchmarks

To compile the benchmark binary run:

```bash
go build -o benchmark
```

The above command produces an executable called `benchmark` in the current directory.

## Running the benchmarks

After building, execute the benchmarks with:

```bash
./benchmark
```

Alternatively the code can be run directly without building using:

```bash
go run ./...
```

## Executing unit tests

If unit tests are present they can be run with:

```bash
go test ./...
```

This command recursively executes all tests in the repository.

## Configurable parameters

Benchmark sizes are controlled via constants within the source code. For example:

* `matrixMultiplicationSize` in `multithreaded/matrix_multiplication.go`
* `streamArraySize` in `multithreaded/stream.go`
* `dataSize` in `multithreaded/crypto.go`

Modifying these constants allows changing the amount of work performed by each benchmark. There are currently no environment variables for configuration.
