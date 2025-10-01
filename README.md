# FavMaster 🔍

A powerful tool for extracting MMH3, SHA256, and MD5 hashes from images, favicons, and URLs. Perfect for security researchers, penetration testers, and developers.

[![Version](https://img.shields.io/badge/Version-1.2.0-blue)](https://github.com/cristophercervantes/favmaster)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Linux%2520%257C%2520Windows%2520%257C%2520macOS-lightgrey)]()
![CI](https://github.com/cristophercervantes/favmaster/actions/workflows/go.yml/badge.svg)

---

## ✨ Features

* 🔍 **Automatic Favicon Discovery** — Intelligently finds favicons from domains
* 🖼️ **Multi-Format Support** — PNG, JPG, JPEG, GIF, SVG, WEBP, ICO
* 🌐 **Flexible Input Sources** — URLs, domains, text files, or local files
* 📊 **Multiple Hash Algorithms** — MMH3, SHA256, and MD5 calculations
* 🚀 **High Performance** — Fast, concurrent processing with robust error handling
* 🔧 **Easy Installation** — Simple `go install` command
* 🎯 **Smart URL Processing** — Auto-detects domains and finds favicons automatically
* 🏗️ **CI/CD Built-in** — Automated testing and multi-platform releases

---

## 🚀 Installation

### Quick Install (Recommended)

```bash
go install github.com/cristophercervantes/favmaster@latest
```

Make sure `$GOPATH/bin` is in your `PATH` environment variable.

### From Pre-built Binaries

Visit the Releases page to download pre-built binaries for:

* 🐧 Linux (AMD64, ARM64)
* 🪟 Windows (AMD64)
*  macOS (AMD64, ARM64)

### From Source

```bash
git clone https://github.com/cristophercervantes/favmaster
cd favmaster
go build -o favmaster .
sudo mv favmaster /usr/local/bin/
```

### Verify Installation

```bash
favmaster --help
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

# Multiple domains from command line (one by one)
favmaster google.com && favmaster github.com
```

### Text File Input Format

Create a text file with one URL or domain per line (comments supported with `#`):

```
# Popular websites for testing
google.com
https://github.com
http://example.com

# Specific image URLs
https://example.com/logo.png
https://site.com/images/favicon.ico

# Mixed entries - tool auto-detects type
stackoverflow.com
assets.example.com/favicon.jpg
localhost:8080

# Comments are ignored
# This is a comment line
```

---

## 🎯 Supported Image Formats

| Format  | Extension   | Common Use        |
| ------- | ----------- | ----------------- |
| Favicon | .ico        | Website icons     |
| PNG     | .png        | Web images, logos |
| JPEG    | .jpg, .jpeg | Photographs       |
| GIF     | .gif        | Animations        |
| SVG     | .svg        | Vector graphics   |
| WebP    | .webp       | Modern web format |

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

### 🛡️ Security Research & Reconnaissance

```bash
# Shodan searches using MMH3 hashes
favmaster target.com
# Use output in Shodan: http.favicon.hash:-1234567890

# Bulk processing for attack surface mapping
favmaster target_list.txt
```

### 🔍 Digital Forensics & Incident Response

* File integrity verification across multiple systems
* Duplicate image detection in evidence collection
* Evidence hashing for chain of custody

### 🌐 Web Development & Operations

* Favicon-based website identification in logs
* Content verification after deployments
* Cache busting with hash-based versioning

### ⚡ Penetration Testing & Red Teaming

```bash
# Quick asset discovery
favmaster company-domain.com

# Bulk target processing
cat domains.txt | xargs -I {} favmaster {}

# Integration with other tools
favmaster target.com | grep MMH3 | awk '{print $2}' > hashes.txt
```

---

## 🏗️ Architecture

```
FavMaster Core
├── HTTP Client (Downloader)
├── URL Scanner & Favicon Discoverer
├── Hash Calculator (MMH3, SHA256, MD5)
└── Output Formatter
```

---

## 🐛 Troubleshooting

### Common Issues & Solutions

**"command not found: favmaster"**

```bash
# Add GOPATH to your PATH
export PATH=$PATH:$(go env GOPATH)/bin

# Make it permanent (add to ~/.bashrc, ~/.zshrc, or ~/.profile)
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc

# Verify Go environment
go env GOPATH
```

**Network Timeouts or Connection Issues**

```bash
# The tool has 30-second timeouts with retry logic
# Check your internet connection and firewall settings

# For corporate environments, you might need to set proxies
export HTTP_PROXY=http://your-proxy:port
export HTTPS_PROXY=http://your-proxy:port
```

**Permission Issues**

```bash
# If you can't move the binary to /usr/local/bin
sudo mv favmaster /usr/local/bin/

# Or install to a user directory
mkdir -p ~/.local/bin
mv favmaster ~/.local/bin/
export PATH=$PATH:~/.local/bin
```

**Go Version Compatibility**

```bash
# Check your Go version
go version

# Update Go if needed (requires 1.21+)
# Visit: https://golang.org/dl/
```

---

## 🤝 Contributing

We love contributions! Here's how to get started:

### Development Setup

```bash
# Fork and clone
git clone https://github.com/cristophercervantes/favmaster
cd favmaster

# Build and test
go build -o favmaster .
./favmaster google.com  # Test functionality

# Run basic validation
go test -v .
```

### Contribution Workflow

* Fork the repository
* Create a feature branch (`git checkout -b feature/amazing-feature`)
* Commit your changes (`git commit -m 'Add amazing feature'`)
* Push to the branch (`git push origin feature/amazing-feature`)
* Open a Pull Request

### Areas for Contribution

* 🆕 New hash algorithms
* 🌐 Additional protocol support
* 📊 Enhanced output formats (JSON, CSV)
* 🔧 Performance optimizations
* 🐛 Bug fixes and documentation

---

## 🛠️ API Reference

### Command Line Options

```bash
favmaster <input>  # Process URL, domain, file, or text file
```

### Input Types

* **Domain:** `example.com` (auto-discovers favicon)
* **URL:** `https://example.com/favicon.ico` (direct access)
* **File:** `image.png` (local file processing)
* **Text File:** `urls.txt` (batch processing)

### Exit Codes

* `0`: Success
* `1`: General error
* `2`: Invalid arguments
* `3`: Network error
* `4`: File system error

---

## 📄 License

This project is licensed under the MIT License - see the `LICENSE` file for details.

---

## 👨‍💻 Author

Cristopher

GitHub: [@cristophercervantes](https://github.com/cristophercervantes)

Tool: FavMaster

---

## 🙏 Acknowledgments

* Go Community - For excellent libraries and tools
* Shodan - For popularizing favicon hashing techniques
* Security Researchers - For inspiration and use cases
* Contributors - For improvements and bug reports

---

## ⭐ Support

If you find this tool useful, please give it a star on GitHub!

### Pro Tip for Security Researchers:

```bash
# Combine with Shodan for powerful reconnaissance
favmaster target.com | grep MMH3 | awk '{print "http.favicon.hash:" $2}' | shodan search --fields ip_str,port,org -
```

Happy Hashing! 🔍✨

**FavMaster - Because every favicon tells a story**

Release page: [https://github.com/cristophercervantes/favmaster/releases/](https://github.com/cristophercervantes/favmaster/releases/)
