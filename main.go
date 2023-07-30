package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX,hasSPF,spfRecord,hasDmarc,dmarcRecord")
	for scanner.Scan() {
		checkdomain(scanner.Text())
	}
	if scanner.Err()!=nil{
		Err(scanner.Err())
	}
}

func checkdomain(domain string) {
   

}

func Err(err error) error {

	return fmt.Errorf("Error Occured", err)
}
