package main

import (
    "crypto/md5"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/PuerkitoBio/goquery"
    "github.com/spaolacci/murmur3"
)

const toolVersion = "v1.0.0"

func computeHashes(data []byte) (md5sum, sha256sum, mmh3sum string) {
    md5hash := md5.Sum(data)
    md5sum = hex.EncodeToString(md5hash[:])

    sha := sha256.Sum256(data)
    sha256sum = hex.EncodeToString(sha[:])

    u := murmur3.Sum32(data)
    mmh3sum = fmt.Sprintf("%d", u)

    return
}

func fetchURL(rawurl string) ([]byte, error) {
    client := &http.Client{Timeout: 30 * time.Second}
    resp, err := client.Get(rawurl)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("HTTP status: %s", resp.Status)
    }

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func discoverFaviconURL(base *url.URL) (string, error) {
    resp, err := http.Get(base.String())
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return "", err
    }

    var href string
    doc.Find("link").EachWithBreak(func(i int, s *goquery.Selection) bool {
        rel, _ := s.Attr("rel")
        relLower := strings.ToLower(rel)
        if strings.Contains(relLower, "icon") {
            h, ok := s.Attr("href")
            if ok {
                href = h
                return false
            }
        }
        return true
    })

    if href == "" {
        return base.Scheme + "://" + base.Host + "/favicon.ico", nil
    }

    u, err := url.Parse(href)
    if err != nil {
        return "", err
    }
    full := base.ResolveReference(u)
    return full.String(), nil
}

func processURLOrDomain(input string) error {
    if !strings.HasPrefix(input, "http://") && !strings.HasPrefix(input, "https://") {
        input = "https://" + input
    }
    parsed, err := url.Parse(input)
    if err != nil {
        return fmt.Errorf("invalid URL: %w", err)
    }

    faviconURL, err := discoverFaviconURL(parsed)
    if err != nil {
        return fmt.Errorf("favicon discovery failed: %w", err)
    }

    data, err := fetchURL(faviconURL)
    if err != nil {
        return fmt.Errorf("fetch favicon failed: %w", err)
    }

    md5sum, sha256sum, mmh3sum := computeHashes(data)

    fmt.Printf("Source: %s\n", faviconURL)
    fmt.Printf("Size: %d bytes\n", len(data))
    fmt.Printf("MMH3: %s\n", mmh3sum)
    fmt.Printf("SHA256: %s\n", sha256sum)
    fmt.Printf("MD5: %s\n", md5sum)
    return nil
}

func processLocalFile(path string) error {
    data, err := os.ReadFile(path)
    if err != nil {
        return fmt.Errorf("read file failed: %w", err)
    }
    md5sum, sha256sum, mmh3sum := computeHashes(data)

    fmt.Printf("File: %s\n", path)
    fmt.Printf("Size: %d bytes\n", len(data))
    fmt.Printf("MMH3: %s\n", mmh3sum)
    fmt.Printf("SHA256: %s\n", sha256sum)
    fmt.Printf("MD5: %s\n", md5sum)
    return nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Printf("Usage: favmaster <url|domain|file>\n")
        os.Exit(1)
    }

    input := os.Args[1]
    if fi, err := os.Stat(input); err == nil && !fi.IsDir() {
        if err := processLocalFile(input); err != nil {
            fmt.Fprintf(os.Stderr, "Error: %v\n", err)
            os.Exit(1)
        }
    } else {
        if err := processURLOrDomain(input); err != nil {
            fmt.Fprintf(os.Stderr, "Error: %v\n", err)
            os.Exit(1)
        }
    }
}
