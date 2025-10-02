import hashlib
import mmh3

def murmurhash3_x86_32(data, seed=0):
    """
    Calculate MurmurHash3 (32-bit) for the given data.
    Returns unsigned integer (compatible with Shodan).
    """
    hash_signed = mmh3.hash(data, seed)
    
    # Convert signed to unsigned for Shodan compatibility
    if hash_signed < 0:
        hash_unsigned = hash_signed + 2**32
    else:
        hash_unsigned = hash_signed
    
    return hash_unsigned

def calculate_hashes(data):
    """
    Calculate MMH3, SHA256, and MD5 hashes for the given data.
    """
    # Calculate MurmurHash3 (mmh3) - unsigned for Shodan
    mmh3_hash = murmurhash3_x86_32(data)
    
    # Calculate SHA256
    sha256_hash = hashlib.sha256(data).hexdigest()
    
    # Calculate MD5
    md5_hash = hashlib.md5(data).hexdigest()
    
    return {
        'mmh3': mmh3_hash,
        'sha256': sha256_hash,
        'md5': md5_hash
    }

def display_hashes(hashes, source, size):
    """
    Display hashes in a formatted box.
    """
    source_display = truncate_string(source, 30)
    
    print("┌─────────────────────────────────────────┐")
    print(f"│ Source: {source_display:<30} │")
    print(f"│ Size:   {size:<30} │")
    print("├─────────────────────────────────────────┤")
    print(f"│ MMH3:   {hashes['mmh3']:<32} │")
    print(f"│ SHA256: {truncate_string(hashes['sha256'], 32):<32} │")
    print(f"│ MD5:    {truncate_string(hashes['md5'], 32):<32} │")
    print("└─────────────────────────────────────────┘")

def truncate_string(s, max_length):
    """
    Truncate string to max_length, adding ... if truncated.
    """
    if len(s) <= max_length:
        return s
    return s[:max_length-3] + "..."
