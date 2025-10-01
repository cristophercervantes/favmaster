package hasher

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/tj/go-murmur"
)

type HashResult struct {
	Source string
	Size   int
	MMH3   uint32
	SHA256 string
	MD5    string
}

func CalculateHashes(data []byte, source string) HashResult {
	mmh3Hash := murmur.Hash32(data, 0)
	sha256Hash := sha256.Sum256(data)
	md5Hash := md5.Sum(data)

	return HashResult{
		Source: source,
		Size:   len(data),
		MMH3:   mmh3Hash,
		SHA256: hex.EncodeToString(sha256Hash[:]),
		MD5:    hex.EncodeToString(md5Hash[:]),
	}
}

func CalculateAndDisplayHashes(data []byte, source string) {
	result := CalculateHashes(data, source)
	DisplayHashes(result)
}

func DisplayHashes(result HashResult) {
	fmt.Println("┌─────────────────────────────────────────┐")
	fmt.Printf("│ Source: %-30s │\n", truncateString(result.Source, 30))
	fmt.Printf("│ Size:   %-30d │\n", result.Size)
	fmt.Println("├─────────────────────────────────────────┤")
	fmt.Printf("│ MMH3:   %-32d │\n", result.MMH3)
	fmt.Printf("│ SHA256: %-32s │\n", truncateString(result.SHA256, 32))
	fmt.Printf("│ MD5:    %-32s │\n", truncateString(result.MD5, 32))
	fmt.Println("└─────────────────────────────────────────┘")
}

func truncateString(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength-3] + "..."
}
