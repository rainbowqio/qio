package cmd

import (
	"testing"
	"time"
)

// TestMD5ken ::: Fails if it's not a 32 character token.
func TestMD5ken(t *testing.T) {
	ttxt := "craquemattic"
	ttkn := MD5ken(ttxt)

	if len(ttkn) != 32 {
		t.Errorf("Token length error, expected %d got %d", 32, len(ttkn))
	}
}

// TestRndCoin ::: Fails if not a 0 or 1
func TestRndCoin(t *testing.T) {
	coin, _ := RndCoin(time.Now().UnixNano())

	if coin < 0 {
		t.Errorf("Impossible value for coin, got %d", coin)
	} else if coin > 1 {
		t.Errorf("Impossible range for coin, got %d", coin)
	}
}
