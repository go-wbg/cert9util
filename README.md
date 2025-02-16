# cert9util

[![Go Report Card](https://goreportcard.com/badge/github.com/eyedeekay/cert9util)](https://goreportcard.com/report/github.com/eyedeekay/cert9util)
[![GoDoc](https://godoc.org/github.com/eyedeekay/cert9util?status.svg)](https://godoc.org/github.com/eyedeekay/cert9util)
[![License](https://img.shields.io/github/license/eyedeekay/cert9util)](LICENSE)

A Go library and command-line tool for managing Mozilla Firefox's cert9.db and key4.db certificate databases. This tool allows you to add, remove, and list certificates in Firefox's NSS database format.

## 🚀 Quick Start

```bash
# Install the tool
go install github.com/eyedeekay/cert9util@latest

# Add a certificate
cert9util -cert mycert.pem -nickname "My Certificate" -db ~/.mozilla/firefox/XXXXXX.default/cert9.db

# List all certificates
cert9util -list -db ~/.mozilla/firefox/XXXXXX.default/cert9.db

# Remove a certificate
cert9util -remove -nickname "My Certificate" -db ~/.mozilla/firefox/XXXXXX.default/cert9.db
```

## ✨ Features

- Add self-signed certificates to Firefox's cert9.db
- Remove certificates by nickname or subject
- List all certificates in the database
- Manage certificate trust attributes
- Support for both cert9.db and key4.db databases
- Pure Go implementation using modernc.org/sqlite

## 📋 Requirements

- Go 1.16 or later
- Firefox profile directory with cert9.db
- Write permissions to the Firefox profile directory

## 📦 Installation

### From Source
```bash
git clone https://github.com/eyedeekay/cert9util.git
cd cert9util
go build
```

### Via Go Install
```bash
go install github.com/eyedeekay/cert9util@latest
```

## 🔧 Configuration

No configuration file is needed. The tool operates directly on Firefox's certificate databases.

Common Firefox profile locations:
- Linux: `~/.mozilla/firefox/<profile>/`
- macOS: `~/Library/Application Support/Firefox/Profiles/<profile>/`
- Windows: `%APPDATA%\Mozilla\Firefox\Profiles\<profile>\`

## 🎯 Usage

### Command Line Interface

```bash
cert9util [options]

Options:
  -cert string
        Path to the certificate file (PEM format)
  -db string
        Path to cert9.db file (default "cert9.db")
  -nickname string
        Nickname for the certificate
  -remove
        Remove certificate instead of adding
  -list
        List all certificates in the database
  -subject string
        Remove certificates by subject
```

### Library Usage

```go
package main

import (
    "log"
    cert9util "github.com/eyedeekay/cert9util/lib"
)

func main() {
    // Open the database
    db, err := cert9util.NewCertificateDB9("cert9.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // List certificates
    certs, err := db.ListCertificates()
    if err != nil {
        log.Fatal(err)
    }
    
    for _, cert := range certs {
        log.Printf("Certificate: %s", cert.Nickname)
    }
}
```

## 🛠️ Development

### Project Structure
- `lib/`: Core library code
  - `cert9.go`: Certificate database operations
  - `key4.go`: Key database operations
  - `schema.go`: Database schema definitions
  - `certificate.go`: Certificate manipulation
  - `rmcertificate.go`: Certificate removal operations
  - `util.go`: Utility functions
  - `sql.go`: SQLite database operations

### Testing

```bash
go test ./...
```

## ✅ Common Operations

### Adding a Self-Signed Certificate
```bash
cert9util -cert self-signed.pem -nickname "My Self-Signed Cert" -db cert9.db
```

### Removing Certificates
```bash
# Remove by nickname
cert9util -remove -nickname "My Self-Signed Cert" -db cert9.db

# Remove by subject
cert9util -remove -subject "CN=example.com" -db cert9.db
```

### Listing Certificates
```bash
cert9util -list -db cert9.db
```

## 🔍 Troubleshooting

### Common Issues

1. **Database Locked**
   - Ensure Firefox is not running when modifying the database
   - Check file permissions

2. **Certificate Import Fails**
   - Verify the certificate is in PEM format
   - Check if the certificate is valid
   - Ensure the database path is correct

3. **Permission Denied**
   - Check file ownership and permissions
   - Run with appropriate privileges

## 👥 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## ⚠️ Security Considerations

- Always backup your cert9.db before modifications
- Keep your Firefox profile secure
- Be cautious when adding self-signed certificates
- Verify certificate contents before import

## 🙏 Acknowledgments

- Mozilla NSS team for database format documentation
- modernc.org/sqlite for pure Go SQLite implementation