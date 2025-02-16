package cert9util

// NSS certificate database schema constants
const (
	// Table names
	TableNSSPublic  = "nssPublic"
	TableMetaData   = "metaData"
	TableNSSPrivate = "nssPrivate"

	// Common queries
	CreateCertificateQuery = `INSERT INTO nssPublic (
		nickname, 
		subject, 
		issuer, 
		serial, 
		derCert, 
		trust,
		certType
	) VALUES (?, ?, ?, ?, ?, ?, ?)`
)

// TrustAttributes represents the trust flags for certificates
type TrustAttributes struct {
	ServerAuth bool
	ClientAuth bool
	EmailProtection bool
	CodeSigning bool
}

// EncodeTrust converts trust attributes to NSS trust string format
func (t TrustAttributes) EncodeTrust() string {
	trust := "u" // User cert by default
	flags := ""
	if t.ServerAuth {
		flags += "C"
	}
	if t.ClientAuth {
		flags += "c"
	}
	if t.EmailProtection {
		flags += "E"
	}
	if t.CodeSigning {
		flags += "w"
	}
	if flags != "" {
		trust = "p" + flags + "," + flags + "," + flags // Format: purpose,verify,issue
	}
	return trust
}