package main

import (
	"flag"
	"log"
	"path/filepath"

	cert9util "github.com/eyedeekay/cert9util/lib"
)

func main() {
	certFile := flag.String("cert", "", "Path to the certificate file (PEM format)")
	nickname := flag.String("nickname", "", "Nickname for the certificate")
	dbPath := flag.String("db", "cert9.db", "Path to cert9.db file")
	flag.Parse()

	if *certFile == "" || *nickname == "" {
		log.Fatal("Certificate file and nickname are required")
	}

	// Open the certificate database
	db, err := cert9util.NewCertificateDB9(*dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Import the certificate
	err = db.ImportCertificateFromFile(*certFile, *nickname)
	if err != nil {
		log.Fatalf("Failed to import certificate: %v", err)
	}

	log.Printf("Successfully imported certificate '%s' from %s", *nickname, filepath.Base(*certFile))
}