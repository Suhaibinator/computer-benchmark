package singlethreaded

import (
	"crypto/aes"
	"crypto/rand"
)

// Cryptographic Algorithm Benchmark (AES Encryption/Decryption)
// LargePlaintextSize configures the workload for CryptoBenchmark when executed
// from the command line.
var LargePlaintextSize = 256 * 10_000_000

// CryptoBenchmark encrypts and decrypts a byte slice of the provided length.
// A default of one million bytes is used if size <= 0.
func CryptoBenchmark(size int) {
	if size <= 0 {
		size = 1_000_000
	}
	key := make([]byte, 32) // 256-bit key
	plaintext := make([]byte, size)

	// Initialize key and plaintext with random data
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	_, err = rand.Read(plaintext)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Use AES in ECB mode (for simplicity)
	encrypted := make([]byte, len(plaintext))
	decrypted := make([]byte, len(plaintext))

	blockSize := block.BlockSize()

	// Encrypt the plaintext
	for i := 0; i < len(plaintext); i += blockSize {
		block.Encrypt(encrypted[i:i+blockSize], plaintext[i:i+blockSize])
	}

	// Decrypt the ciphertext
	for i := 0; i < len(encrypted); i += blockSize {
		block.Decrypt(decrypted[i:i+blockSize], encrypted[i:i+blockSize])
	}
}
