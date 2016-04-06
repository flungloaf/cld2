# cld2
--
    import "cld2"

Package cld2 implements language detection using the Compact Language Detector.

This package includes the relevant sources from the cld2 project, so it doesn't
require any external dependencies. For more information about CLD2, see
https://code.google.com/p/cld2/.

## Usage

#### func  Detect

```go
func Detect(text string) string
```
Detect returns the language code for detected language in the given text.

#### func DetectLang

```go
func DetectLang(text string) Language {
```

DetectLang returns the Language type for detected language in the given text.

#### func DetectThree

```go
func DetectThree(text string) Languages
```

DetectThree returns up to three language guesses.

```go
// Languages are probable languages of the supplied text
type Languages struct {
	Estimates []Estimate // Possible languages returned in order of confidence
	TextBytes int        // the amount of non-tag/letters-only text found
	Reliable  bool       // Does CLD2 see the result as reliable?
}

// Single Language estimate
type Estimate struct {
	Language Language
	Percent  int // text percentage 0..100 of the top 3 languages.

	// NormScore is internal language scores as a ratio to normal score for real text in that language.
	// Scores close to 1.0 indicate normal text, while scores far away
	// from 1.0 indicate badly-skewed text or gibberish.
	NormScore float64
}
```