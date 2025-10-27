## Hybrid PQC + Classical VPN Handshake Prototype

This project is a hybrid **post-quantum cryptography (PQC) and classical VPN handshake** implementation built on top of WireGuard-Go.  
It demonstrates how classical elliptic-curve-based VPN handshakes can be extended with **post-quantum key encapsulation mechanisms (KEMs)** such as **Kyber512**, ensuring future-proof security against quantum attacks.

 Features

-  **Hybrid handshake:** Combines classical WireGuard Noise Protocol with PQC (Kyber512).
-  **High-performance:** Implemented in Go for speed and simplicity.
-  **Modular cryptography layer:** Easy to integrate other PQC algorithms from [Open Quantum Safe (liboqs)](https://github.com/open-quantum-safe/liboqs).
-  **Benchmark-ready:** Includes benchmarking for PQC key generation, encapsulation, and decapsulation.
-  **Fully compatible:** Works alongside WireGuard’s classical handshake mechanism.


## License
This project is released under the MIT License.
