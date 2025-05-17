package singlethreaded

import (
	"crypto/aes"
	"crypto/rand"
)

// Cryptographic Algorithm Benchmark (AES Encryption/Decryption)
func CryptoBenchmark() {
	key := make([]byte, 32) // 256-bit key
	var plaintextSize int64 = 256 * 10_000_000
	plaintext := make([]byte, int(plaintextSize))

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
