package cmd

import (
	"bytes"
	"testing"

	toml "github.com/pelletier/go-toml"
	"github.com/spf13/viper"
)

// TestListRainbow ::: Test the Rainbow list function by comparing its output to the source request.
func TestListRainbow(t *testing.T) {
	// Test data is a one-Almanac Rainbow with two Plugs.
	var testmap = []byte(`
[meta.last]
dt = "2021-03-31T03:00:00Z"
editor = "craque@craque.net"
`)

	// Use viper to access the value in this testmap
	viper.SetConfigType("toml")
	viper.ReadConfig(bytes.NewBuffer(testmap))
	viperSent := viper.Get("meta.last.editor")

	// Use viper to submit the config to listRainbow()
	// which returns a TOML object.
	v := viper.AllSettings()
	list, rerr := listRainbow(v)
	if rerr != nil {
		t.Errorf("Error on Return: %s", rerr)
	}
	configmatch, err := toml.Load(list)
	if err != nil {
		t.Errorf("Malformed TOML return value: %s", err)
	}
	listReturn := configmatch.Get("meta.last.editor").(string)

	// Test does not pass if these values do not match
	if viperSent != listReturn {
		t.Errorf("Mismatch error, expected %s got %s", listReturn, viperSent)
	}
}
