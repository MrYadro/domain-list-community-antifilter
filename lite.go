package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func antiLite() {
	// Open input file
	inputFile, err := os.Open("./code/antifilter")
	if err != nil {
		log.Fatalf("Failed to open input file: %v", err)
	}
	defer inputFile.Close()

	// Create output file
	outputFile, err := os.Create("./code/data/antifilter")
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	// Use a map to store unique root domains
	uniqueRoots := make(map[string]bool)

	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		if domain == "" {
			continue
		}

		root, err := publicsuffix.EffectiveTLDPlusOne(domain)
		if err != nil {
			log.Printf("Skipping domain %s: %v", domain, err)
			continue
		}

		// Add to map if not seen before
		if !uniqueRoots[root] {
			uniqueRoots[root] = true
			fmt.Printf("%s → %s\n", domain, root)
		}
	}

	// Write unique roots to output file
	for root := range uniqueRoots {
		_, err := writer.WriteString(root + "\n")
		if err != nil {
			log.Printf("Failed to write to output file: %v", err)
		}
	}

	fmt.Println("Finished! Unique results saved to domains_new.lst")
}
