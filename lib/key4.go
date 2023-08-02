package cert9util

type KeyDB9 struct {
	SQLiteDB
}

func NewKeyDB9(path string) (*KeyDB9, error) {
	db, err := NewSQLiteDB(path)
	if err != nil {
		return nil, err
	}
	return &KeyDB9{*db}, nil
}

func (db *KeyDB9) Columns() ([]string, error) {
	md, err := db.MetaDataColumns()
	if err != nil {
		return nil, err
	}
	nssPrivate, err := db.NSSPrivateColumns()
	if err != nil {
		return nil, err
	}
	return append(md, nssPrivate...), nil
}

func (db *KeyDB9) MetaDataColumns() ([]string, error) {
	md, err := db.SQLiteDB.Columns("metaData")
	if err != nil {
		return nil, err
	}
	return md, nil
}

func (db *KeyDB9) NSSPrivateColumns() ([]string, error) {
	nssPrivate, err := db.SQLiteDB.Columns("nssPrivate")
	if err != nil {
		return nil, err
	}
	return nssPrivate, nil
}
