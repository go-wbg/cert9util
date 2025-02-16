package cert9util

import (
	//"crypto/x509"
	"database/sql"
	//"encoding/base64"
	"fmt"
)

// RemoveCertificate removes a certificate from the cert9 database by its nickname
func (db *CertificateDB9) RemoveCertificate(nickname string) error {
	result, err := db.Exec("DELETE FROM nssPublic WHERE nickname = ?", nickname)
	if err != nil {
		return fmt.Errorf("failed to remove certificate: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("no certificate found with nickname: %s", nickname)
	}

	return nil
}

// RemoveCertificateBySubject removes all certificates matching a specific subject
func (db *CertificateDB9) RemoveCertificateBySubject(subject string) error {
	result, err := db.Exec("DELETE FROM nssPublic WHERE subject = ?", subject)
	if err != nil {
		return fmt.Errorf("failed to remove certificates: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("no certificates found with subject: %s", subject)
	}

	return nil
}

// RemoveCertificateBySerialAndIssuer removes a certificate by its serial number and issuer
func (db *CertificateDB9) RemoveCertificateBySerialAndIssuer(serial, issuer string) error {
	result, err := db.Exec("DELETE FROM nssPublic WHERE serial = ? AND issuer = ?", serial, issuer)
	if err != nil {
		return fmt.Errorf("failed to remove certificate: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("no certificate found with serial %s and issuer %s", serial, issuer)
	}

	return nil
}

// ListCertificates returns a list of all certificates in the database
func (db *CertificateDB9) ListCertificates() ([]CertificateInfo, error) {
	rows, err := db.Query("SELECT nickname, subject, issuer, serial FROM nssPublic")
	if err != nil {
		return nil, fmt.Errorf("failed to query certificates: %v", err)
	}
	defer rows.Close()

	var certificates []CertificateInfo
	for rows.Next() {
		var cert CertificateInfo
		err := rows.Scan(&cert.Nickname, &cert.Subject, &cert.Issuer, &cert.Serial)
		if err != nil {
			return nil, fmt.Errorf("failed to scan certificate row: %v", err)
		}
		certificates = append(certificates, cert)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating certificate rows: %v", err)
	}

	return certificates, nil
}

// CertificateInfo represents the basic information about a certificate in the database
type CertificateInfo struct {
	Nickname string
	Subject  string
	Issuer   string
	Serial   string
}

// GetCertificateByNickname retrieves certificate information by its nickname
func (db *CertificateDB9) GetCertificateByNickname(nickname string) (*CertificateInfo, error) {
	var cert CertificateInfo
	err := db.QueryRow("SELECT nickname, subject, issuer, serial FROM nssPublic WHERE nickname = ?", 
		nickname).Scan(&cert.Nickname, &cert.Subject, &cert.Issuer, &cert.Serial)
	
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("no certificate found with nickname: %s", nickname)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query certificate: %v", err)
	}

	return &cert, nil
}