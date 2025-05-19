package multithreaded

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"runtime"
	"sync"
)

const DefaultCryptoSize = 1_000_000

// CryptoBenchmark performs multithreaded AES encryption on a buffer of the
// given size. When size is non-positive, DefaultCryptoSize is used.
func CryptoBenchmark(size int) {
	if size <= 0 {
		size = DefaultCryptoSize
	}
	key := make([]byte, 32) // AES-256
	rand.Read(key)

	data := make([]byte, size)
	rand.Read(data)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	numWorkers := runtime.NumCPU() * 3
	var wg sync.WaitGroup

	chunkSize := size / numWorkers
	remainder := size % numWorkers

	encrypted := make([]byte, size)

	startIndex := 0
	for w := 0; w < numWorkers; w++ {
		endIndex := startIndex + chunkSize
		if w < remainder {
			endIndex++
		}
		if endIndex > size {
			endIndex = size
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
		if startIndex >= size {
			break
		}
	}
	wg.Wait()
}
