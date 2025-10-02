import requests
from urllib.parse import urljoin, urlparse

def download_file(url, timeout=30):
    """
    Download file from URL and return bytes.
    Use consistent headers to get the same content every time.
    """
    headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36',
        'Accept': 'image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8',
        'Accept-Language': 'en-US,en;q=0.9',
        'Accept-Encoding': 'gzip, deflate, br',
        'DNT': '1',
        'Connection': 'keep-alive',
        'Sec-Fetch-Dest': 'image',
        'Sec-Fetch-Mode': 'no-cors',
        'Sec-Fetch-Site': 'cross-site',
        'Cache-Control': 'no-cache',
        'Pragma': 'no-cache'
    }
    
    try:
        response = requests.get(url, headers=headers, timeout=timeout)
        response.raise_for_status()
        return response.content
    except requests.exceptions.RequestException as e:
        raise Exception(f"Failed to download {url}: {e}")

def check_url_exists(url, timeout=10):
    """
    Check if URL exists and returns successful status code.
    """
    try:
        headers = {
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36',
            'Accept': '*/*',
            'Cache-Control': 'no-cache'
        }
        response = requests.head(url, headers=headers, timeout=timeout, allow_redirects=True)
        return response.status_code == 200
    except:
        return False

def get_content_type(url, timeout=10):
    """
    Get content type of URL.
    """
    try:
        headers = {
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36',
            'Accept': '*/*',
            'Cache-Control': 'no-cache'
        }
        response = requests.head(url, headers=headers, timeout=timeout, allow_redirects=True)
        return response.headers.get('content-type', '')
    except:
        return ''

def normalize_url(domain):
    """
    Normalize URL by adding https:// if missing.
    """
    if domain.startswith(('http://', 'https://')):
        return domain
    return f"https://{domain}"

def resolve_url(base, relative):
    """
    Resolve relative URL against base URL.
    """
    return urljoin(base, relative)
