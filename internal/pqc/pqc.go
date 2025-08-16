package pqc

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "io"

    kyber "github.com/cloudflare/circl/kem/kyber/kyber768"
)

// PQCManager manages key pairs and provides encrypt/decrypt using Kyber KEM and AES-GCM.
type PQCManager struct {
    priv kyber.PrivateKey
    pub  kyber.PublicKey
}

// NewPQC generates a new Kyber key pair and returns a PQCManager.
func NewPQC() (*PQCManager, error) {
    scheme := kyber.Scheme()
    pub, priv, err := scheme.GenerateKeyPair()
    if err != nil {
        return nil, err
    }
    return &PQCManager{priv: priv, pub: pub}, nil
}

// Encrypt encrypts the provided data using AES-GCM with a shared secret derived via Kyber KEM.
// It returns the encrypted data and the Kyber ciphertext needed for decryption.
func (m *PQCManager) Encrypt(data []byte) ([]byte, []byte, error) {
    // Encapsulate to get a shared secret and Kyber ciphertext
    ct, ss, err := m.pub.Encapsulate()
    if err != nil {
        return nil, nil, err
    }
    // Use shared secret as 32-byte AES key
    block, err := aes.NewCipher(ss)
    if err != nil {
        return nil, nil, err
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, nil, err
    }
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, nil, err
    }
    ciphertext := gcm.Seal(nonce, nonce, data, nil)
    return ciphertext, ct, nil
}

// Decrypt decrypts the given ciphertext using the Kyber ciphertext to derive the shared secret.
func (m *PQCManager) Decrypt(ciphertext, ct []byte) ([]byte, error) {
    ss, err := m.priv.Decapsulate(ct)
    if err != nil {
        return nil, err
    }
    block, err := aes.NewCipher(ss)
    if err != nil {
        return nil, err
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, fmt.Errorf("ciphertext too short")
    }
    nonce, data := ciphertext[:nonceSize], ciphertext[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, data, nil)
    if err != nil {
        return nil, err
    }
    return plaintext, nil
}
