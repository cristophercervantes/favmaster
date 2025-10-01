import time
from pathlib import Path
from bs4 import BeautifulSoup
from .downloader import download_file, check_url_exists, get_content_type, normalize_url, resolve_url

# Supported image extensions
SUPPORTED_IMAGE_EXTENSIONS = {
    '.ico', '.png', '.jpg', '.jpeg', '.gif', '.svg', '.webp'
}

# Common favicon paths
FAVICON_PATHS = [
    '/favicon.ico',
    '/favicon.png',
    '/favicon.jpg',
    '/favicon.jpeg',
    '/favicon.svg',
    '/favicon.webp',
    '/images/favicon.ico',
    '/img/favicon.ico',
    '/assets/favicon.ico',
    '/static/favicon.ico',
    '/icon.png',
    '/logo.png',
    '/apple-touch-icon.png',
]

def is_image_url(url):
    """
    Check if URL points to an image based on extension.
    """
    url_lower = url.lower()
    return any(url_lower.endswith(ext) for ext in SUPPORTED_IMAGE_EXTENSIONS)

def find_and_process_favicon(domain):
    """
    Find and process favicon for a domain.
    """
    print("   Searching for favicon...")
    
    # Try common favicon paths first
    found = try_common_favicon_paths(domain)
    if not found:
        # If not found, parse HTML to find favicon
        found = parse_html_for_favicon(domain)
    
    if not found:
        print(f"❌ No favicon found for {domain}")

def try_common_favicon_paths(domain):
    """
    Try common favicon paths for a domain.
    """
    for path in FAVICON_PATHS:
        favicon_url = normalize_url(domain) + path
        print(f"   Trying: {favicon_url}")
        
        if check_and_process_image_url(favicon_url):
            return True
        
        # Small delay to be respectful to servers
        time.sleep(0.1)
    return False

def parse_html_for_favicon(domain):
    """
    Parse HTML to find favicon links.
    """
    try:
        data = download_file(domain)
        soup = BeautifulSoup(data, 'html.parser')
        found = False
        
        # Look for favicon in link tags
        for link in soup.find_all('link', rel=True):
            rel = link.get('rel', [])
            href = link.get('href', '')
            
            # Check for various favicon rel types
            favicon_rels = ['icon', 'shortcut icon', 'apple-touch-icon', 'apple-touch-icon-precomposed']
            if any(icon_type in favicon_rels for icon_type in rel) and href:
                favicon_url = resolve_url(domain, href)
                print(f"   Found in HTML: {favicon_url}")
                if check_and_process_image_url(favicon_url):
                    found = True
        
        return found
    except Exception as e:
        print(f"   Error parsing HTML: {e}")
        return False

def check_and_process_image_url(url):
    """
    Check if URL is an image and process it.
    """
    from .cli import process_single_url
    
    if not check_url_exists(url):
        return False
    
    content_type = get_content_type(url)
    if content_type and content_type.startswith('image/'):
        process_single_url(url)
        return True
    
    return False

def read_urls_from_file(filename):
    """
    Read URLs from a text file.
    """
    urls = []
    try:
        with open(filename, 'r') as f:
            for line in f:
                line = line.strip()
                if line and not line.startswith('#'):
                    urls.append(line)
        return urls
    except Exception as e:
        raise Exception(f"Could not read file {filename}: {e}")
