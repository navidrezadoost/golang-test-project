package utils

import (
	"fmt"
)

// ByteToSize converts a byte size into a human-readable string representation.
// It accepts an int64 representing the size in bytes and returns a string
// indicating the size in a convenient format (e.g., "5 MB", "1.2 GB").
func ByteToSize(byte int64) string {
	const (
		KB float64 = 1024
		MB         = KB * 1024
		GB         = MB * 1024
		TB         = GB * 1024
	)

	if byte == 0 {
		return "n/a"
	}

	var unit string
	var value float64

	switch {
	case float64(byte) >= TB:
		unit = "TB"
		value = float64(byte) / TB
	case float64(byte) >= GB:
		unit = "GB"
		value = float64(byte) / GB
	case float64(byte) >= MB:
		unit = "MB"
		value = float64(byte) / MB
	case float64(byte) >= KB:
		unit = "KB"
		value = float64(byte) / KB
	default:
		unit = "Bytes"
		value = float64(byte)
	}

	if value == float64(int64(value)) {
		return fmt.Sprintf("%d %s", int64(value), unit)
	}
	return fmt.Sprintf("%.1f %s", value, unit)
}

// BytesToSize converts a list of byte sizes into a human-readable string representation
// of their total combined size. It accepts a slice of int64 values representing sizes in bytes
// and returns a string indicating the total size in a convenient format (e.g., "5 MB", "1.2 GB").
func BytesToSize(bytes []int64) string {
	totalBytes := int64(0)
	for _, b := range bytes {
		totalBytes += b
	}

	const (
		_        = iota // Ignore zero
		kb int64 = 1 << (10 * iota)
		mb
		gb
		tb
	)

	switch {
	case totalBytes >= tb:
		return fmt.Sprintf("%.1f TB", float64(totalBytes)/float64(tb))
	case totalBytes >= gb:
		return fmt.Sprintf("%.1f GB", float64(totalBytes)/float64(gb))
	case totalBytes >= mb:
		return fmt.Sprintf("%.1f MB", float64(totalBytes)/float64(mb))
	case totalBytes >= kb:
		return fmt.Sprintf("%.1f KB", float64(totalBytes)/float64(kb))
	default:
		return fmt.Sprintf("%d Bytes", totalBytes)
	}
}
