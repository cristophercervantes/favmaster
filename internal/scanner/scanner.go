package scanner

import (
	"bufio"
	"favmaster/internal/downloader"
	"favmaster/internal/hasher"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	SupportedImageExtensions = map[string]bool{
		".ico": true, ".png": true, ".jpg": true, ".jpeg": true,
		".gif": true, ".svg": true, ".webp": true,
	}

	FaviconPaths = []string{
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

func IsImageURL(url string) bool {
	lowerURL := strings.ToLower(url)
	for ext := range SupportedImageExtensions {
		if strings.HasSuffix(lowerURL, ext) {
			return true
		}
	}
	return false
}

func FindAndProcessFavicon(domain string) {
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
	for _, path := range FaviconPaths {
		faviconURL := normalizeURL(domain) + path
		fmt.Printf("   Trying: %s\n", faviconURL)

		if checkAndProcessImageURL(faviconURL) {
			return true
		}
	}
	return false
}

func parseHTMLForFavicon(domain string) bool {
	data, err := downloader.DownloadFile(domain)
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
	if !downloader.CheckURLExists(url) {
		return false
	}

	contentType := downloader.GetContentType(url)
	if strings.HasPrefix(contentType, "image/") {
		processImageURL(url)
		return true
	}

	return false
}

func processImageURL(url string) {
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

func ReadURLsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			urls = append(urls, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
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
