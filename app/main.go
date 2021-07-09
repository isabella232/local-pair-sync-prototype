package app

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
)

const (
	KeysDir = "./keys/"
	TLSCertName = KeysDir + "tls.crt"
	TLSKeyName = KeysDir + "tls.key"
)

type App struct {
	CertPemBytes []byte
	KeyPemBytes []byte
	TLSConfig tls.Config
}

func (a *App) Init() error {
	return a.generateKeyAndCert()
}

func (a *App) generateKeyAndCert() error {
	crtWriter := bytes.NewBuffer([]byte{})
	keyWriter := bytes.NewBuffer([]byte{})

	err := GenerateKeyAndCert(crtWriter, keyWriter)
	if err != nil {
		return err
	}

	a.KeyPemBytes = keyWriter.Bytes()
	a.CertPemBytes = crtWriter.Bytes()
	return nil
}

func (a *App) MakeTLS() error {
	cert, err := tls.X509KeyPair(a.CertPemBytes, a.KeyPemBytes)
	if err != nil {
		return err
	}

	a.TLSConfig = tls.Config{Certificates: []tls.Certificate{cert}, Rand: rand.Reader}

	return nil
}