# FavMaster 🔍

![Python 3.7+](https://img.shields.io/badge/Python-3.7+-blue)
![License MIT](https://img.shields.io/badge/License-MIT-green)
![Status Ready for PyPI](https://img.shields.io/badge/Status-Ready_for_PyPI-orange)

A powerful, feature-rich Python tool for extracting MMH3, SHA256, and MD5 hashes from images, favicons, and URLs. Perfect for security researchers, penetration testers, and developers.

## ✨ Features

* **Automatic Favicon Discovery** — Intelligently finds favicons from domains.
* **Multi-Format Support** — PNG, JPG, JPEG, GIF, SVG, WEBP, ICO.
* **Flexible Input Sources** — URLs, domains, text files, or local files.
* **Multiple Hash Algorithms** — MMH3, SHA256, and MD5 calculations.
* **High Performance** — Fast, concurrent processing with robust error handling.
* **Easy Installation** — Multiple installation options (pipx / pip / local / dev).
* **Smart URL Processing** — Auto-detects domains and finds favicons automatically.
* **Bulk Processing** — Process multiple targets from text files.
* **Security Focused** — Built for security professionals and researchers.

---

## 🚀 Installation

**Method 1: Install from GitHub using pipx (Recommended - Isolated Environment)**

```bash
pipx install git+https://github.com/cristophercervantes/favmaster.git
```

**Method 2: Install from GitHub using pip**

```bash
pip install git+https://github.com/cristophercervantes/favmaster.git
```

**Method 3: Install from Local Source**

```bash
# Clone the repository
git clone https://github.com/cristophercervantes/favmaster
cd favmaster

# Install using pip
pip install .

# Or install using pipx (isolated environment)
pipx install .
```

**Method 4: Development Installation**

```bash
git clone https://github.com/cristophercervantes/favmaster
cd favmaster
pip install -e .  # Editable mode for development
```

### 🔧 Prerequisites

Install `pipx` (if you don't have it)

```bash
# On Linux/macOS
python3 -m pip install --user pipx
python3 -m pipx ensurepath

# On Windows
py -m pip install --user pipx
py -m pipx ensurepath

# Restart your terminal after installation
```

**Why Choose pipx?**

* No dependency conflicts with other Python packages
* Global command — `favmaster` available everywhere
* Easy updates — `pipx upgrade favmaster`
* Clean removal — `pipx uninstall favmaster`
* Security — Runs in isolated environment

---

## 📖 Quick Start

### Basic Usage

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

### Advanced Usage

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

---

## 🎯 Supported Image Formats

|  Format | Extension   | Common Use                     |
| ------: | ----------- | ------------------------------ |
| Favicon | .ico        | Website icons                  |
|     PNG | .png        | Web images, logos, screenshots |
|    JPEG | .jpg, .jpeg | Photographs, complex images    |
|     GIF | .gif        | Animations, simple graphics    |
|     SVG | .svg        | Vector graphics, logos         |
|    WebP | .webp       | Modern web format, optimized   |

---

## 📊 Output Example

```
╔═══════════════════════════════════════╗
║              FAVMASTER                ║
║            v1.0.0 by Cristopher       ║
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

## 📁 Input File Format

Create a text file with one URL or domain per line (comments supported with `#`):

```
# Popular websites for reconnaissance
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
# This line won't be processed
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

# Integration with other security tools
favmaster company.com | grep MMH3 | awk '{print $3}' > hashes.txt
```

### 🔍 Digital Forensics & Incident Response

* File integrity verification across multiple systems
* Duplicate image detection in evidence collection
* Evidence hashing for chain of custody
* Malware analysis — hash known malicious images

### 🌐 Web Development & Operations

* Favicon-based website identification in logs
* Content verification after deployments
* Cache busting with hash-based versioning
* Asset management and tracking

### ⚡ Penetration Testing & Red Teaming

```bash
# Quick asset discovery during engagements
favmaster target-company.com

# Bulk target processing for large scopes
cat domains.txt | xargs -I {} favmaster {}

# Integration with reconnaissance pipelines
subfinder -d target.com | favmaster
```

---

## 🛠️ Advanced Usage

### Integration with Other Tools

```bash
# Combine with subdomain enumeration
subfinder -d example.com | favmaster

# Process results from asset discovery tools
assetfinder example.com | favmaster

# Use with parallel processing for large datasets
cat large_domain_list.txt | parallel -j 10 favmaster
```

---

## 🐛 Troubleshooting

**Common Issues & Solutions**

* `"Command not found: favmaster"`

```bash
# If using pipx, ensure it's in PATH
pipx ensurepath

# If using pip, ensure Python scripts are in PATH
export PATH=$PATH:$(python -m site --user-base)/bin

# Or use Python module directly
python -m favmaster.cli google.com
```

* `pipx not found`

```bash
# Install pipx first
python3 -m pip install --user pipx
python3 -m pipx ensurepath
source ~/.bashrc  # or restart terminal
```

* **Network Timeouts**

```bash
# The tool includes 30-second timeouts with retry logic
# Check your internet connection and firewall settings

# For corporate environments, set proxies
export HTTP_PROXY=http://your-proxy:port
export HTTPS_PROXY=http://your-proxy:port
```

---

## 🔄 Updating

**Update with pipx**

```bash
pipx upgrade favmaster
```

**Update with pip**

```bash
pip install --upgrade git+https://github.com/cristophercervantes/favmaster.git
```

---

## 🤝 Contributing

We welcome contributions! Here's how to get started:

**Development Setup**

```bash
# Fork and clone
git clone https://github.com/cristophercervantes/favmaster
cd favmaster

# Install in development mode
pip install -e .

# Make your changes and test
```

**Pull Request Process**

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

---

## 📄 License

MIT License — see LICENSE file for details.

---

## 👨‍💻 Author

Cristopher
GitHub: [@cristophercervantes](https://github.com/cristophercervantes)

---

If you want this README saved with a different filename, or edited (shorter, longer, translated, or tailored for PyPI), tell me and I will update it.
