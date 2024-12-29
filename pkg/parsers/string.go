package parsers

import (
	"strings"
	"unicode"
)

func CapitalizeStr(value string) string {
	result := ""
	value = strings.ToLower(strings.Trim(value, " "))

	for i, world := range strings.Split(value, " ") {
		switch world {
		// states
		case "ac", "al", "ap", "am", "ba", "ce", "df", "es", "go", "ma", "mt", "ms", "mg", "pa", "pb", "pr", "pe", "pi", "rj", "rn", "rs", "ro", "rr", "sc", "sp", "se", "to":
			world = strings.ToUpper(world)
		// Preposições, Artigos e Conjuções
		case "o", "a", "os", "as", "e", "mas", "ou", "porque", "de", "da", "do", "em", "com", "para", "ante", "entre":
		default:
			world = strings.ToTitle(world)
		}

		if i == 0 {
			result = result + world
		} else {
			result = result + " " + world
		}
	}

	return result
}

func FormatCPF(v string) string {
	return v[0:2] + "." + v[3:5] + "." + v[6:8] + "-" + v[9:10]
}

func GetOnlyDigits(v string) string {
	result := ""
	for _, r := range v {
		if unicode.IsDigit(r) {
			result += string(r)
		}
	}

	return result
}
