package cmd

import (
	"fmt"
	"testing"
)

// Test fails if an RFC3339 datetime is not extracted.
// To get: TZ=Zulu date +%Y-%m-%dT%H:%m:%SZ
func TestReadRainbow(t *testing.T) {
}

func TestListAlmanac(t *testing.T) {
	fmt.Println(listAlmanac())
}
