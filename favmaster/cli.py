#!/usr/bin/env python3

import sys
from pathlib import Path
from .downloader import download_file
from .hasher import calculate_hashes, display_hashes
from .scanner import find_and_process_favicon, read_urls_from_file, is_image_url

def show_banner():
    banner = """
╔═══════════════════════════════════════╗
║              FAVMASTER              ║
║            v1.0.0 by Cristopher           ║
╚═══════════════════════════════════════╝
"""
    print(banner)

def show_usage():
    print("\nUsage:")
    print("  favmaster <url/domain>  - Process single URL or domain")
    print("  favmaster <file.txt>    - Process URLs from text file")
    print("  favmaster <file>        - Process local file")
    print("\nSupported image formats: .ico, .png, .jpg, .jpeg, .gif, .svg, .webp")
    print("\nExamples:")
    print("  favmaster https://example.com/favicon.ico")
    print("  favmaster example.com")
    print("  favmaster https://example.com/image.png")
    print("  favmaster urls.txt")
    print("  favmaster image.jpg")

def is_url(input_str):
    return input_str.startswith(('http://', 'https://'))

def process_url_or_domain(input_str):
    if is_url(input_str):
        if is_image_url(input_str):
            print(f"\n🔗 Processing image URL: {input_str}")
            process_single_url(input_str)
        else:
            print(f"\n🌐 Processing domain: {input_str}")
            find_and_process_favicon(input_str)
    else:
        print(f"\n🌐 Processing domain: {input_str}")
        find_and_process_favicon(f"https://{input_str}")

def process_single_url(url):
    try:
        data = download_file(url)
        if not data:
            print(f"❌ Empty response from {url}")
            return
        
        calculate_and_display_hashes(data, url)
    except Exception as e:
        print(f"❌ Error downloading {url}: {e}")

def process_url_file(filename):
    print(f"\n📄 Processing URLs from: {filename}")
    try:
        urls = read_urls_from_file(filename)
        processed = 0
        
        for url in urls:
            process_url_or_domain(url)
            processed += 1
        
        print(f"\n✅ Processed {processed}/{len(urls)} entries from {filename}")
    except Exception as e:
        print(f"❌ Error reading file {filename}: {e}")

def process_local_file(filename):
    print(f"\n📁 Processing local file: {filename}")
    try:
        with open(filename, 'rb') as f:
            data = f.read()
        
        if not data:
            print(f"❌ Empty file: {filename}")
            return
        
        calculate_and_display_hashes(data, filename)
    except Exception as e:
        print(f"❌ Error reading file {filename}: {e}")

def calculate_and_display_hashes(data, source):
    hashes = calculate_hashes(data)
    display_hashes(hashes, source, len(data))

def main():
    if len(sys.argv) < 2:
        show_banner()
        show_usage()
        sys.exit(1)
    
    show_banner()
    input_str = sys.argv[1]
    
    # Check if it's a text file
    if input_str.lower().endswith('.txt'):
        process_url_file(input_str)
    # Check if it's a local file that exists
    elif Path(input_str).exists():
        process_local_file(input_str)
    # Assume it's a URL or domain
    else:
        process_url_or_domain(input_str)

if __name__ == "__main__":
    main()
