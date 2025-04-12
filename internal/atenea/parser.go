package atenea

import (
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type CourseMetadata struct {
	Codigo              string
	Nombre              string
	Uab                 string
	Vigente             bool
	HorasPresenciales   int
	HorasNoPresenciales int
	Creditos            int
	Validable           bool
	Electiva            bool
	Descripcion         string
}

type CourseCareer struct {
	Code string
	Name string
}

func normalizeString(str string) (string, error) {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r)
	}), norm.NFC)

	normalizedString, _, err := transform.String(t, strings.ToLower(str))
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(normalizedString), nil
}
