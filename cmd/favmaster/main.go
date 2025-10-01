package main

import (
	"favmaster/internal/downloader"
	"favmaster/internal/hasher"
	"favmaster/internal/scanner"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ToolName  = "favmaster"
	Developer = "Cristopher"
	Version   = "1.2.0"
)

func main() {
	if len(os.Args) < 2 {
		showBanner()
		showUsage()
		return
	}

	showBanner()
	input := os.Args[1]

	if isURL(input) || strings.Contains(input, ".") {
		processURLOrDomain(input)
	} else if strings.HasSuffix(strings.ToLower(input), ".txt") {
		processURLFile(input)
	} else {
		processFile(input)
	}
}

func showBanner() {
	fmt.Printf(`
╔═══════════════════════════════════════╗
║              %s              ║
║            v%s by %s           ║
╚═══════════════════════════════════════╝
`, strings.ToUpper(ToolName), Version, Developer)
}

func showUsage() {
	fmt.Println("\nUsage:")
	fmt.Printf("  %s <url/domain>  - Process single URL or domain\n", ToolName)
	fmt.Printf("  %s <file.txt>    - Process URLs from text file\n", ToolName)
	fmt.Printf("  %s <file>        - Process local file\n", ToolName)
	fmt.Println("\nSupported image formats: .ico, .png, .jpg, .jpeg, .gif, .svg, .webp")
	fmt.Println("\nExamples:")
	fmt.Printf("  %s https://example.com/favicon.ico\n", ToolName)
	fmt.Printf("  %s example.com\n", ToolName)
	fmt.Printf("  %s https://example.com/image.png\n", ToolName)
	fmt.Printf("  %s urls.txt\n", ToolName)
	fmt.Printf("  %s image.jpg\n", ToolName)
}

func isURL(input string) bool {
	return strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://")
}

func processURLOrDomain(input string) {
	if isURL(input) {
		if scanner.IsImageURL(input) {
			fmt.Printf("\n🔗 Processing image URL: %s\n", input)
			processSingleURL(input)
		} else {
			fmt.Printf("\n🌐 Processing domain: %s\n", input)
			scanner.FindAndProcessFavicon(input)
		}
	} else {
		fmt.Printf("\n🌐 Processing domain: %s\n", input)
		scanner.FindAndProcessFavicon("https://" + input)
	}
}

func processSingleURL(url string) {
	data, err := downloader.DownloadFile(url)
	if err != nil {
		log.Printf("❌ Error downloading %s: %v\n", url, err)
		return
	}

	if len(data) == 0 {
		log.Printf("❌ Empty response from %s\n", url)
		return
	}

	hasher.CalculateAndDisplayHashes(data, url)
}

func processURLFile(filename string) {
	fmt.Printf("\n📄 Processing URLs from: %s\n", filename)
	urls, err := scanner.ReadURLsFromFile(filename)
	if err != nil {
		log.Printf("❌ Error reading file: %v\n", err)
		return
	}

	processed := 0
	for _, url := range urls {
		processURLOrDomain(url)
		processed++
	}

	fmt.Printf("\n✅ Processed %d/%d entries from %s\n", processed, len(urls), filename)
}

func processFile(filename string) {
	fmt.Printf("\n📁 Processing local file: %s\n", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("❌ Error reading file %s: %v\n", filename, err)
		return
	}

	if len(data) == 0 {
		log.Printf("❌ Empty file: %s\n", filename)
		return
	}

	hasher.CalculateAndDisplayHashes(data, filename)
}
