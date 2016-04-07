//+build !cld2_disable

package cld2

import (
	"testing"
)

var dkText = `Omkring 4.000 personer har gennem de seneste år forladt EU-landene for at deltage i krigen i Irak og Syrien.
Det viser en ny rapport fra tænketanken ’International Centre for Counter-Terrorism’, som den hollandske regering har fået udarbejdet, da landet i øjeblikket har formandsskabet for EU.
Rapporten dokumenterer også, at 125 syrienkrigere kommer fra Danmark.
På et punkt skiller Danmark sig ud.
- Danmark er det land, der har den højeste andel af krigere, der er vendt tilbage. Cirka 50 procent af de krigere, der er taget til Syrien og Irak, er kommet tilbage nu, siger Tore Hamming, der som dansker forsker i militant islamisme ved European University Institute i Firenze.
LÆS OGSÅ: Joanna kæmpede mod IS: Retten skal nu bestemme om hun skal have sit pas tilbage

Ifølge forskeren er det dog ikke ensbetydende med, at truslen mod Danmark er højere end i andre EU-lande.
- Der er en masse mennesker i Danmark, som har erfaring i at kæmpe og er kommet tilbage med et netværk, de kan bruge til noget. Men hvorvidt, de også er kommet tilbage med intentionen, er sværere at sige, siger han.
Til gengæld ser han fremmedkrigerne som en stor trussel mod alle EU-lande.
- Jeg er desværre lidt pessimistisk og tror, at truslen er værre, end vi overhovedet kan forestille os, siger Tore Hamming.
- Det tror jeg på grund af det vanvittigt høje antal folk, der er draget af sted til Syrien og Irak, og som er kommet tilbage igen.
De hjemvendte er ifølge forskeren yngre end før set, og de vender hjem med 'et forstyrret billede af, hvad vold er, og hvad vold kan bruges til'.`

func TestDetect(t *testing.T) {
	lang := Detect(dkText)
	if lang != "da" {
		t.Fatalf("want 'da', got '%s'", lang)
	}
}

func TestDetectLang(t *testing.T) {
	lang := DetectLang(dkText)
	if lang != DANISH {
		t.Fatalf("want 'DANISH', got '%v'", lang)
	}
}

func TestDetectThree(t *testing.T) {
	guesses := DetectThree(dkText)
	t.Logf("dkText: %+v", guesses)
	if !guesses.Reliable {
		t.Error("want result to be reliable")
	}
	if len(guesses.Estimates) < 1 {
		t.Error("want at least one language estimate")
		return
	}
	est := guesses.Estimates[0]
	if est.Percent < 10 {
		t.Errorf("want percent to be >10 in first estimate: %+v", est)
	}
	if est.Language != DANISH {
		t.Errorf("want language to be DANISH in first estimate: %+v", est)
	}

	guesses = DetectThree(``)
	t.Logf("empty: %+v", guesses)
	if guesses.Reliable {
		t.Error("do not want result to be reliable")
	}
	if len(guesses.Estimates) > 0 {
		t.Error("want no language estimates")
		return
	}

	guesses = DetectThree(`Stringer works best with`)
	t.Logf("short: %+v", guesses)
	if !guesses.Reliable {
		t.Error("want result to be reliable")
	}
	if len(guesses.Estimates) == 0 {
		t.Error("want at least one language estimate")
		return
	}
}

type TestData struct {
	ExpectLanguageCode string
	ExpectLanguageName string
	ExpectIsReliable   bool
	Text               string
}

var testData = [...]TestData{
	{"en", "English", true, "The quick brown fox jumped over the lazy dog"},
	{"fr", "French", true, "Le rapide renard brun sauta par dessus le chien paresseux"},
	{"de", "German", true, "Der schnelle braune Fuchs über den faulen Hund sprang"},
	{"es", "Spanish", true, "el zorro marrón rápido saltó sobre el perro perezoso"},
	{"mk", "Macedonian", true, "брзо кафеава лисица прескокна мрзливи куче"},
	{"zh", "Chinese", true, "敏捷的棕色狐狸跳过了懒狗，目的也许这语料库文本的宽度足以决定"},
	{"ja", "Japanese", true, "速い茶色のキツネは、怠け者の犬を飛び越えました"},
	{"ko", "Korean", true, "빠른 갈색 여우가 게으른 개를 뛰어 넘었다"},
	{"th", "Thai", true, "สุนัขจิ้งจอกสีน้ำตาลได้อย่างรวดเร็วเพิ่มขึ้นกว่าสุนัขขี้เกียจ"},
	{"ar", "Arabic", true, "قفز الثعلب البني السريع فوق الكلب الكسول"},
	{"iw", "Hebrew", true, "שועל החום הזריז קפץ מעל הכלב העצלן, לכוון אולי קורפוס זה של טקסט הוא רחב מספיק כדי להחליט"},
	{"un", "Unknown", false, "no"},
}

func TestDetectShort(t *testing.T) {
	for _, input := range testData {
		actualLanguageCode := Detect(input.Text)
		if actualLanguageCode != input.ExpectLanguageCode {
			t.Errorf("expected `%s`, got `%s` (%s)", input.ExpectLanguageCode, actualLanguageCode, input.Text)
		}
	}
}

func TestDetectShortEstimates(t *testing.T) {
	for _, item := range testData {
		three := DetectThree(item.Text)
		if !three.Reliable {
			if item.ExpectIsReliable {
				t.Error("wanted unreliable result")
			}
			continue
		}
		actual := three.Estimates[0].Language
		if actual.Code() != item.ExpectLanguageCode {
			t.Errorf("expected `%s`, got `%s` (%s)", item.ExpectLanguageCode, actual.Code(), item.Text)
		}
		if actual.String() != item.ExpectLanguageName {
			t.Errorf("expected `%s`, got `%s` (%s)", item.ExpectLanguageName, actual.String(), item.Text)
		}
		t.Logf("`%s` is with %d%% certainty %s (%s)", item.Text, three.Estimates[0].Percent, actual.String(), actual.Code())
	}
}

func BenchmarkDetectLong(b *testing.B) {
	b.SetBytes(int64(len(dkText)))
	for i := 0; i < b.N; i++ {
		_ = Detect(dkText)
	}
}

func BenchmarkDetectLangLong(b *testing.B) {
	b.SetBytes(int64(len(dkText)))
	for i := 0; i < b.N; i++ {
		_ = DetectLang(dkText)
	}
}

func BenchmarkDetectThreeLong(b *testing.B) {
	b.SetBytes(int64(len(dkText)))
	for i := 0; i < b.N; i++ {
		_ = DetectThree(dkText)
	}
}

var shortText = `Freuen Sie sich auf eine Berlin-Story zur Wiedervereinigung und eine bewegende Ost-West-Liebesgeschichte.`

func BenchmarkDetectShort(b *testing.B) {
	b.SetBytes(int64(len(shortText)))
	for i := 0; i < b.N; i++ {
		_ = Detect(shortText)
	}
}

func BenchmarkDetectLangShort(b *testing.B) {
	b.SetBytes(int64(len(shortText)))
	for i := 0; i < b.N; i++ {
		_ = DetectLang(shortText)
	}
}

func BenchmarkDetectThreeShort(b *testing.B) {
	b.SetBytes(int64(len(shortText)))
	for i := 0; i < b.N; i++ {
		_ = DetectThree(shortText)
	}
}
