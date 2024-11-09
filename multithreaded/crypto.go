package multithreaded

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"runtime"
	"sync"
)

const dataSize = 3 * 1_000_000_000 // Fixed data size

// CryptoBenchmark performs multithreaded AES encryption
func CryptoBenchmark() {
	key := make([]byte, 32) // AES-256
	rand.Read(key)

	data := make([]byte, dataSize)
	rand.Read(data)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	numWorkers := runtime.NumCPU() * 3
	var wg sync.WaitGroup

	chunkSize := dataSize / numWorkers
	remainder := dataSize % numWorkers

	encrypted := make([]byte, dataSize)

	startIndex := 0
	for w := 0; w < numWorkers; w++ {
		endIndex := startIndex + chunkSize
		if w < remainder {
			endIndex++
		}
		if endIndex > dataSize {
			endIndex = dataSize
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			iv := make([]byte, aes.BlockSize)
			rand.Read(iv)
			stream := cipher.NewCTR(block, iv)
			stream.XORKeyStream(encrypted[start:end], data[start:end])
		}(startIndex, endIndex)

		startIndex = endIndex
		if startIndex >= dataSize {
			break
		}
	}
	wg.Wait()
}
