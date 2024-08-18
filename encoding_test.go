package main

import "testing"

func TestEncodingMode(t *testing.T) {
	mode := getModeFromText("0123456789")
	if mode != NumericMode {
		t.Errorf("got %s mode, wanted NUMERIC mode", mode)
	}

	mode = getModeFromText("ABC123 %$")
	if mode != AlphanumericMode {
		t.Errorf("got %s mode, wanted ALPHANUMERIC mode", mode)
	}

	mode = getModeFromText("こんにちは")
	if mode != KanjiMode {
		t.Errorf("got %s mode, wanted KANJI mode", mode)
	}

	mode = getModeFromText("Héllo, Wørld!")
	if mode != ByteMode {
		t.Errorf("got %s mode, wanted BYTE mode", mode)
	}
	mode = getModeFromText("こんにちは😊")
	if mode != ByteMode {
		t.Errorf("got %s mode, wanted BYTE mode", mode)
	}
}