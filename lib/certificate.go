package cert9util

import (
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

// AddCertificate adds a new certificate to the cert9 database
func (db *CertificateDB9) AddCertificate(cert *x509.Certificate, nickname string, trust TrustAttributes) error {
	// Convert certificate to DER format
	derCert := cert.Raw

	// Extract subject and issuer
	subject := cert.Subject.String()
	issuer := cert.Issuer.String()
	
	// Convert serial to string
	serial := base64.StdEncoding.EncodeToString(cert.SerialNumber.Bytes())

	// Set certificate type (1 for user certificates)
	certType := uint(1)

	// Encode trust attributes
	trustFlags := trust.EncodeTrust()

	// Insert the certificate
	_, err := db.Exec(CreateCertificateQuery,
		nickname,
		subject,
		issuer,
		serial,
		derCert,
		trustFlags,
		certType,
	)
	if err != nil {
		return fmt.Errorf("failed to add certificate: %v", err)
	}

	return nil
}

// AddSelfSignedCertificate adds a self-signed certificate to the database
func (db *CertificateDB9) AddSelfSignedCertificate(cert *x509.Certificate, nickname string) error {
	// Default trust for self-signed certificates
	trust := TrustAttributes{
		ServerAuth: true,
		ClientAuth: true,
		EmailProtection: true,
		CodeSigning: false,
	}

	return db.AddCertificate(cert, nickname, trust)
}