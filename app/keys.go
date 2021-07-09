package app

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io"
	"math/big"
	"os"
	"time"
)

const (
	OrganisationName = "Status IM"
)

func getCertTemplate() *x509.Certificate {
	return &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{OrganisationName},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
}

func generateCert(certWriter io.Writer, key *ecdsa.PrivateKey) error {
	certTemplate := getCertTemplate()
	certBytes, err := x509.CreateCertificate(rand.Reader, certTemplate, certTemplate, &key.PublicKey, key)
	if err != nil {
		return err
	}

	return pem.Encode(certWriter, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
}

func generateKey(keyWriter io.Writer, key *ecdsa.PrivateKey) error {
	b, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}

	return pem.Encode(keyWriter, &pem.Block{Type: "EC PRIVATE KEY", Bytes: b})
}

func GenerateKeyAndCert(certWriter io.Writer, keyWriter io.Writer) error {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		 return err
	}

	err = generateCert(certWriter, privKey)
	if err != nil {
		return err
	}

	err = generateKey(keyWriter, privKey)
	if err != nil {
		return err
	}

	return nil
}

func DeleteKeyAndCert() error {
	err := os.Remove(TLSKeyName)
	if err != nil {
		return err
	}

	return os.Remove(TLSCertName)
}