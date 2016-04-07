//+build !cld2_disable

// Package cld2 implements language detection using the
// Compact Language Detector.
//
// This package includes the relevant sources from the cld2
// project, so it doesn't require any external dependencies.
// For more information about CLD2, see https://github.com/CLD2Owners/cld2
package cld2

// #include <stdlib.h>
// #include "cld2.h"
import "C"
import (
	"unsafe"
)

// Detect returns the language code for detected language
// in the given text.
func Detect(text string) string {
	cs := C.CString(text)
	res := C.DetectLang(cs, -1)
	C.free(unsafe.Pointer(cs))
	var lang string
	if res != nil {
		lang = C.GoString(res)
	}
	return lang
}

// DetectLang returns the language code for detected language
// in the given text.
// ENGLISH is returned if the language cannot be detected.
func DetectLang(text string) Language {
	cs := C.CString(text)
	res := C.DetectLangCode(cs, -1)
	C.free(unsafe.Pointer(cs))
	return Language(res)
}

// DetectThree returns up to three language guesses.
// Extended languages are enabled.
// Unknown languages are removed from the resultset.
func DetectThree(text string) Languages {
	cs := C.CString(text)
	dst := new(C.struct__result)
	C.DetectThree(dst, cs, -1)
	C.free(unsafe.Pointer(cs))
	res := make([]Estimate, 0, 3)
	for i := range dst.language {
		var est Estimate
		est.Language = Language(dst.language[i])
		if est.Language == UNKNOWN_LANGUAGE {
			continue
		}
		est.Percent = int(dst.percent[i])
		est.NormScore = float64(dst.normalized_score[i])
		res = append(res, est)
	}
	rel := true
	if dst.reliable == 0 {
		rel = false
	}
	return Languages{Estimates: res, Reliable: rel, TextBytes: int(dst.text_bytes)}
}
