package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX,hasSPF,spfRecord,hasDmarc,dmarcRecord")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Err could not read from input", err)
	}
}

func checkDomain(domain string) {
	var hasMx, hasSpf, hasDmarc bool
	var spfRecord, dmarcRecord string
	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain) //Spf Records

	if err != nil {
		log.Printf("Error %v\n", err)

	}

	for _, record := range txtRecords {

		if strings.HasPrefix(record, "v=spf1") {
			hasSpf = true
			spfRecord = record
			break
		}

	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error  %v \n", err)

	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDmarc = true
			dmarcRecord = record
			break
		}
	}
}
