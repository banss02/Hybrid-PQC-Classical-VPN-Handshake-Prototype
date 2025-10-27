package main_test

import (
	"crypto/sha256"
	"fmt"
	"testing"
	"time"

	"github.com/open-quantum-safe/liboqs-go/oqs"
)

func BenchmarkPQCHandshake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		start := time.Now()
		fmt.Println("🔐 Starting PQC handshake test (Kyber512 + SHA256)...")

		// Initialize Kyber512 KEM
		kem := oqs.KeyEncapsulation{}
		if err := kem.Init("Kyber512", nil); err != nil {
			b.Fatalf("❌ Kyber initialization failed: %v", err)
		}
		defer kem.Clean()

		// Generate keypair
		publicKey, err := kem.GenerateKeyPair()
		if err != nil {
			b.Fatalf("❌ Key pair generation failed: %v", err)
		}

		// Encapsulate (simulate peer side)
		ciphertext, sharedSecretEnc, err := kem.EncapSecret(publicKey)
		if err != nil {
			b.Fatalf("❌ Encapsulation failed: %v", err)
		}

		// Decapsulate (simulate our side)
		sharedSecretDec, err := kem.DecapSecret(ciphertext)
		if err != nil {
			b.Fatalf("❌ Decapsulation failed: %v", err)
		}

		// Verify equality
		if string(sharedSecretEnc) != string(sharedSecretDec) {
			b.Fatal("❌ Shared secrets do not match!")
		}

		// Combine and hash
		finalKey := sha256.Sum256(sharedSecretDec)
		elapsed := time.Since(start)

		fmt.Println("✅ PQC handshake complete! Combined key hash:", fmt.Sprintf("%x", finalKey))
		fmt.Printf("⏱️  Handshake time: %v ms\n", elapsed.Milliseconds())
	}
}
