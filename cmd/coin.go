/*
Copyright © 2021 Matt Davis <maroda@rainbowq.io>
*/
package cmd

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"

	b64 "encoding/base64"

	"github.com/spf13/cobra"
)

// coinCmd represents the coin command
var coinCmd = &cobra.Command{
	Use:   "coin",
	Short: "Toss a coin. Use the '-t|--token' flag for a randomized 32 character token.",
	Run: func(cmd *cobra.Command, args []string) {
		c, t := RndCoin(time.Now().UnixNano())
		if tokenOut {
			fmt.Println(t)
		} else if tok64Out {
			fmt.Println(b64.URLEncoding.EncodeToString([]byte(t)))
		} else {
			switch c {
			case 0:
				fmt.Println("TAILS")
			case 1:
				fmt.Println("HEADS")
			}
		}
	},
}

var tokenOut bool
var tok64Out bool

func init() {
	rootCmd.AddCommand(coinCmd)
	coinCmd.PersistentFlags().BoolVarP(&tokenOut, "token", "t", false, "print a 32-character token instead of a coin")
	coinCmd.PersistentFlags().BoolVarP(&tok64Out, "base64", "b", false, "print a base64 URL-encoded 32-char token instead of a coin")
}

// RndCoin ::: Produce a coin toss.
//	Takes a salt value to create a second salt value to
//	randomize a third number that is taken as odd or even
//	and returned respectively as heads (1) or tails (0),
//	along with a 32 character MD5 of the final random output.
func RndCoin(salt int64) (int, string) {
	rand.Seed(salt)               // seed math.rand
	input := rand.Int()           // get a random int
	tlas := time.Now().UnixNano() // timestamp nanoseconds
	rand.Seed(tlas)               // re-seed math.rand
	output := rand.Intn(input)    // get a second random int
	encode := MD5ken(fmt.Sprint(output))

	if output%2 == 0 {
		return 0, encode // TAILS
	} else {
		return 1, encode // HEADS
	}
}

// MD5ken ::: Creates an MD5 hash appropriate as a
//	nice 32-character random string for token use.
func MD5ken(k string) string {
	s := md5.New()
	s.Write([]byte(k))
	bash := s.Sum(nil)
	hash := fmt.Sprintf("%x", bash)
	return hash
}
