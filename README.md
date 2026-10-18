# PQC-SecureLink: Post-Quantum Cryptography in Go

This project demonstrates a Quantum-Resistant Key Encapsulation Mechanism (KEM) using the NIST-standardized **ML-KEM (Kyber-768)** algorithm. It is designed to protect critical infrastructure communications against the "Store Now, Decrypt Later" threat posed by future quantum computers.

## Mathematical Core (LWE)
The security of this implementation relies on the **Learning With Errors (LWE)** problem over module lattices. By introducing deterministic mathematical noise into the public key matrix, the algorithm neutralizes Shor's Algorithm, rendering quantum superposition and period-finding techniques useless.

## Features & Workflow
- **Language:** Go (Golang)
- **Library:** Cloudflare CIRCL (`github.com/cloudflare/circl`)
- **Key Generation:** Generates 1184-byte quantum-resistant Public Keys.
- **Encapsulation:** Creates a 1088-byte ciphertext capsule encapsulating a 32-byte shared secret.
- **Decapsulation:** Decrypts the capsule to establish a secure AES-256 symmetric key matrix for secure SCADA/ICS communication.