package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/Suhaibinator/computer-benchmark/multithreaded"
	"github.com/Suhaibinator/computer-benchmark/singlethreaded"
)

// BenchmarkResult holds the execution time of a benchmark
type BenchmarkResult struct {
	Name     string
	Duration time.Duration
}

// Benchmark represents a named benchmark function
type Benchmark struct {
	Name string
	Run  func()
}

var benchmarkList = []Benchmark{
	{"Matrix Multiplication", singlethreaded.MatrixMultiplicationBenchmark},
	{"STREAM Memory Bandwidth", singlethreaded.StreamBenchmark},
	{"Fast Fourier Transform", singlethreaded.FftBenchmark},
	{"Memory Copy and Set", singlethreaded.MemoryCopySetBenchmark},
	{"Sorting Algorithms", singlethreaded.SortingBenchmark},
	{"AES Encryption/Decryption", singlethreaded.CryptoBenchmark},
	{"Multithreaded Matrix Multiplication", multithreaded.MatrixMultiplicationBenchmark},
	{"Multithreaded STREAM Memory Bandwidth", multithreaded.StreamBenchmark},
	{"Multithreaded Fast Fourier Transform", multithreaded.FftBenchmark},
	{"Multithreaded Memory Copy and Set", multithreaded.MemoryCopySetBenchmark},
	{"Multithreaded Sorting Algorithms", multithreaded.SortingBenchmark},
	{"Multithreaded AES Encryption", multithreaded.CryptoBenchmark},
}

func runBenchmark(name string, fn func()) BenchmarkResult {
	start := time.Now()
	fn()
	duration := time.Since(start)
	return BenchmarkResult{Name: name, Duration: duration}
}

func runAllBenchmarks() []BenchmarkResult {
	results := make([]BenchmarkResult, 0, len(benchmarkList))
	for _, b := range benchmarkList {
		fmt.Println("Starting " + b.Name)
		results = append(results, runBenchmark(b.Name, b.Run))
	}
	return results
}

func printResults(results []BenchmarkResult) {
	fmt.Println("\nBenchmark Results:")
	for _, result := range results {
		fmt.Printf("%-40s: %v\n", result.Name, result.Duration)
	}
	fmt.Println()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Select a benchmark to run:")
		fmt.Println("0) Run all benchmarks")
		for i, b := range benchmarkList {
			fmt.Printf("%d) %s\n", i+1, b.Name)
		}
		fmt.Print("q) Quit\n> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.EqualFold(input, "q") {
			return
		}

		if input == "0" {
			results := runAllBenchmarks()
			printResults(results)
			continue
		}

		idx, err := strconv.Atoi(input)
		if err != nil || idx < 1 || idx > len(benchmarkList) {

			fmt.Println("Invalid selection\n")

			continue
		}

		b := benchmarkList[idx-1]
		fmt.Println("Starting " + b.Name)
		result := runBenchmark(b.Name, b.Run)
		printResults([]BenchmarkResult{result})
	}
}
