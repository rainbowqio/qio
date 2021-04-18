package cmd

import (
	"testing"
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
