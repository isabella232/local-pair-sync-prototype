package app

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"testing"
)

func TestGenerateKeyAndCert(t *testing.T) {
	crtWriter := bytes.NewBuffer([]byte{})
	keyWriter := bytes.NewBuffer([]byte{})

	err := GenerateKeyAndCert(crtWriter, keyWriter)
	if err != nil {
		t.Error(err)
	}

	expected := []*struct {
		Name string
		Bytes []byte
		Len int
		BeginWith string
		EndWith string
		PemBlock *pem.Block
	}{
		{
			"certificate",
			crtWriter.Bytes(),
			509,
			"-----BEGIN CERTIFICATE-----",
			"-----END CERTIFICATE-----",
			nil,
		},
		{
			"private key",
			keyWriter.Bytes(),
			227,
			"-----BEGIN EC PRIVATE KEY-----",
			"-----END EC PRIVATE KEY-----",
			nil,
		},
	}

	for _, e := range expected {
		if len(e.Bytes) != e.Len {
			t.Errorf("%s pem bytes should be %d in length, received %d",
				e.Name,
				e.Len,
				len(e.Bytes),
			)
		}

		bwl := len(e.BeginWith)
		if string(e.Bytes[:bwl]) != e.BeginWith {
			t.Errorf("%s pem should begin with '%s', recieved '%s'",
				e.Name,
				e.BeginWith,
				string(e.Bytes[:bwl]),
			)
		}

		ewl := len(e.EndWith)
		if string(e.Bytes[len(e.Bytes)-ewl-1:len(e.Bytes)-1]) != e.EndWith{
			t.Errorf("%s pem should end with '%s', recieved '%s'",
				e.Name,
				e.EndWith,
				e.Bytes[len(e.Bytes)-ewl-1:len(e.Bytes)-1],
			)
		}

		pb, r := pem.Decode(e.Bytes)
		if len(r) != 0 {
			t.Errorf("%s pem decode should be empty, remaining bytes %v", e.Name, r)
		}
		e.PemBlock = pb
	}

	_, err = x509.ParseCertificate(expected[0].PemBlock.Bytes)
	if err != nil {
		t.Error(err)
	}

	_, err = x509.ParseECPrivateKey(expected[1].PemBlock.Bytes)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteKeyAndCert(t *testing.T) {
	err := DeleteKeyAndCert()
	if err != nil {
		t.Error(err)
	}
}