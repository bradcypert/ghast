package config

import (
    "fmt"
    "testing"
)

func TestConfigParser(t *testing.T) {
    t.Run("Should parse a YAML config into JSON properly", func(t *testing.T) {
        var data = `
a: Easy!
b:
  c: 2
`

        parsed, _ := parseYamlString([]byte(data))

        fmt.Printf("--- m:\n%v\n\n", ((*parsed)["b"]))

        if (*parsed)["a"] != "Easy!" {
            t.Errorf("Failed to parse element A")
        }

        if ((*parsed)["b"]).(map[interface{}]interface{})["c"] != 2 {
            t.Errorf("Failed to parse element B.C")
        }
    })

    t.Run("Should parse a JSON config properly", func(t *testing.T) {
        k := map[string]interface{}{
            "abc": map[string]interface{}{
                "def": "ghi",
                "jkl": 1234,
            },
            "r": 2138,
        }

        parsed, err := ParsedConfigToContainerKeys(&k)

        if err != nil {
            t.Errorf("failed parsing config")
        }

        if parsed["r"] != "2138" {
            t.Errorf("r not properly parsed")
        }

        if parsed["abc.def"] != "ghi" {
            t.Errorf("Unable to find adc.def\nexpected: ghi\ngot: %v", parsed["abc.def"])
        }

        if parsed["abc.jkl"] != "1234" {
            t.Errorf("Unable to find adc.def\nexpected: ghi\ngot: %v", parsed["abc.def"])
        }
    })
}
