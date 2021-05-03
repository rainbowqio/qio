package cmd

import (
	"bytes"
	"sort"
	"testing"

	"github.com/spf13/viper"
)

// TestReadRainbow ::: Tests the argument set that gets passed around.
func TestReadRainbow(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"arg single", []string{"meta"}},
		{"arg diff", []string{"meta", "last"}},
		{"arg bad", []string{"meta", "..notanarg"}},
	}

	var acp string
	var display string

	for _, tts := range tests {
		switch tts.name {
		case "arg single":
			// When Run sends identical arguments, the user only used one.
			// So a single Almanac (no Plugs) should be returned.
			localacp := tts.args[0]
			acp, display, _ = readRainbow(tts.args[0], tts.args[0]) // cmd.Run sends this

			if acp != localacp {
				t.Errorf("Return value error, expected %s got %s", tts.args[0], acp)
			}

			// These were dummy values, they shouldn't be found.
			if display != "Not Found" {
				t.Errorf("Content error, expected %s got %s", "Not Found", display)
			}

		case "arg diff":
			// The correct format for Almanac:Plug is constructed with two different arguments.
			localacp := string(tts.args[0] + "." + tts.args[1])
			acp, display, _ = readRainbow(tts.args[0], tts.args[1]) // cmd.Run sends this

			if acp != localacp {
				t.Errorf("Return value error, expected %s got %s", localacp, acp)
			}

			// These were dummy values, they shouldn't be found.
			if display != "Not Found" {
				t.Errorf("Content error, expected %s got %s", "Not Found", display)
			}

		default:
			// The correct format for Almanac:Plug is constructed with two different arguments.
			localacp := string(tts.args[0] + "." + tts.args[1])
			acp, display, _ = readRainbow(tts.args[0], tts.args[1]) // cmd.Run sends this

			if acp != localacp {
				t.Errorf("Return value error, expected %s got %s", localacp, acp)
			}

			// These were dummy values, they shouldn't be found.
			if display != "Not Found" {
				t.Errorf("Content error, expected %s got %s", "Not Found", display)
			}
		}
	}
}

// TestListAlmanacs ::: Check that the Almanac listing is working.
func TestListAlmanacs(t *testing.T) {
	// Test data is a one-Almanac Rainbow with two Plugs.
	var testmap = []byte(`
[meta.last]
editor = "craque@craque.net"

[status.last]
color = "green"
`)
	var actual = []string{"meta", "status"}

	// Use viper to access the value in this testmap
	viper.SetConfigType("toml")
	viper.ReadConfig(bytes.NewBuffer(testmap))

	// Use viper to submit the config to listRainbow()
	// which returns a TOML object.
	list := listAlmanacs(viper.AllSettings())

	if len(list) != len(actual) {
		t.Errorf("Not equal size lists, expected %d got %d", len(actual), len(list))
	}

	sort.Strings(list)
	sort.Strings(actual)
	for i, v := range list {
		if v != actual[i] {
			t.Errorf("Entry mismatch, expected %s got %s", actual[i], v)
		}
	}
}
