package main

import (
	"strings"
	"unicode"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type Version int
type Mode string
type ErrorCorrectionLevel rune

const (
	NumericMode Mode = "NUMERIC"
	AlphanumericMode Mode = "ALPHANUMERIC"
	ByteMode Mode = "BYTE"
	KanjiMode Mode = "NUMERIC"
)

func getModeFromText(text string) Mode {
	if hasOnlyDigits(text) {
		return NumericMode
	}
	if isTextAlphanumeric(text) {
		return AlphanumericMode
	}
	if isShiftJIS(text) {
		return KanjiMode
	}
	return ByteMode
}

func hasOnlyDigits(text string) bool {
	for _, r := range text {
		if unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isTextAlphanumeric(text string) bool {
	CHARS := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:"
	for _, r := range text {
		if !strings.Contains(CHARS, string(r)) {
			return false
		}
	}
	return true
}

func isShiftJIS(s string) bool {
	encoder := japanese.ShiftJIS.NewEncoder()
	_, _, err := transform.String(encoder, s)
	return err == nil
}
