# FavMaster

[![Python](https://img.shields.io/badge/python-3.6%2B-blue.svg)](https://www.python.org/downloads/)
[![Platform](https://img.shields.io/badge/platform-linux%20%7C%20macos%20%7C%20windows-lightgrey)]()
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## FavMaster is a command-line tool to compute hashes from website
favicons. It supports:

-   **MMH3** (Shodan/FOFA/Censys/CriminalIP-compatible, using
    base64-encoded data with newlines)
-   **MD5**
-   **SHA256**

You can provide either a root URL (e.g., `https://example.com`) to
automatically discover and download the favicon, or a local favicon file
(e.g., `.ico`, `.png`, `.jpg`, `.jpeg`, `.webp`, `.svg`).

**Creator**: Cristopher

## Features

-   **URL Mode**: Automatically discovers favicon URLs by parsing
    `<link rel="icon">` tags or falling back to `/favicon.ico`.
-   **File Mode**: Reads local favicon files in various formats (.ico,
    .png, .jpg, .jpeg, .webp, .svg, or any binary file).
-   **Hashes**:
    -   MMH3: Shodan-compatible (base64-encoded favicon data, including
        newlines, then hashed with MurmurHash3).
    -   MD5 and SHA256: Computed on raw favicon bytes.
-   **Installable**: Available via `pip` or `pipx` from GitHub for easy
    global access.
-   **Cross-Platform**: Works on Windows, macOS, and Linux.

## Installation

FavMaster can be installed directly from GitHub using `pip` or `pipx`.
The latter is recommended for CLI tools to keep dependencies isolated.

### Prerequisites

-   Python 3.6 or higher
-   `pip` or `pipx` installed
-   Ensure your Python `bin` directory is in your system PATH (e.g.,
    `~/.local/bin` on Linux/macOS or `Scripts` in your Python install on
    Windows).

### Install via pip

``` bash
pip install git+https://github.com/yourusername/favmaster.git
```

### Install via pipx (recommended for CLI tools)

``` bash
pipx install git+https://github.com/yourusername/favmaster.git
```

### Local Installation (for development)

Clone the repository:

``` bash
git clone https://github.com/yourusername/favmaster.git
cd favmaster
```

Install locally:

``` bash
pip install .
```

## Usage

Run the tool using the favmaster command, followed by either a URL or a
file path.

### Syntax

``` bash
favmaster <input>
```

`<input>`: Either a root URL (e.g., https://example.com) or a local file
path (e.g., ./favicon.ico).

### Examples

#### Hash a favicon from a URL:

``` bash
favmaster https://example.com
```

Output:

    text______              ___  ____          
      |  ____|            |__ \|___ \         
      | |__ __ ___   ____  __) | __) | ___ 
      |  __| '_ ` _ \ / __||__ <|__ < / __|
      | |  | | | | | | (__ ___) |__) | (__ 
      | |_|  |_| |_| |_| \___|____/____/ \___|

      FavMaster – MMH3 (Shodan/FOFA/Censys/CriminalIP), MD5, SHA256
      Creator: Cristopher

    [+] Discovered favicon URL: https://example.com/favicon.ico

    Hashes
    ----------------------------------------
    MMH3 (Shodan/FOFA/Censys/CriminalIP): -1370005593
    MD5: d41d8cd98f00b204e9800998ecf8427e
    SHA256: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855

#### Hash a local favicon file:

``` bash
favmaster ./myfavicon.ico
```

Output:

    text______              ___  ____          
      |  ____|            |__ \|___ \         
      | |__ __ ___   ____  __) | __) | ___ 
      |  __| '_ ` _ \ / __||__ <|__ < / __|
      | |  | | | | | | (__ ___) |__) | (__ 
      | |_|  |_| |_| |_| \___|____/____/ \___|

      FavMaster – MMH3 (Shodan/FOFA/Censys/CriminalIP), MD5, SHA256
      Creator: Cristopher

    [+] Reading local file: ./myfavicon.ico

    Hashes
    ----------------------------------------
    MMH3 (Shodan/FOFA/Censys/CriminalIP): -1370005593
    MD5: d41d8cd98f00b204e9800998ecf8427e
    SHA256: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855

## Supported File Formats

-   .ico
-   .png
-   .jpg/.jpeg
-   .webp
-   .svg
-   Any binary file (hashes raw bytes)

## Dependencies

Listed in requirements.txt:

-   requests\>=2.28.0: For downloading favicons from URLs.
-   beautifulsoup4\>=4.11.0: For parsing HTML to find favicon URLs.
-   mmh3\>=4.0.0: For computing MurmurHash3 (MMH3).

These are automatically installed during setup.

## How It Works

### URL Mode:

If the input starts with http:// or https://, FavMaster:

-   Fetches the root URL.
-   Parses HTML for `<link rel="icon">`{=html} (or similar,
    case-insensitive).
-   Falls back to /favicon.ico if no link is found.
-   Downloads the favicon.

### File Mode:

If the input is a file path, FavMaster reads the file directly.

### Hashing:

-   MD5 and SHA256: Computed on raw favicon bytes.
-   MMH3: Base64-encodes the favicon (with newlines, per
    Shodan/FOFA/Censys/CriminalIP convention), then applies MurmurHash3.

### Output:

Displays the hashes in a clean format with a banner.

## Troubleshooting

-   **Command not found**: Ensure your Python bin directory is in your
    PATH:
    -   Linux/macOS: `export PATH=$PATH:~/.local/bin`
    -   Windows: Add `C:\Path\To\Python\Scripts` to your system PATH.
-   **Network errors**: Check your internet connection or the URL's
    validity.
-   **File not found**: Verify the file path is correct and the file
    exists.

## Contributing

Contributions are welcome! To contribute:

1.  Fork the repository.
2.  Create a new branch: `git checkout -b feature/your-feature`.
3.  Make changes and commit: `git commit -m "Add your feature"`.
4.  Push to your fork: `git push origin feature/your-feature`.
5.  Open a pull request.

Please include tests and update documentation as needed.

## License

MIT License. See LICENSE file (you can add one to the repository).

## Contact

For issues or feature requests, open a GitHub issue or contact
Cristopher via GitHub.
