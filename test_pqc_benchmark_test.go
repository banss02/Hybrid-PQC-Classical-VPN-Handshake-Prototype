package main_test

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/open-quantum-safe/liboqs-go/oqs"
)

func runKyberHandshake(alg string) error {
	kem := oqs.KeyEncapsulation{}
	if err := kem.Init(alg, nil); err != nil {
		return err
	}
	defer kem.Clean()

	pubKey, err := kem.GenerateKeyPair()
	if err != nil {
		return err
	}

	ciphertext, sharedEnc, err := kem.EncapSecret(pubKey)
	if err != nil {
		return err
	}

	sharedDec, err := kem.DecapSecret(ciphertext)
	if err != nil {
		return err
	}

	if string(sharedEnc) != string(sharedDec) {
		return fmt.Errorf("shared secrets mismatch")
	}

	_ = sha256.Sum256(sharedDec)
	return nil
}

func BenchmarkKyber(b *testing.B) {
	algos := []string{"Kyber512", "Kyber768", "Kyber1024"}

	for _, alg := range algos {
		b.Run(alg, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if err := runKyberHandshake(alg); err != nil {
					b.Fatalf("%s failed: %v", alg, err)
				}
			}
		})
	}
}
