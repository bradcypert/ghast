package config

import (
	"testing"
)

func TestConfigParser(t *testing.T) {
	t.Run("Should parse a config properly", func(t *testing.T) {
		k := map[string]interface{}{
			"abc": map[string]interface{}{
				"def": "ghi",
			},
			"r": 2138,
		}

		parsed, err := ParsedConfigToContainerKeys(k)

		if err != nil {
			t.Errorf("failed parsing config")
		}

		if parsed["r"] != 2138 {
			t.Errorf("r not properly parsed")
		}

		if parsed["abc.def"] != "ghi" {
			t.Errorf("Unable to find adc.def\nexpected: ghi\ngot: %v", parsed["abc.def"])
		}
	})
}
