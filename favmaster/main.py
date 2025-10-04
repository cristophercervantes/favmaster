#!/usr/bin/env python3
import argparse
import base64
import hashlib
import mmh3
import os
import sys
import requests
from bs4 import BeautifulSoup
from urllib.parse import urljoin

BANNER = r"""
   ______              ___  ____          
  |  ____|            |__ \|___ \         
  | |__ __ ___   ____  __) | __) | ___ 
  |  __| '_ ` _ \ / __||__ <|__ < / __|
  | |  | | | | | | (__ ___) |__) | (__ 
  |_|  |_| |_| |_| \___|____/____/ \___|

  FavMaster – MMH3 (Shodan/FOFA/Censys/CriminalIP), MD5, SHA256
  Creator: Cristopher
"""

def print_banner():
    print(BANNER)

def get_favicon_url(root_url: str) -> str:
    """Discover the favicon URL by parsing HTML or fallback to /favicon.ico."""
    if not root_url.startswith(("http://", "https://")):
        root_url = "https://" + root_url

    try:
        resp = requests.get(root_url, timeout=10)
        resp.raise_for_status()
    except requests.RequestException as e:
        raise ValueError(f"Failed to fetch root URL: {e}")

    soup = BeautifulSoup(resp.text, "html.parser")
    link = soup.find(
        "link",
        attrs={"rel": lambda x: x and "icon" in x.lower()}
    )

    if link and link.get("href"):
        favicon_url = link["href"]
        if not favicon_url.startswith(("http://", "https://")):
            favicon_url = urljoin(root_url, favicon_url)
        return favicon_url

    # Default fallback
    return urljoin(root_url, "/favicon.ico")

def download_favicon(url: str) -> bytes:
    """Download the favicon binary."""
    try:
        resp = requests.get(url, timeout=10)
        resp.raise_for_status()
        return resp.content
    except requests.RequestException as e:
        raise ValueError(f"Failed to download favicon: {e}")

def read_local_file(path: str) -> bytes:
    """Read a local image file."""
    if not os.path.isfile(path):
        raise ValueError(f"File not found: {path}")
    with open(path, "rb") as f:
        return f.read()

def compute_hashes(data: bytes) -> dict:
    """
    Compute:
      * MD5  – raw bytes
      * SHA-256 – raw bytes
      * MMH3 – Shodan-compatible (base64-encode with newlines, then mmh3)
    """
    md5 = hashlib.md5(data).hexdigest()
    sha256 = hashlib.sha256(data).hexdigest()

    # Shodan-compatible MMH3
    b64 = base64.encodebytes(data)          # includes trailing newlines
    mmh3_val = mmh3.hash(b64)

    return {"mmh3": mmh3_val, "md5": md5, "sha256": sha256}

def main():
    print_banner()

    parser = argparse.ArgumentParser(
        description="Compute MMH3 (Shodan/FOFA/Censys/CriminalIP-compatible), MD5, and SHA256 hashes from a favicon."
    )
    parser.add_argument(
        "input",
        help="Root URL (e.g., https://example.com) or local file path (ico/png/jpg/webp/svg/etc.)"
    )
    args = parser.parse_args()

    try:
        if args.input.startswith(("http://", "https://")):
            # ---- URL mode -------------------------------------------------
            favicon_url = get_favicon_url(args.input)
            print(f"[+] Discovered favicon URL: {favicon_url}")
            data = download_favicon(favicon_url)
        else:
            # ---- File mode ------------------------------------------------
            print(f"[+] Reading local file: {args.input}")
            data = read_local_file(args.input)

        hashes = compute_hashes(data)

        print("\nHashes")
        print("-" * 40)
        print(f"MMH3 (Shodan/FOFA/Censys/CriminalIP): {hashes['mmh3']}")
        print(f"MD5: {hashes['md5']}")
        print(f"SHA256: {hashes['sha256']}")

    except Exception as e:
        print(f"[!] Error: {e}", file=sys.stderr)
        sys.exit(1)

if __name__ == "__main__":
    main()
