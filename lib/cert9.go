package cert9util

type CertificateDB9 struct {
	SQLiteDB
}

func NewCertificateDB9(path string) (*CertificateDB9, error) {
	db, err := NewSQLiteDB(path)
	if err != nil {
		return nil, err
	}
	return &CertificateDB9{*db}, nil
}

func (db *CertificateDB9) Columns() ([]string, error) {
	return db.SQLiteDB.Columns("nssPublic")
}
