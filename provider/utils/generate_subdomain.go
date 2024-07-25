package utils

import (
	ung "github.com/dillonstreator/go-unique-name-generator"
	"github.com/dillonstreator/go-unique-name-generator/dictionaries"
)


func GenerateRandomSubdomain() string {
	generator := ung.NewUniqueNameGenerator(
		ung.WithDictionaries(
			[][]string{
				dictionaries.Colors,
				dictionaries.Adjectives,
				dictionaries.Animals,
			},
		),
		ung.WithSeparator("-"),
		ung.WithStyle(ung.Lower),
	)

	return generator.Generate()

}