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

	//Take input from terminal and call the checkDomain to check various parameters
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Domain. hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error: could not read from input: %v\n", err)
	}

}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	//Check for hasMX using LookupMX func of net package

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	//Check for hasSPF

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error : %v", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v= spf1") {
			hasSPF = true
			spfRecord = record
			break
		}

	}

	// Check for hasDMARC

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error %v \n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	//Check all the Returned parameters
	fmt.Printf("%v, %v, %v, %v, %v ,%v\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
	fmt.Printf("Exiting the program...\n")
	os.Exit(1)

}
