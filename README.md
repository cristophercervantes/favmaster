# FavMaster 🔍

A powerful and versatile tool for extracting MMH3, SHA256, and MD5 hashes from images, favicons, and URLs. Perfect for security researchers, penetration testers, and developers.

[![Version](https://img.shields.io/badge/Version-1.2.0-blue)](https://github.com/cristophercervantes/favmaster)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Linux%2520%257C%2520Windows%2520%257C%2520macOS-lightgrey)]()

---

## ✨ Features

* 🔍 **Automatic Favicon Discovery** — Intelligently finds favicons from domains
* 🖼️ **Multi-Format Support** — PNG, JPG, JPEG, GIF, SVG, WEBP, ICO
* 🌐 **Flexible Input Sources** — URLs, domains, text files, or local files
* 📊 **Multiple Hash Algorithms** — MMH3, SHA256, and MD5 calculations
* 🚀 **High Performance** — Fast, concurrent processing with robust error handling
* 🔧 **Easy Installation** — Simple `go install` command
* 🎯 **Smart URL Processing** — Auto-detects domains and finds favicons automatically

---

## 🚀 Installation

### Quick Install (Recommended)

```bash
go install github.com/cristophercervantes/favmaster@latest
```

Make sure `$GOPATH/bin` is in your `PATH` environment variable.

### From Source

```bash
git clone https://github.com/cristophercervantes/favmaster
cd favmaster
go build -o favmaster cmd/favmaster/main.go
sudo mv favmaster /usr/local/bin/  # Optional: move to PATH
```

### Using Make (Development)

```bash
git clone https://github.com/cristophercervantes/favmaster
cd favmaster
make build        # Build binary
make install-local # Install from local source
```

---

## 📖 Usage

### Basic Examples

```bash
# Process a domain (auto-find favicon)
favmaster google.com

# Process specific image URL
favmaster https://example.com/logo.png

# Process from text file with URLs
favmaster urls.txt

# Process local image file
favmaster image.jpg

# Direct favicon URL
favmaster https://github.com/favicon.ico
```

### Advanced Examples

```bash
# Domain without protocol (auto-adds https://)
favmaster example.com

# HTTP URL
favmaster http://example.com/favicon.ico

# Mixed input file with comments
favmaster targets.txt

# Process current directory image
favmaster ./photo.png
```

### Text File Input Format

Create a text file with one URL or domain per line (comments supported with `#`):

```
# Popular websites
google.com
https://github.com
http://example.com

# Specific image URLs
https://example.com/logo.png
https://site.com/images/favicon.ico

# Mixed entries
stackoverflow.com
assets.example.com/favicon.jpg
```

---

## 🎯 Supported Image Formats

| Format  | Extension   | Description                      |
| ------- | ----------- | -------------------------------- |
| Favicon | .ico        | Windows icon format              |
| PNG     | .png        | Portable Network Graphics        |
| JPEG    | .jpg, .jpeg | Joint Photographic Experts Group |
| GIF     | .gif        | Graphics Interchange Format      |
| SVG     | .svg        | Scalable Vector Graphics         |
| WebP    | .webp       | Modern web image format          |

---

## 📊 Output Example

```
╔═══════════════════════════════════════╗
║              FAVMASTER              ║
║            v1.2.0 by Cristopher           ║
╚═══════════════════════════════════════╝

🌐 Processing domain: google.com
   Searching for favicon...
   Trying: https://google.com/favicon.ico
┌─────────────────────────────────────────┐
│ Source: https://google.com/favicon.ico  │
│ Size:   15486                           │
├─────────────────────────────────────────┤
│ MMH3:   -1234567890                     │
│ SHA256: e3b0c44298fc1c149afbf4c8996...  │
│ MD5:    d41d8cd98f00b204e9800998ec...   │
└─────────────────────────────────────────┘
```

---

## 🔧 Use Cases

* 🛡️ **Security Research**

  * Favicon hashing for Shodan reconnaissance (`http.favicon.hash:-1234567890`)
  * Asset discovery and fingerprinting
  * Threat intelligence gathering

* 🔍 **Digital Forensics**

  * File integrity verification
  * Duplicate image detection
  * Evidence collection

* 🌐 **Web Development**

  * Favicon-based website identification
  * Content verification
  * Cache busting with hash checks

* ⚡ **Penetration Testing**

  * Target enumeration
  * Attack surface mapping
  * Security assessment

---

## 🏗️ Architecture

```
FavMaster
├── Downloader (HTTP client)
├── Scanner (URL discovery)
├── Hasher (MMH3, SHA256, MD5)
└── Output Formatter
```

---

## 🛠️ Development

### Building from Source

```bash
# Clone repository
git clone https://github.com/cristophercervantes/favmaster
cd favmaster

# Build
go build -o favmaster cmd/favmaster

# Test
go test ./...

# Install locally
go install ./cmd/favmaster
```

### Make Commands

```bash
make build        # Build binary
make install      # Install using go install
make test         # Run tests
make clean        # Clean build artifacts
make build-all    # Cross-compile for all platforms
```

---

## 🤝 Contributing

We welcome contributions! Please feel free to submit pull requests.

**Suggested workflow**:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Setup

```bash
git clone https://github.com/cristophercervantes/favmaster
cd favmaster
go mod tidy
make test
```

---

## 🐛 Troubleshooting

**Common Issues**

* `command not found: favmaster`

```bash
# Add GOPATH to your PATH
export PATH=$PATH:$(go env GOPATH)/bin
# Add to your ~/.bashrc or ~/.zshrc for persistence
```

* `go: cannot find main module`

```bash
# Make sure you're in the project directory
cd favmaster
```

* **Network Timeouts**

```bash
# The tool includes retry logic, but you can adjust timeouts in the code
# Modify internal/downloader/downloader.go for custom timeouts
```

---

## 📄 License

This project is licensed under the MIT License - see the `LICENSE` file for details.

---

## 👨‍💻 Author

**Cristopher**
GitHub: [@cristophercervantes](https://github.com/cristophercervantes)

---

## ⭐ Support

If you find this tool useful, please give it a star on GitHub!

**Pro Tip:** Use with Shodan for powerful reconnaissance: `http.favicon.hash:YOUR_MMH3_HASH`

Happy Hashing! 🔍✨
