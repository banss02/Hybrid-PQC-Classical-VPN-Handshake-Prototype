package main_test

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/open-quantum-safe/liboqs-go/oqs"
)

func runKyberBenchmark(alg string) {
	start := time.Now()
	kem := oqs.KeyEncapsulation{}
	if err := kem.Init(alg, nil); err != nil {
		fmt.Printf("%-10s ❌ Init failed: %v\n", alg, err)
		return
	}
	defer kem.Clean()

	publicKey, err := kem.GenerateKeyPair()
	if err != nil {
		fmt.Printf("%-10s ❌ KeyGen failed: %v\n", alg, err)
		return
	}

	ciphertext, sharedEnc, err := kem.EncapSecret(publicKey)
	if err != nil {
		fmt.Printf("%-10s ❌ Encapsulation failed: %v\n", alg, err)
		return
	}

	sharedDec, err := kem.DecapSecret(ciphertext)
	if err != nil {
		fmt.Printf("%-10s ❌ Decapsulation failed: %v\n", alg, err)
		return
	}

	if string(sharedEnc) != string(sharedDec) {
		fmt.Printf("%-10s ❌ Shared secrets mismatch\n", alg)
		return
	}

	combined := sha256.Sum256(sharedDec)
	elapsed := time.Since(start)
	fmt.Printf("%-10s ✅ Success | Time: %4d ms | Hash: %.12x\n", alg, elapsed.Milliseconds(), combined)
}

func main() {
	fmt.Println("🔐 Kyber PQC Benchmark Comparison")
	fmt.Println("-----------------------------------------")
	runKyberBenchmark("Kyber512")
	runKyberBenchmark("Kyber768")
	runKyberBenchmark("Kyber1024")
	fmt.Println("-----------------------------------------")
	fmt.Println("Benchmark complete.")
}
