# Hybrid PQC + Classical VPN Handshake Prototype

This project is a **team-based academic implementation** of a hybrid post-quantum cryptography (PQC) and classical VPN handshake built on top of **WireGuard-Go**. It demonstrates how classical elliptic-curve-based VPN handshakes can be extended with post-quantum key encapsulation mechanisms (KEMs) such as **Kyber512**, enabling quantum-resistant secure communication.

## Features

- **Hybrid handshake:** Combines the WireGuard Noise Protocol with Kyber512-based post-quantum key exchange.
- **High-performance implementation:** Developed in Go for efficient VPN communication.
- **Modular cryptography layer:** Supports integration of post-quantum algorithms using Open Quantum Safe (liboqs).
- **Performance evaluation:** Benchmarks key generation, encapsulation, decapsulation, latency, and throughput.
- **WireGuard integration:** Extends the classical WireGuard handshake with hybrid PQC capabilities.

## Technologies

- Go
- WireGuard-Go
- Open Quantum Safe (liboqs)
- Kyber512
- WireGuard
- Tailscale
- Ubuntu Linux
- Bash

## License

This project is released under the MIT License.
