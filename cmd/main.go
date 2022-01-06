package main

import (
	"fmt"
	"log"
	"sandbox/bios"
)

func main() {
	bi, err := bios.BIOS()
	if err != nil {
		log.Fatalf("ERROR: cannot get bios info: %v", err)
	}
	fmt.Printf("vendor: %s\nversion: %s\ndate: %s\n", bi.Vendor, bi.Version, bi.Date)
}
