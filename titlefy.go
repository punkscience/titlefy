package titlefy

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func FormatEnglishTitle(str string) string {
	conjunctions := []string{
		"for",
		"and",
		"nor",
		"but",
		"or",
		"yet",
		"so",
	}

	output := ""

	// Now break up the string
	mots := strings.Split(str, " ")

	for i, mot := range mots {

		// Skip the first token or anything larger than 3 letters
		if i == 0 || len(mot) > 3 {
			output += mot + " "
			continue
		}

		// Look for the ones we want to shift
		for _, conjunction := range conjunctions {
			if strings.EqualFold(mot, conjunction) {
				output += cases.Title(language.English).String(conjunction) + " "
				break
			}
		}
	}

	return strings.Trim(output, " ")

}
