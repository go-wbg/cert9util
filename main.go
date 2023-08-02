package main

import (
	"log"

	cert9util "github.com/eyedeekay/cert9util/lib"
)

func main() {
	cdbtest, err := cert9util.NewCertificateDB9("cert9.db")
	if err != nil {
		panic(err)
	}
	defer cdbtest.Close()
	ccols, err := cdbtest.Columns()
	if err != nil {
		panic(err)
	}
	for _, col := range ccols {
		log.Println("cert column:" + col)
	}

	kdbtest, err := cert9util.NewKeyDB9("key4.db")
	if err != nil {
		panic(err)
	}
	defer kdbtest.Close()
	kcols, err := kdbtest.Columns()
	if err != nil {
		panic(err)
	}
	for _, col := range kcols {
		if str, ok := contains(ccols, col); ok {
			log.Println("key column:"+col, "appears in cert:"+str)
		} else {
			log.Println("key column:" + col, "does not appear in cert")
		}
	}
}

func contains(s []string, e string) (string, bool) {
	for _, a := range s {
		if a == e {
			return a, true
		}
	}
	return "", false

}
