package main

import (
	"flag"
	"fmt"
	"log"
	//"os"
	"path/filepath"

	cert9util "github.com/eyedeekay/cert9util/lib"
)

func main() {
	// Command-line flags
	certFile := flag.String("cert", "", "Path to the certificate file (PEM format)")
	nickname := flag.String("nickname", "", "Nickname for the certificate")
	dbPath := flag.String("db", "cert9.db", "Path to cert9.db file")
	remove := flag.Bool("remove", false, "Remove certificate instead of adding")
	list := flag.Bool("list", false, "List all certificates in the database")
	subject := flag.String("subject", "", "Remove certificates by subject")
	flag.Parse()

	// Open the certificate database
	db, err := cert9util.NewCertificateDB9(*dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Handle list operation
	if *list {
		certs, err := db.ListCertificates()
		if err != nil {
			log.Fatalf("Failed to list certificates: %v", err)
		}
		fmt.Println("Certificates in database:")
		for _, cert := range certs {
			fmt.Printf("Nickname: %s\n  Subject: %s\n  Issuer: %s\n  Serial: %s\n\n",
				cert.Nickname, cert.Subject, cert.Issuer, cert.Serial)
		}
		return
	}

	// Handle remove operation
	if *remove {
		if *nickname != "" {
			if err := db.RemoveCertificate(*nickname); err != nil {
				log.Fatalf("Failed to remove certificate: %v", err)
			}
			log.Printf("Successfully removed certificate with nickname: %s", *nickname)
		} else if *subject != "" {
			if err := db.RemoveCertificateBySubject(*subject); err != nil {
				log.Fatalf("Failed to remove certificates: %v", err)
			}
			log.Printf("Successfully removed certificates with subject: %s", *subject)
		} else {
			log.Fatal("Must specify either -nickname or -subject when removing certificates")
		}
		return
	}

	// Handle add operation
	if *certFile == "" || *nickname == "" {
		log.Fatal("Certificate file and nickname are required for adding certificates")
	}

	if err := db.ImportCertificateFromFile(*certFile, *nickname); err != nil {
		log.Fatalf("Failed to import certificate: %v", err)
	}

	log.Printf("Successfully imported certificate '%s' from %s", *nickname, filepath.Base(*certFile))
}