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

	"github.com/PuerkitoBio/goquery"
	"github.com/tj/go-murmur"
)

const (
	ToolName    = "favmaster"
	Developer   = "Cristopher"
	Version     = "1.2.0"
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

// [ALL OTHER FUNCTIONS REMAIN EXACTLY THE SAME AS BEFORE]
// ... (include all the same functions from your previous main.go)

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

// [Include all other functions exactly as they were in your main.go]
