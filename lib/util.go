package cert9util

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// ImportCertificateFromFile imports a certificate from a PEM file
func (db *CertificateDB9) ImportCertificateFromFile(pemFile, nickname string) error {
	// Read the certificate file
	certPEM, err := ioutil.ReadFile(pemFile)
	if err != nil {
		return fmt.Errorf("failed to read certificate file: %v", err)
	}

	// Decode the PEM block
	block, _ := pem.Decode(certPEM)
	if block == nil {
		return fmt.Errorf("failed to decode PEM block")
	}

	// Parse the certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse certificate: %v", err)
	}

	// Add the certificate to the database
	return db.AddSelfSignedCertificate(cert, nickname)
}