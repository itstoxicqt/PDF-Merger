package main

import (
	"fmt"
	"log"

	"pdfmerger/cmd" // ✅ FIXED: Correct import path
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("PDF merging completed successfully!")
}
