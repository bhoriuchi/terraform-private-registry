package server

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)

const rsaBits = 2048

var validFor = 365 * 24 * time.Hour

func generateSelfSigned() (err error) {
	var privateKey *rsa.PrivateKey
	var serialNumber *big.Int
	var certOut *os.File
	var derBytes []byte
	var privBytes []byte

	privateKey, err = rsa.GenerateKey(rand.Reader, rsaBits)
	keyUsage := x509.KeyUsageDigitalSignature
	keyUsage |= x509.KeyUsageKeyEncipherment
	notBefore := time.Now()
	notAfter := notBefore.Add(validFor)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err = rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		err = fmt.Errorf("Failed to generate serial number: %v", err)
		return
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames: []string{
			"localhost",
		},
	}

	derBytes, err = x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Fatalf("Failed to create certificate: %v", err)
	}

	if certOut, err = os.Create("cert.pem"); err != nil {
		return
	}

	if err = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		err = fmt.Errorf("Failed to write data to cert.pem: %v", err)
		return
	}

	if err = certOut.Close(); err != nil {
		err = fmt.Errorf("Error closing cert.pem: %v", err)
		return
	}

	keyOut, err := os.OpenFile("key.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		err = fmt.Errorf("Failed to open key.pem for writing: %v", err)
		return
	}
	privBytes, err = x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		log.Fatalf("Unable to marshal private key: %v", err)
		return
	}
	if err = pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		err = fmt.Errorf("Failed to write data to key.pem: %v", err)
		return
	}
	if err = keyOut.Close(); err != nil {
		err = fmt.Errorf("Error closing key.pem: %v", err)
		return
	}

	return
}
