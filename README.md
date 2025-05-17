# computer-benchmark

This repository contains a set of Go programs used to stress different components of a computer. Each benchmark focuses on a specific workload such as matrix multiplication, memory bandwidth (STREAM), FFT, sorting and cryptography. There are both single-threaded and multithreaded variants. When the program starts it presents an interactive CLI so you can choose which benchmark to run.

## Building the benchmarks

To compile the benchmark binary run:

```bash
go build -o benchmark
```

This produces an executable called `benchmark` in the current directory.

## Running the benchmarks

After building, execute the benchmarks with:

```bash
./benchmark
```

You can also run the code directly without building using:

```bash
go run .
```

Once started, a menu lets you select an individual benchmark or run them all. Enter `q` to quit the program.

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
