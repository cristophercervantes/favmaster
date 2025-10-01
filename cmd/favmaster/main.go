package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unsafe"

	"github.com/PuerkitoBio/goquery"
)

const (
	ToolName    = "favmaster"
	Developer   = "Cristopher"
	Version     = "1.2.2"
)

var (
	supportedImageExtensions = map[string]bool{
		".ico": true, ".png": true, ".jpg": true, ".jpeg": true,
		".gif": true, ".svg": true, ".webp": true,
	}
	
	faviconPaths = []string{
		"/favicon.ico",
		"/favicon.png",
		"/favicon.jpg",
		"/favicon.jpeg",
		"/favicon.svg",
		"/favicon.webp",
		"/images/favicon.ico",
		"/img/favicon.ico",
		"/assets/favicon.ico",
		"/static/favicon.ico",
		"/icon.png",
		"/logo.png",
		"/apple-touch-icon.png",
	}
)

// MurmurHash3 implementation (32-bit version)
func murmurHash3(data []byte, seed uint32) uint32 {
	const (
		c1 uint32 = 0xcc9e2d51
		c2 uint32 = 0x1b873593
		r1 uint32 = 15
		r2 uint32 = 13
		m  uint32 = 5
		n  uint32 = 0xe6546b64
	)

	h := seed
	length := len(data)
	roundedEnd := (length / 4) * 4

	for i := 0; i < roundedEnd; i += 4 {
		k := *(*uint32)(unsafe.Pointer(&data[i]))
		k *= c1
		k = (k << r1) | (k >> (32 - r1))
		k *= c2

		h ^= k
		h = (h << r2) | (h >> (32 - r2))
		h = h*m + n
	}

	var k uint32
	switch length - roundedEnd {
	case 3:
		k ^= uint32(data[roundedEnd+2]) << 16
		fallthrough
	case 2:
		k ^= uint32(data[roundedEnd+1]) << 8
		fallthrough
	case 1:
		k ^= uint32(data[roundedEnd])
		k *= c1
		k = (k << r1) | (k >> (32 - r1))
		k *= c2
		h ^= k
	}

	h ^= uint32(length)
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16

	return h
}

// Safe MurmurHash3 that converts to int32 (for Shodan compatibility)
func mmh3Hash32(data []byte) int32 {
	return int32(murmurHash3(data, 0))
}

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
		if isImageURL(input) {
			fmt.Printf("\n🔗 Processing image URL: %s\n", input)
			processSingleURL(input)
		} else {
			fmt.Printf("\n🌐 Processing domain: %s\n", input)
			findAndProcessFavicon(input)
		}
	} else {
		fmt.Printf("\n🌐 Processing domain: %s\n", input)
		findAndProcessFavicon("https://" + input)
	}
}

func isImageURL(url string) bool {
	lowerURL := strings.ToLower(url)
	for ext := range supportedImageExtensions {
		if strings.HasSuffix(lowerURL, ext) {
			return true
		}
	}
	return false
}

func findAndProcessFavicon(domain string) {
	fmt.Printf("   Searching for favicon...\n")
	
	// Try common favicon paths first
	found := tryCommonFaviconPaths(domain)
	if !found {
		// If not found, parse HTML to find favicon
		found = parseHTMLForFavicon(domain)
	}
	
	if !found {
		log.Printf("❌ No favicon found for %s\n", domain)
	}
}

func tryCommonFaviconPaths(domain string) bool {
	for _, path := range faviconPaths {
		faviconURL := normalizeURL(domain) + path
		fmt.Printf("   Trying: %s\n", faviconURL)
		
		if checkAndProcessImageURL(faviconURL) {
			return true
		}
	}
	return false
}

func parseHTMLForFavicon(domain string) bool {
	data, err := downloadFile(domain)
	if err != nil {
		log.Printf("   Error fetching page: %v\n", err)
		return false
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		return false
	}

	// Look for favicon in link tags
	found := false
	doc.Find("link[rel*='icon'], link[rel*='apple-touch-icon']").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && href != "" {
			faviconURL := resolveURL(domain, href)
			fmt.Printf("   Found in HTML: %s\n", faviconURL)
			if checkAndProcessImageURL(faviconURL) {
				found = true
			}
		}
	})

	return found
}

func checkAndProcessImageURL(url string) bool {
	if !checkURLExists(url) {
		return false
	}

	contentType := getContentType(url)
	if strings.HasPrefix(contentType, "image/") {
		processSingleURL(url)
		return true
	}

	return false
}

func processSingleURL(url string) {
	data, err := downloadFile(url)
	if err != nil {
		log.Printf("❌ Error downloading %s: %v\n", url, err)
		return
	}

	if len(data) == 0 {
		log.Printf("❌ Empty response from %s\n", url)
		return
	}

	calculateAndDisplayHashes(data, url)
}

func processURLFile(filename string) {
	fmt.Printf("\n📄 Processing URLs from: %s\n", filename)
	
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("❌ Error opening file %s: %v\n", filename, err)
		return
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	urlCount := 0
	processedCount := 0
	
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			processURLOrDomain(line)
			urlCount++
			processedCount++
			time.Sleep(100 * time.Millisecond)
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Printf("❌ Error reading file: %v\n", err)
	}
	
	fmt.Printf("\n✅ Processed %d/%d entries from %s\n", processedCount, urlCount, filename)
}

func processFile(filename string) {
	fmt.Printf("\n📁 Processing local file: %s\n", filename)
	
	// Check if it's a supported image format
	ext := strings.ToLower(filepath.Ext(filename))
	if !supportedImageExtensions[ext] {
		fmt.Printf("⚠️  Note: %s may not be an image file\n", filename)
	}
	
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("❌ Error reading file %s: %v\n", filename, err)
		return
	}
	
	if len(data) == 0 {
		log.Printf("❌ Empty file: %s\n", filename)
		return
	}
	
	calculateAndDisplayHashes(data, filename)
}

func downloadFile(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}
	
	return io.ReadAll(resp.Body)
}

func checkURLExists(url string) bool {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	resp, err := client.Head(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	
	return resp.StatusCode == http.StatusOK
}

func getContentType(url string) string {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	resp, err := client.Head(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	
	return resp.Header.Get("Content-Type")
}

func calculateAndDisplayHashes(data []byte, source string) {
	// Calculate MurmurHash3 (mmh3) - using our own implementation
	mmh3Hash := mmh3Hash32(data)
	
	// Calculate SHA256
	sha256Hash := sha256.Sum256(data)
	
	// Calculate MD5
	md5Hash := md5.Sum(data)
	
	// Display results
	fmt.Println("┌─────────────────────────────────────────┐")
	fmt.Printf("│ Source: %-30s │\n", truncateString(source, 30))
	fmt.Printf("│ Size:   %-30d │\n", len(data))
	fmt.Println("├─────────────────────────────────────────┤")
	fmt.Printf("│ MMH3:   %-32d │\n", mmh3Hash)
	fmt.Printf("│ SHA256: %-32s │\n", truncateString(hex.EncodeToString(sha256Hash[:]), 32))
	fmt.Printf("│ MD5:    %-32s │\n", truncateString(hex.EncodeToString(md5Hash[:]), 32))
	fmt.Println("└─────────────────────────────────────────┘")
}

func truncateString(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength-3] + "..."
}

func normalizeURL(domain string) string {
	if strings.HasPrefix(domain, "http://") || strings.HasPrefix(domain, "https://") {
		return domain
	}
	return "https://" + domain
}

func resolveURL(base, relative string) string {
	baseURL, err := url.Parse(base)
	if err != nil {
		return relative
	}
	
	relativeURL, err := url.Parse(relative)
	if err != nil {
		return relative
	}
	
	return baseURL.ResolveReference(relativeURL).String()
}
