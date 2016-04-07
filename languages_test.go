package cld2

import "testing"

func TestLanguageFromCode(t *testing.T) {
	l := LanguageFromCode("da")
	if l != DANISH {
		t.Error("expected 'da' code to return Danish, got %s", l.String())
	}

	l = LanguageFromCode("something")
	if l != UNKNOWN_LANGUAGE {
		t.Error("expected 'something' code to return Unknown Language, got %s", l.String())
	}

	l = LanguageFromCode("un")
	if l != UNKNOWN_LANGUAGE {
		t.Error("expected 'something' code to return Unknown Language, got %s", l.String())
	}
}
